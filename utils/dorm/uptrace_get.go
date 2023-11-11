package dorm

import (
	"github.com/uptrace/bun"
)

// GetDb 获取驱动
func (c *UptraceClient) GetDb() *bun.DB {
	return c.db
}
