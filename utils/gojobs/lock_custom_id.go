package gojobs

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gotime"
	"time"
)

// LockCustomId 上锁
func (c *Client) LockCustomId(ctx context.Context, info *GormModelTask) (string, error) {
	return c.lock(ctx, fmt.Sprintf("%s%s%v%s%v", c.redisConfig.lockKeyPrefix, c.redisConfig.lockKeySeparator, info.Type, c.redisConfig.lockKeySeparator, info.CustomID), fmt.Sprintf("[LockCustomId] 已在%s@%s机器上锁成功，时间：%v", c.config.systemInsideIP, c.config.systemOutsideIP, gotime.Current().Format()), time.Duration(info.Frequency)*3*time.Second)
}

// UnlockCustomId 解锁
func (c *Client) UnlockCustomId(ctx context.Context, info *GormModelTask) error {
	return c.unlock(ctx, fmt.Sprintf("%s%s%v%s%v", c.redisConfig.lockKeyPrefix, c.redisConfig.lockKeySeparator, info.Type, c.redisConfig.lockKeySeparator, info.CustomID))
}

// LockForeverCustomId 永远上锁
func (c *Client) LockForeverCustomId(ctx context.Context, info *GormModelTask) (string, error) {
	return c.lockForever(ctx, fmt.Sprintf("%s%s%v%s%v", c.redisConfig.lockKeyPrefix, c.redisConfig.lockKeySeparator, info.Type, c.redisConfig.lockKeySeparator, info.CustomID), fmt.Sprintf("[LockCustomId] 已在%s@%s机器永远上锁成功，时间：%v", c.config.systemInsideIP, c.config.systemOutsideIP, gotime.Current().Format()))
}
