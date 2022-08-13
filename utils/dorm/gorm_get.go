package dorm

import "gorm.io/gorm"

// GetDb 获取驱动
func (c *GormClient) GetDb() *gorm.DB {
	return c.Db
}
