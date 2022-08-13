package dorm

import (
	"context"
	"github.com/go-redis/redis/v9"
	"time"
)

type StringOperation struct {
	db  *redis.Client
	ctx context.Context
}

func (r *RedisClient) NewStringOperation() *StringOperation {
	return &StringOperation{
		db:  r.Db,
		ctx: context.Background(),
	}
}

// Set 设置
func (o *StringOperation) Set(key string, value interface{}, attrs ...*OperationAttr) *StringResult {
	exp := OperationAttrs(attrs).Find(AttrExpr)
	if exp == nil {
		exp = time.Second * 0
	}
	return NewStringResult(o.db.Set(o.ctx, key, value, exp.(time.Duration)).Result())
}

// Get 获取单个
func (o *StringOperation) Get(key string) *StringResult {
	return NewStringResult(o.db.Get(o.ctx, key).Result())
}

// MGet 获取多个
func (o *StringOperation) MGet(keys ...string) *SliceResult {
	return NewSliceResult(o.db.MGet(o.ctx, keys...).Result())
}

// Del 删除key操作，支持批量删除
func (o *StringOperation) Del(keys ...string) *redis.IntCmd {
	return o.db.Del(o.ctx, keys...)
}
