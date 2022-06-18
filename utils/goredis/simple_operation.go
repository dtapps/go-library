package goredis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type SimpleOperation struct {
	db  *redis.Client
	ctx context.Context
}

func (app *App) NewSimpleOperation() *SimpleOperation {
	return &SimpleOperation{
		db:  app.Db,
		ctx: context.Background(),
	}
}

// Set 设置
func (o *SimpleOperation) Set(key string, value interface{}, attrs ...*OperationAttr) *SimpleResult {
	exp := OperationAttrs(attrs).Find(AttrExpr)
	if exp == nil {
		exp = time.Second * 0
	}
	return NewSimpleResult(o.db.Set(o.ctx, key, value, exp.(time.Duration)).Result())
}

// Get 获取单个
func (o *SimpleOperation) Get(key string) *SimpleResult {
	return NewSimpleResult(o.db.Get(o.ctx, key).Result())
}

// Del 删除key操作，支持批量删除
func (o *SimpleOperation) Del(keys ...string) *redis.IntCmd {
	return o.db.Del(o.ctx, keys...)
}
