package dorm

import (
	"context"
	"encoding/json"
	"time"
)

type DBGttJsonFunc func() interface{}

// SimpleJsonCache 缓存
type SimpleJsonCache struct {
	Operation *StringOperation // 操作类
	Expire    time.Duration    // 过期时间
	DBGetter  DBGttJsonFunc    // 缓存不存在的操作 DB
}

// NewSimpleJsonCache 构造函数
func (r *RedisClient) NewSimpleJsonCache(operation *StringOperation, expire time.Duration) *SimpleJsonCache {
	return &SimpleJsonCache{
		Operation: operation, // 操作类
		Expire:    expire,    // 过期时间
	}
}

// SetCache 设置缓存
func (c *SimpleJsonCache) SetCache(ctx context.Context, key string, value interface{}) {
	c.Operation.Set(ctx, key, value, WithExpire(c.Expire)).Unwrap()
}

// GetCache 获取缓存
func (c *SimpleJsonCache) GetCache(ctx context.Context, key string) (ret interface{}) {
	f := func() string {
		obj := c.DBGetter()
		b, err := json.Marshal(obj)
		if err != nil {
			return ""
		}
		return string(b)
	}
	ret = c.Operation.Get(ctx, key).UnwrapOrElse(f)
	c.SetCache(ctx, key, ret)
	return
}
