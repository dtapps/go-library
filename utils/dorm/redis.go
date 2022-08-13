package dorm

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"time"
)

type ConfigRedisClient struct {
	Addr     string // 地址
	Password string // 密码
	DB       int    // 数据库
	PoolSize int    // 连接池大小
}

// RedisClient
// https://redis.uptrace.dev/
type RedisClient struct {
	Db     *redis.Client      // 驱动
	config *ConfigRedisClient // 配置
}

func NewRedisClient(config *ConfigRedisClient) (*RedisClient, error) {

	c := &RedisClient{}
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
