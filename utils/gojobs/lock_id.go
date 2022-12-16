package gojobs

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojobs/jobs_gorm_model"
	"github.com/dtapps/go-library/utils/gotime"
	"time"
)

// LockId 上锁
func (c *Client) LockId(ctx context.Context, info jobs_gorm_model.Task) (string, error) {
	return c.cache.redisLockClient.Lock(ctx, fmt.Sprintf("%s%s%v%s%v", c.cache.lockKeyPrefix, c.cache.lockKeySeparator, info.Type, c.cache.lockKeySeparator, info.Id), fmt.Sprintf("[LockId] 已在%s@%s机器上锁成功，时间：%v", c.config.systemInsideIp, c.config.systemOutsideIp, gotime.Current().Format()), time.Duration(info.Frequency)*3*time.Second)
}

// UnlockId 解锁
func (c *Client) UnlockId(ctx context.Context, info jobs_gorm_model.Task) error {
	return c.cache.redisLockClient.Unlock(ctx, fmt.Sprintf("%s%s%v%s%v", c.cache.lockKeyPrefix, c.cache.lockKeySeparator, info.Type, c.cache.lockKeySeparator, info.Id))
}

// LockForeverId 永远上锁
func (c *Client) LockForeverId(ctx context.Context, info jobs_gorm_model.Task) (string, error) {
	return c.cache.redisLockClient.LockForever(ctx, fmt.Sprintf("%s%s%v%s%v", c.cache.lockKeyPrefix, c.cache.lockKeySeparator, info.Type, c.cache.lockKeySeparator, info.Id), fmt.Sprintf("[LockForeverId] 已在%s@%s机器永远上锁成功，时间：%v", c.config.systemInsideIp, c.config.systemOutsideIp, gotime.Current().Format()))
}
