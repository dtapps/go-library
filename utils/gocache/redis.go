package gocache

import (
	"context"
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/gojson"
	"time"
)

// RedisConfig 配置
type RedisConfig struct {
	DefaultExpiration time.Duration     // 默认过期时间
	Client            *dorm.RedisClient // 驱动，可选
	Debug             bool              // 调试，可选
}

// Redis https://github.com/go-redis/redis
type Redis struct {
	config *RedisConfig      // 配置
	Client *dorm.RedisClient // 驱动
}

// NewRedis 实例化
func NewRedis(config *RedisConfig) *Redis {
	c := &Redis{config: config}
	c.Client = config.Client
	return c
}

// Set 设置一个key的值
func (r *Redis) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	r.setLog(key)
	return r.Client.GetDb().Set(context.Background(), key, value, expiration).Result()
}

// SetInterface 设置一个key的值
func (r *Redis) SetInterface(key string, value interface{}, expiration time.Duration) (string, error) {
	r.setLog(key)
	marshal, _ := gojson.Marshal(value)
	return r.Client.GetDb().Set(context.Background(), key, marshal, expiration).Result()
}

// SetDefaultExpiration 设置一个key的值，使用全局默认过期时间
func (r *Redis) SetDefaultExpiration(key string, value interface{}) (string, error) {
	r.setLog(key)
	return r.Client.GetDb().Set(context.Background(), key, value, r.config.DefaultExpiration).Result()
}

// SetInterfaceDefaultExpiration 设置一个key的值，使用全局默认过期时间
func (r *Redis) SetInterfaceDefaultExpiration(key string, value interface{}) (string, error) {
	r.setLog(key)
	marshal, _ := gojson.Marshal(value)
	return r.Client.GetDb().Set(context.Background(), key, marshal, r.config.DefaultExpiration).Result()
}

// Get 查询key的值
func (r *Redis) Get(key string) (string, error) {
	r.getLog(key)
	return r.Client.GetDb().Get(context.Background(), key).Result()
}

// GetInterface 查询key的值
func (r *Redis) GetInterface(key string, result interface{}) error {
	r.getLog(key)
	ret, err := r.Client.GetDb().Get(context.Background(), key).Result()
	if err != nil {
		return err
	}
	err = gojson.Unmarshal([]byte(ret), result)
	return nil
}

// GetSet 设置一个key的值，并返回这个key的旧值
func (r *Redis) GetSet(key string, value interface{}) (string, error) {
	return r.Client.GetDb().GetSet(context.Background(), key, value).Result()
}

// SetNX 如果key不存在，则设置这个key的值
func (r *Redis) SetNX(key string, value interface{}, expiration time.Duration) error {
	return r.Client.GetDb().SetNX(context.Background(), key, value, expiration).Err()
}

// SetNXDefaultExpiration 如果key不存在，则设置这个key的值，使用全局默认过期时间
func (r *Redis) SetNXDefaultExpiration(key string, value interface{}) error {
	return r.Client.GetDb().SetNX(context.Background(), key, value, r.config.DefaultExpiration).Err()
}

// MGet 批量查询key的值
func (r *Redis) MGet(keys ...string) ([]interface{}, error) {
	return r.Client.GetDb().MGet(context.Background(), keys...).Result()
}

// MSet 批量设置key的值
// MSet(map[string]interface{}{"key1": "value1", "key2": "value2"})
func (r *Redis) MSet(values map[string]interface{}) error {
	return r.Client.GetDb().MSet(context.Background(), values).Err()
}

// Incr 针对一个key的数值进行递增操作
func (r *Redis) Incr(key string) (int64, error) {
	return r.Client.GetDb().Incr(context.Background(), key).Result()
}

// IncrBy 针对一个key的数值进行递增操作，指定每次递增多少
func (r *Redis) IncrBy(key string, value int64) (int64, error) {
	return r.Client.GetDb().IncrBy(context.Background(), key, value).Result()
}

// Decr 针对一个key的数值进行递减操作
func (r *Redis) Decr(key string) (int64, error) {
	return r.Client.GetDb().Decr(context.Background(), key).Result()
}

// DecrBy 针对一个key的数值进行递减操作，指定每次递减多少
func (r *Redis) DecrBy(key string, value int64) (int64, error) {
	return r.Client.GetDb().DecrBy(context.Background(), key, value).Result()
}

// Del 删除key操作，支持批量删除
func (r *Redis) Del(keys ...string) error {
	r.delLog(keys...)
	return r.Client.GetDb().Del(context.Background(), keys...).Err()
}
