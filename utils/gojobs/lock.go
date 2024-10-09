package gojobs

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.dtapp.net/library/utils/gotime"
	"time"
)

// Lock 上锁
func (c *Client) Lock(ctx context.Context, info *GormModelTask, id any) (string, error) {
	return c.lock(ctx, fmt.Sprintf("%s%s%v%s%v", c.redisConfig.lockKeyPrefix, c.redisConfig.lockKeySeparator, info.Type, c.redisConfig.lockKeySeparator, id), fmt.Sprintf("[Lock] 已在%s@%s机器上锁成功，时间：%v", c.config.systemInsideIP, c.config.systemOutsideIP, gotime.Current().Format()), time.Duration(info.Frequency)*3*time.Second)
}

// LockMinute 上锁
func (c *Client) LockMinute(ctx context.Context, info *GormModelTask, id any, minute int64) (string, error) {
	return c.lock(ctx, fmt.Sprintf("%s%s%v%s%v", c.redisConfig.lockKeyPrefix, c.redisConfig.lockKeySeparator, info.Type, c.redisConfig.lockKeySeparator, id), fmt.Sprintf("[Lock] 已在%s@%s机器上锁成功，时间：%v", c.config.systemInsideIP, c.config.systemOutsideIP, gotime.Current().Format()), time.Duration(minute)*time.Minute)
}

// Unlock Lock 解锁
func (c *Client) Unlock(ctx context.Context, info *GormModelTask, id any) error {
	return c.unlock(ctx, fmt.Sprintf("%s%s%v%s%v", c.redisConfig.lockKeyPrefix, c.redisConfig.lockKeySeparator, info.Type, c.redisConfig.lockKeySeparator, id))
}

// LockForever 永远上锁
func (c *Client) LockForever(ctx context.Context, info *GormModelTask, id any) (string, error) {
	return c.lockForever(ctx, fmt.Sprintf("%s%s%v%s%v", c.redisConfig.lockKeyPrefix, c.redisConfig.lockKeySeparator, info.Type, c.redisConfig.lockKeySeparator, id), fmt.Sprintf("[LockForever] 已在%s@%s机器永远上锁成功，时间：%v", c.config.systemInsideIP, c.config.systemOutsideIP, gotime.Current().Format()))
}

// Lock 上锁
// key 锁名
// val 锁内容
// ttl 锁过期时间
func (c *Client) lock(ctx context.Context, key string, val string, ttl time.Duration) (resp string, err error) {
	if ttl <= 0 {
		return resp, errors.New("长期请使用 LockForever 方法")
	}
	// 获取
	get, err := c.redisConfig.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		// 设置
		err = c.redisConfig.client.Set(ctx, key, val, ttl).Err()
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
func (c *Client) unlock(ctx context.Context, key string) error {
	_, err := c.redisConfig.client.Del(ctx, key).Result()
	if err != nil {
		return errors.New(fmt.Sprintf("解锁失败：%s", err.Error()))
	}
	return err
}

// LockForever 永远上锁
// key 锁名
// val 锁内容
func (c *Client) lockForever(ctx context.Context, key string, val string) (resp string, err error) {
	// 获取
	get, err := c.redisConfig.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		// 设置
		err = c.redisConfig.client.Set(ctx, key, val, 0).Err()
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
