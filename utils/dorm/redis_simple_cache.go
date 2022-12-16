package dorm

import (
	"context"
	"encoding/json"
	"time"
)

const (
	SerializerJson   = "json"
	SerializerString = "string"
)

type JsonGttFunc func() interface{}

type DBGttFunc func() string

// SimpleCache 缓存
type SimpleCache struct {
	Operation  *StringOperation // 操作类
	Expire     time.Duration    // 过去时间
	DBGetter   DBGttFunc        // 缓存不存在的操作 DB
	JsonGetter JsonGttFunc      // 缓存不存在的操作 JSON
	Serializer string           // 序列化方式
}

// NewSimpleCache 构造函数
func (r *RedisClient) NewSimpleCache(operation *StringOperation, expire time.Duration, serializer string) *SimpleCache {
	return &SimpleCache{
		Operation:  operation,  // 操作类
		Expire:     expire,     // 过去时间
		Serializer: serializer, // 缓存不存在的操作 DB
	}
}

// SetCache 设置缓存
func (c *SimpleCache) SetCache(ctx context.Context, key string, value interface{}) {
	c.Operation.Set(ctx, key, value, WithExpire(c.Expire)).Unwrap()
}

// GetCache 获取缓存
func (c *SimpleCache) GetCache(ctx context.Context, key string) (ret interface{}) {
	if c.Serializer == SerializerJson {
		f := func() string {
			obj := c.JsonGetter()
			b, err := json.Marshal(obj)
			if err != nil {
				return ""
			}
			return string(b)
		}
		ret = c.Operation.Get(ctx, key).UnwrapOrElse(f)
		c.SetCache(ctx, key, ret)
	} else if c.Serializer == SerializerString {
		f := func() string {
			return c.DBGetter()
		}
		ret = c.Operation.Get(ctx, key).UnwrapOrElse(f)
		c.SetCache(ctx, key, ret)
	}
	return
}
