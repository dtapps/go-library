package dorm

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"time"
)

// RedisClientLock https://github.com/go-redis/redis
type RedisClientLock struct {
	operation *RedisClient // 操作
}

// NewLock 实例化锁
func (r *RedisClient) NewLock() *RedisClientLock {
	return &RedisClientLock{r}
}

// Lock 上锁
// key 锁名
// val 锁内容
// ttl 锁过期时间
func (rl *RedisClientLock) Lock(ctx context.Context, key string, val string, ttl time.Duration) (resp string, err error) {
	if ttl <= 0 {
		return resp, errors.New("长期请使用 LockForever 方法")
	}
	// 获取
	get, err := rl.operation.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		// 设置
		err = rl.operation.Set(ctx, key, val, ttl).Err()
		if err != nil {
			return resp, errors.New(fmt.Sprintf("上锁失败：%s", err.Error()))
		}
		return val, nil
	}
	if get != "" {
		return resp, errors.New("上锁失败，已存在")
	}
	return resp, errors.New(fmt.Sprintf("获取异常：%s", err.Error()))
}

// Unlock 解锁
// key 锁名
func (rl *RedisClientLock) Unlock(ctx context.Context, key string) error {
	_, err := rl.operation.Del(ctx, key).Result()
	if err != nil {
		return errors.New(fmt.Sprintf("解锁失败：%s", err.Error()))
	}
	return err
}

// LockForever 永远上锁
// key 锁名
// val 锁内容
func (rl *RedisClientLock) LockForever(ctx context.Context, key string, val string) (resp string, err error) {
	// 获取
	get, err := rl.operation.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		// 设置
		err = rl.operation.Set(ctx, key, val, 0).Err()
		if err != nil {
			return resp, errors.New(fmt.Sprintf("上锁失败：%s", err.Error()))
		}
		return val, nil
	}
	if get != "" {
		return resp, errors.New("上锁失败，已存在")
	}
	return resp, errors.New(fmt.Sprintf("获取异常：%s", err.Error()))
}
