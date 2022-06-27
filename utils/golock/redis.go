package golock

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type LockRedis struct {
	redisClient *redis.Client // 驱动
}

func NewLockRedis(redisClient *redis.Client) *LockRedis {
	return &LockRedis{redisClient: redisClient}
}

// Lock 上锁
// key 锁名
// val 锁内容
// ttl 锁过期时间
func (r *LockRedis) Lock(key string, val string, ttl int64) (string, error) {
	if ttl <= 0 {
		return "", errors.New("长期请使用 LockForever 方法")
	}
	// 1、获取
	get, err := r.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		return "", errors.New("获取异常")
	}
	if get != "" {
		return "", errors.New("上锁失败，已存在")
	}
	// 2、设置
	err = r.redisClient.Set(context.Background(), key, val, time.Duration(ttl)).Err()
	if err != nil {
		return "", errors.New("上锁失败")
	}
	return val, nil
}

// Unlock 解锁
// key 锁名
func (r *LockRedis) Unlock(key string) error {
	_, err := r.redisClient.Del(context.Background(), key).Result()
	return err
}

// LockForever 永远上锁
// key 锁名
// val 锁内容
func (r *LockRedis) LockForever(key string, val string) (string, error) {
	// 1、获取
	get, err := r.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		return "", errors.New("获取异常")
	}
	if get != "" {
		return "", errors.New("上锁失败，已存在")
	}
	// 2、设置
	err = r.redisClient.Set(context.Background(), key, val, 0).Err()
	if err != nil {
		return "", errors.New("上锁失败")
	}
	return val, nil
}
