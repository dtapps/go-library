package dorm

import (
	"context"
	"github.com/go-redis/redis/v9"
)

type HashOperation struct {
	db  *redis.Client
	ctx context.Context
}

// NewHashOperation hash类型数据操作 https://www.tizi365.com/archives/296.html
func NewHashOperation(db *redis.Client, ctx context.Context) *HashOperation {
	return &HashOperation{db: db, ctx: ctx}
}

// Set 根据key和field字段设置，field字段的值
func (cl *HashOperation) Set(key string, value interface{}) *redis.IntCmd {
	return cl.db.HSet(cl.ctx, key, value)
}

// Get 根据key和field字段设置，field字段的值
func (cl *HashOperation) Get(key, field string) *redis.StringCmd {
	return cl.db.HGet(cl.ctx, key, field)
}
