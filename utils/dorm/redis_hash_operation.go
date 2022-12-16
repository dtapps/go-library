package dorm

import (
	"context"
	"github.com/go-redis/redis/v9"
)

type HashOperation struct {
	db *redis.Client
}

// NewHashOperation hash类型数据操作 https://www.tizi365.com/archives/296.html
func NewHashOperation(db *redis.Client) *HashOperation {
	return &HashOperation{db: db}
}

// Set 根据key和field字段设置，field字段的值
func (cl *HashOperation) Set(ctx context.Context, key string, value interface{}) *redis.IntCmd {
	return cl.db.HSet(ctx, key, value)
}

// Get 根据key和field字段设置，field字段的值
func (cl *HashOperation) Get(ctx context.Context, key, field string) *redis.StringCmd {
	return cl.db.HGet(ctx, key, field)
}
