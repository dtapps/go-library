package gojobs

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojobs/jobs_gorm_model"
	"github.com/dtapps/go-library/utils/gotime"
	"time"
)

// Lock 上锁
func (c *Client) Lock(ctx context.Context, info jobs_gorm_model.Task, id any) (string, error) {
	return c.cache.redisLockClient.Lock(ctx, fmt.Sprintf("%s%s%v%s%v", c.cache.lockKeyPrefix, c.cache.lockKeySeparator, info.Type, c.cache.lockKeySeparator, id), fmt.Sprintf("[Lock] 已在%s@%s机器上锁成功，时间：%v", c.config.systemInsideIp, c.config.systemOutsideIp, gotime.Current().Format()), time.Duration(info.Frequency)*3*time.Second)
}

// Unlock Lock 解锁
func (c *Client) Unlock(ctx context.Context, info jobs_gorm_model.Task, id any) error {
	return c.cache.redisLockClient.Unlock(ctx, fmt.Sprintf("%s%s%v%s%v", c.cache.lockKeyPrefix, c.cache.lockKeySeparator, info.Type, c.cache.lockKeySeparator, id))
}

// LockForever 永远上锁
func (c *Client) LockForever(ctx context.Context, info jobs_gorm_model.Task, id any) (string, error) {
	return c.cache.redisLockClient.LockForever(ctx, fmt.Sprintf("%s%s%v%s%v", c.cache.lockKeyPrefix, c.cache.lockKeySeparator, info.Type, c.cache.lockKeySeparator, id), fmt.Sprintf("[LockForever] 已在%s@%s机器永远上锁成功，时间：%v", c.config.systemInsideIp, c.config.systemOutsideIp, gotime.Current().Format()))
}
