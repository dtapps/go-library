package gojobs

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gotime"
	"time"
)

// LockId 上锁
func (c *Client) LockId(ctx context.Context, info *GormModelTask) (string, error) {
	return c.lock(ctx, fmt.Sprintf("%s%s%v%s%v", c.redisConfig.lockKeyPrefix, c.redisConfig.lockKeySeparator, info.Type, c.redisConfig.lockKeySeparator, info.ID), fmt.Sprintf("[LockId] 已在%s@%s机器上锁成功，时间：%v", c.config.systemInsideIP, c.config.systemOutsideIP, gotime.Current().Format()), time.Duration(info.Frequency)*3*time.Second)
}

// LockIdMinute 上锁
func (c *Client) LockIdMinute(ctx context.Context, info *GormModelTask, minute int64) (string, error) {
	return c.lock(ctx, fmt.Sprintf("%s%s%v%s%v", c.redisConfig.lockKeyPrefix, c.redisConfig.lockKeySeparator, info.Type, c.redisConfig.lockKeySeparator, info.ID), fmt.Sprintf("[LockId] 已在%s@%s机器上锁成功，时间：%v", c.config.systemInsideIP, c.config.systemOutsideIP, gotime.Current().Format()), time.Duration(minute)*time.Minute)
}

// UnlockId 解锁
func (c *Client) UnlockId(ctx context.Context, info *GormModelTask) error {
	return c.unlock(ctx, fmt.Sprintf("%s%s%v%s%v", c.redisConfig.lockKeyPrefix, c.redisConfig.lockKeySeparator, info.Type, c.redisConfig.lockKeySeparator, info.ID))
}

// LockForeverId 永远上锁
func (c *Client) LockForeverId(ctx context.Context, info *GormModelTask) (string, error) {
	return c.lockForever(ctx, fmt.Sprintf("%s%s%v%s%v", c.redisConfig.lockKeyPrefix, c.redisConfig.lockKeySeparator, info.Type, c.redisConfig.lockKeySeparator, info.ID), fmt.Sprintf("[LockForeverId] 已在%s@%s机器永远上锁成功，时间：%v", c.config.systemInsideIP, c.config.systemOutsideIP, gotime.Current().Format()))
}
