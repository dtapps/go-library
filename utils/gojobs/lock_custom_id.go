package gojobs

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojobs/jobs_gorm_model"
	"github.com/dtapps/go-library/utils/gotime"
	"time"
)

// LockCustomId 上锁
func (c *Client) LockCustomId(ctx context.Context, info jobs_gorm_model.Task) (string, error) {
	return c.cache.redisLockClient.Lock(ctx, fmt.Sprintf("%s%s%v%s%v", c.cache.lockKeyPrefix, c.cache.lockKeySeparator, info.Type, c.cache.lockKeySeparator, info.CustomId), fmt.Sprintf("[LockCustomId] 已在%s@%s机器上锁成功，时间：%v", c.config.systemInsideIp, c.config.systemOutsideIp, gotime.Current().Format()), time.Duration(info.Frequency)*3*time.Second)
}

// UnlockCustomId 解锁
func (c *Client) UnlockCustomId(ctx context.Context, info jobs_gorm_model.Task) error {
	return c.cache.redisLockClient.Unlock(ctx, fmt.Sprintf("%s%s%v%s%v", c.cache.lockKeyPrefix, c.cache.lockKeySeparator, info.Type, c.cache.lockKeySeparator, info.CustomId))
}

// LockForeverCustomId 永远上锁
func (c *Client) LockForeverCustomId(ctx context.Context, info jobs_gorm_model.Task) (string, error) {
	return c.cache.redisLockClient.LockForever(ctx, fmt.Sprintf("%s%s%v%s%v", c.cache.lockKeyPrefix, c.cache.lockKeySeparator, info.Type, c.cache.lockKeySeparator, info.CustomId), fmt.Sprintf("[LockCustomId] 已在%s@%s机器永远上锁成功，时间：%v", c.config.systemInsideIp, c.config.systemOutsideIp, gotime.Current().Format()))
}
