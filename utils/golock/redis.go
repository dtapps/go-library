package golock

import (
	"context"
	"errors"
	"github.com/dtapps/go-library/utils/dorm"
	"time"
)

type LockRedis struct {
	client *dorm.RedisClient // 驱动
}

func NewLockRedis(client *dorm.RedisClient) *LockRedis {
	return &LockRedis{client: client}
}

// Lock 上锁
// key 锁名
// val 锁内容
// ttl 锁过期时间
func (r *LockRedis) Lock(ctx context.Context, key string, val string, ttl time.Duration) (resp string, err error) {
	if ttl <= 0 {
		return resp, errors.New("长期请使用 LockForever 方法")
	}
	// 获取
	get, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return resp, errors.New("获取异常")
	}
	if get != "" {
		return resp, errors.New("上锁失败，已存在")
	}
	// 设置
	err = r.client.Set(ctx, key, val, ttl).Err()
	if err != nil {
		return resp, errors.New("上锁失败")
	}
	return val, nil
}

// Unlock 解锁
// key 锁名
func (r *LockRedis) Unlock(ctx context.Context, key string) error {
	_, err := r.client.Del(ctx, key).Result()
	return err
}

// LockForever 永远上锁
// key 锁名
// val 锁内容
func (r *LockRedis) LockForever(ctx context.Context, key string, val string) (resp string, err error) {
	// 获取
	get, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return resp, errors.New("获取异常")
	}
	if get != "" {
		return resp, errors.New("上锁失败，已存在")
	}
	// 设置
	err = r.client.Set(ctx, key, val, 0).Err()
	if err != nil {
		return resp, errors.New("上锁失败")
	}
	return val, nil
}
