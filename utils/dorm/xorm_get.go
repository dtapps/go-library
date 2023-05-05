package dorm

import (
	"xorm.io/builder"
	"xorm.io/xorm"
)

// GetDb 获取驱动
func (c *XormClient) GetDb() *xorm.Engine {
	return c.db
}

func (c *XormClient) GetBuilder(dialect string) *builder.Builder {
	return builder.Dialect(dialect)
}
