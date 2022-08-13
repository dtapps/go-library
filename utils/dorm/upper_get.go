package dorm

import (
	"github.com/upper/db/v4"
)

// GetDb 获取驱动
func (c *UpperClient) GetDb() *db.Session {
	return c.Db
}
