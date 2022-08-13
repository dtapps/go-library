package dorm

import (
	"context"
	"encoding/json"
	"time"
)

// GttStringFunc String缓存结构
type GttStringFunc func() string

// GttInterfaceFunc Interface缓存结构
type GttInterfaceFunc func() interface{}

// RedisCacheConfig 配置
type RedisCacheConfig struct {
	Expiration time.Duration // 过期时间
}

// RedisClientCache https://github.com/go-redis/redis
type RedisClientCache struct {
	defaultExpiration time.Duration    // 过期时间
	operation         *RedisClient     // 操作
	GetterString      GttStringFunc    // 不存在的操作
	GetterInterface   GttInterfaceFunc // 不存在的操作
}

// NewCache 实例化
func (r *RedisClient) NewCache(config *RedisCacheConfig) *RedisClientCache {
	return &RedisClientCache{
		defaultExpiration: config.Expiration,
		operation:         r,
	}
}

// NewCacheDefaultExpiration 实例化
func (r *RedisClient) NewCacheDefaultExpiration() *RedisClientCache {
	return &RedisClientCache{
		defaultExpiration: time.Minute * 30,
		operation:         r,
	}
}

// GetString 缓存操作
func (rc *RedisClientCache) GetString(ctx context.Context, key string) (ret string) {

	f := func() string {
		return rc.GetterString()
	}

	// 如果不存在，则调用GetterString
	ret, err := rc.operation.Get(ctx, key).Result()
	if err != nil {
		rc.operation.Set(ctx, key, f(), rc.defaultExpiration)
		ret, _ = rc.operation.Get(ctx, key).Result()
	}

	return
}

// GetInterface 缓存操作
func (rc *RedisClientCache) GetInterface(ctx context.Context, key string, result interface{}) {

	f := func() string {
		marshal, _ := json.Marshal(rc.GetterInterface())
		return string(marshal)
	}

	// 如果不存在，则调用GetterInterface
	ret, err := rc.operation.Get(ctx, key).Result()

	if err != nil {
		rc.operation.Set(ctx, key, f(), rc.defaultExpiration)
		ret, _ = rc.operation.Get(ctx, key).Result()
	}

	err = json.Unmarshal([]byte(ret), result)

	return
}

// GetInterfaceKey 获取key值
func (rc *RedisClientCache) GetInterfaceKey(ctx context.Context, key string, result interface{}) error {
	ret, err := rc.operation.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(ret), result)
	return nil
}

// SetInterfaceKey 设置key值
func (rc *RedisClientCache) SetInterfaceKey(ctx context.Context, key string, value interface{}) (string, error) {
	marshal, _ := json.Marshal(value)
	return rc.operation.Set(ctx, key, marshal, rc.defaultExpiration).Result()
}
