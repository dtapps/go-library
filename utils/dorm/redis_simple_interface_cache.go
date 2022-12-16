package dorm

import (
	"context"
	"time"
)

type DBGttInterfaceFunc func() interface{}

// SimpleInterfaceCache 缓存
type SimpleInterfaceCache struct {
	Operation *SimpleOperation   // 操作类
	Expire    time.Duration      // 过期时间
	DBGetter  DBGttInterfaceFunc // 缓存不存在的操作 DB
}

// NewSimpleInterfaceCache 构造函数
func (r *RedisClient) NewSimpleInterfaceCache(operation *SimpleOperation, expire time.Duration) *SimpleInterfaceCache {
	return &SimpleInterfaceCache{
		Operation: operation, // 操作类
		Expire:    expire,    // 过期时间
	}
}

// SetCache 设置缓存
func (c *SimpleInterfaceCache) SetCache(ctx context.Context, key string, value interface{}) {
	c.Operation.Set(ctx, key, value, WithExpire(c.Expire)).Unwrap()
}

// GetCache 获取缓存
func (c *SimpleInterfaceCache) GetCache(ctx context.Context, key string) (ret interface{}) {
	f := func() interface{} {
		return c.DBGetter()
	}
	ret = c.Operation.Get(ctx, key).UnwrapOrElse(f)
	c.SetCache(ctx, key, ret)
	return ret
}
