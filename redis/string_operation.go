package redis

import (
	"context"
	"time"
)

type StringOperation struct {
	ctx context.Context
}

func NewStringOperation() *StringOperation {
	return &StringOperation{ctx: context.Background()}
}

// Set 设置
func (o *StringOperation) Set(key string, value interface{}, attrs ...*OperationAttr) *StringResult {
	exp := OperationAttrs(attrs).Find(AttrExpr)
	if exp == nil {
		exp = time.Second * 0
	}
	return NewStringResult(Rdb.Set(o.ctx, key, value, exp.(time.Duration)).Result())
}

// Get 获取单个
func (o *StringOperation) Get(key string) *StringResult {
	return NewStringResult(Rdb.Get(o.ctx, key).Result())
}

// MGet 获取多个
func (o *StringOperation) MGet(keys ...string) *SliceResult {
	return NewSliceResult(Rdb.MGet(o.ctx, keys...).Result())
}
