package gojobs

import (
	"github.com/robfig/cron/v3"
)

// GetDrive 获取驱动
func (c *Cron) GetDrive() *cron.Cron {
	return c.inner
}
