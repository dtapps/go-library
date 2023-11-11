package dorm

import (
	"context"
	"xorm.io/builder"
)

func (c *XormClient) GetBuilder(ctx context.Context, dialect string) *builder.Builder {
	return builder.Dialect(dialect)
}
