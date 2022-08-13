package gocache

import (
	"encoding/json"
	"time"
)

// RedisCacheConfig 配置
type RedisCacheConfig struct {
	expiration time.Duration // 过期时间
}

// RedisCache https://github.com/go-redis/redis
type RedisCache struct {
	config          *RedisCacheConfig
	operation       *Redis           // 操作
	GetterString    GttStringFunc    // 不存在的操作
	GetterInterface GttInterfaceFunc // 不存在的操作
}

// NewCache 实例化
func (r *Redis) NewCache(config *RedisCacheConfig) *RedisCache {
	c := &RedisCache{config: config}
	c.operation = r
	return c
}

// NewCacheDefaultExpiration 实例化
func (r *Redis) NewCacheDefaultExpiration() *RedisCache {
	c := &RedisCache{}
	c.config.expiration = r.config.DefaultExpiration
	c.operation = r
	return c
}

// GetString 缓存操作
func (rc *RedisCache) GetString(key string) (ret string) {

	f := func() string {
		return rc.GetterString()
	}

	// 如果不存在，则调用GetterString
	ret, err := rc.operation.Get(key)
	if err != nil {
		rc.operation.Set(key, f(), rc.config.expiration)
		ret, _ = rc.operation.Get(key)
	}

	return
}

// GetInterface 缓存操作
func (rc *RedisCache) GetInterface(key string, result interface{}) {

	f := func() string {
		marshal, _ := json.Marshal(rc.GetterInterface())
		return string(marshal)
	}

	// 如果不存在，则调用GetterInterface
	ret, err := rc.operation.Get(key)

	if err != nil {
		rc.operation.Set(key, f(), rc.config.expiration)
		ret, _ = rc.operation.Get(key)
	}

	err = json.Unmarshal([]byte(ret), result)

	return
}
