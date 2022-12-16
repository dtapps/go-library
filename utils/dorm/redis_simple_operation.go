package dorm

import (
	"context"
	"github.com/go-redis/redis/v9"
	"time"
)

type SimpleOperation struct {
	db *redis.Client
}

func (r *RedisClient) NewSimpleOperation() *SimpleOperation {
	return &SimpleOperation{db: r.Db}
}

// Set 设置
func (o *SimpleOperation) Set(ctx context.Context, key string, value interface{}, attrs ...*OperationAttr) *SimpleResult {
	exp := OperationAttrs(attrs).Find(AttrExpr)
	if exp == nil {
		exp = time.Second * 0
	}
	return NewSimpleResult(o.db.Set(ctx, key, value, exp.(time.Duration)).Result())
}

// Get 获取单个
func (o *SimpleOperation) Get(ctx context.Context, key string) *SimpleResult {
	return NewSimpleResult(o.db.Get(ctx, key).Result())
}

// Del 删除key操作，支持批量删除
func (o *SimpleOperation) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return o.db.Del(ctx, keys...)
}
