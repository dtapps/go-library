package dorm

import "github.com/go-redis/redis/v9"

// GetDb 获取驱动
func (r *RedisClient) GetDb() *redis.Client {
	return r.Db
}
