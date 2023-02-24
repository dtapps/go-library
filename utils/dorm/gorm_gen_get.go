package dorm

import (
	"gorm.io/gen"
	"gorm.io/gorm"
)

// GetDb 获取驱动
func (c *GormGenClient) GetDb() *gorm.DB {
	return c.db
}

// GetGenerator 获取驱动
func (c *GormGenClient) GetGenerator() *gen.Generator {
	return c.generator
}
