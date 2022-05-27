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
	RedisCacheConfig
	db              *Redis           // 驱动
	GetterString    GttStringFunc    // 不存在的操作
	GetterInterface GttInterfaceFunc // 不存在的操作
}

// NewCache 实例化
func (r *Redis) NewCache(config *RedisCacheConfig) *RedisCache {
	app := &RedisCache{}
	app.expiration = config.expiration
	app.db = r
	return app
}

// NewCacheDefaultExpiration 实例化
func (r *Redis) NewCacheDefaultExpiration() *RedisCache {
	app := &RedisCache{}
	app.expiration = r.DefaultExpiration
	app.db = r
	return app
}

// GetString 缓存操作
func (rc *RedisCache) GetString(key string) (ret string) {

	f := func() string {
		return rc.GetterString()
	}

	// 如果不存在，则调用GetterString
	ret, err := rc.db.Get(key)
	if err != nil {
		rc.db.Set(key, f(), rc.expiration)
		ret, _ = rc.db.Get(key)
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
	ret, err := rc.db.Get(key)

	if err != nil {
		rc.db.Set(key, f(), rc.expiration)
		ret, _ = rc.db.Get(key)
	}

	err = json.Unmarshal([]byte(ret), result)

	return
}
