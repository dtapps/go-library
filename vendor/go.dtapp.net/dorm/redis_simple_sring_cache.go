package dorm

import (
	"time"
)

type DBGttStringFunc func() string

// SimpleStringCache 缓存
type SimpleStringCache struct {
	Operation *StringOperation // 操作类
	Expire    time.Duration    // 过期时间
	DBGetter  DBGttStringFunc  // 缓存不存在的操作 DB
}

// NewSimpleStringCache 构造函数
func (c *RedisClient) NewSimpleStringCache(operation *StringOperation, expire time.Duration) *SimpleStringCache {
	return &SimpleStringCache{
		Operation: operation, // 操作类
		Expire:    expire,    // 过期时间
	}
}

// SetCache 设置缓存
func (c *SimpleStringCache) SetCache(key string, value string) {
	c.Operation.Set(key, value, WithExpire(c.Expire)).Unwrap()
}

// GetCache 获取缓存
func (c *SimpleStringCache) GetCache(key string) (ret string) {
	f := func() string {
		return c.DBGetter()
	}
	ret = c.Operation.Get(key).UnwrapOrElse(f)
	c.SetCache(key, ret)
	return
}
