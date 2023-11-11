package dorm

import "xorm.io/builder"

func (c *XormClient) GetBuilder(dialect string) *builder.Builder {
	return builder.Dialect(dialect)
}
