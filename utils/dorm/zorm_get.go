package dorm

import (
	"gitee.com/chunanyong/zorm"
)

// GetDb 获取驱动
func (c *ZormClient) GetDb() *zorm.DBDao {
	return c.db
}
