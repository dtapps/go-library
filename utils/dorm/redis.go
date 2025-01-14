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
	db *redis.Client // 驱动
}

// NewRedisClient 创建实例
func NewRedisClient(config *RedisClientConfig) (*RedisClient, error) {

	c := &RedisClient{}
	if config.PoolSize == 0 {
		config.PoolSize = 100
	}

	c.db = redis.NewClient(&redis.Options{
		Addr:        config.Addr,        // 地址
		Password:    config.Password,    // 密码
		DB:          config.DB,          // 数据库
		PoolSize:    config.PoolSize,    // 连接池大小
		ReadTimeout: config.ReadTimeout, // 读取超时
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 检测 Redis 连接是否正常连接
	_, err := c.db.Ping(ctx).Result()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	return c, nil
}

// NewRedisClientURL 创建实例
func NewRedisClientURL(redisURL string) (*RedisClient, error) {

	c := &RedisClient{}

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return c, nil
	}

	// 创建 Redis 客户端
	c.db = redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 检测 Redis 连接是否正常连接
	_, err = c.db.Ping(ctx).Result()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	return c, nil
}
