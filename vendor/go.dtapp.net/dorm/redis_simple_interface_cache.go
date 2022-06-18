package dorm

import (
	"log"
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
func (c *RedisClient) NewSimpleInterfaceCache(operation *SimpleOperation, expire time.Duration) *SimpleInterfaceCache {
	return &SimpleInterfaceCache{
		Operation: operation, // 操作类
		Expire:    expire,    // 过期时间
	}
}

// SetCache 设置缓存
func (c *SimpleInterfaceCache) SetCache(key string, value interface{}) {
	c.Operation.Set(key, value, WithExpire(c.Expire)).Unwrap()
}

// GetCache 获取缓存
func (c *SimpleInterfaceCache) GetCache(key string) (ret interface{}) {
	f := func() interface{} {
		return c.DBGetter()
	}
	ret = c.Operation.Get(key).UnwrapOrElse(f)
	c.SetCache(key, ret)
	log.Println(ret)
	return ret
}
