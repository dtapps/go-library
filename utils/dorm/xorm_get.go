package dorm

import (
	"xorm.io/xorm"
)

// GetDb 获取驱动
func (c *XormClient) GetDb() *xorm.Engine {
	return c.Db
}
