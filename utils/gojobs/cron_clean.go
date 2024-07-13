package gojobs

import (
	"context"
	"github.com/robfig/cron/v3"
)

// StartCronClean 定时清理任务日志
func (c *Client) StartCronClean(ctx context.Context, cr *cron.Cron, cp string, hour int64) (cron.EntryID, error) {
	return cr.AddFunc(cp, func() {
		_ = c.GormTaskLogDelete(ctx, hour)
	})
}
