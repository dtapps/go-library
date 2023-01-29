package dorm

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// RedisClientFun *RedisClient 驱动
type RedisClientFun func() *RedisClient

type RedisClientConfig struct {
	Addr        string        // 地址
	Password    string        // 密码
	DB          int           // 数据库
	PoolSize    int           // 连接池大小
	ReadTimeout time.Duration // 读取超时
}

// RedisClient
// https://redis.uptrace.dev/
type RedisClient struct {
	Db     *redis.Client      // 驱动
	config *RedisClientConfig // 配置
}

func NewRedisClient(config *RedisClientConfig) (*RedisClient, error) {

	c := &RedisClient{}
	c.config = config
	if c.config.PoolSize == 0 {
		c.config.PoolSize = 100
	}

	c.Db = redis.NewClient(&redis.Options{
		Addr:        c.config.Addr,        // 地址
		Password:    c.config.Password,    // 密码
		DB:          c.config.DB,          // 数据库
		PoolSize:    c.config.PoolSize,    // 连接池大小
		ReadTimeout: c.config.ReadTimeout, // 读取超时
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.Db.Ping(ctx).Result()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	return c, nil
}
