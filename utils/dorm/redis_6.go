package dorm

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// RedisClient6
// https://redis.uptrace.dev/
type RedisClient6 struct {
	Db     *redis.Client      // 驱动
	config *ConfigRedisClient // 配置
}

// NewRedisClient6 Redis 6
func NewRedisClient6(config *ConfigRedisClient) (*RedisClient6, error) {

	c := &RedisClient6{}
	c.config = config
	if c.config.PoolSize == 0 {
		c.config.PoolSize = 100
	}

	c.Db = redis.NewClient(&redis.Options{
		Addr:     c.config.Addr,     // 地址
		Password: c.config.Password, // 密码
		DB:       c.config.DB,       // 数据库
		PoolSize: c.config.PoolSize, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.Db.Ping(ctx).Result()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	return c, nil
}

func (c *RedisClient6) GetDb() *redis.Client {
	return c.Db
}
