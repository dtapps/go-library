package dorm

import (
	"context"
	"github.com/go-redis/redis/v9"
)

type ListOperation struct {
	db  *redis.Client
	ctx context.Context
}

// NewListOperation 列表(list)类型数据操作 https://www.tizi365.com/archives/299.html
func (r *RedisClient) NewListOperation() *ListOperation {
	return &ListOperation{db: r.Db, ctx: context.Background()}
}

// LPush 从列表左边插入数据
func (cl *ListOperation) LPush(key string, value interface{}) *redis.IntCmd {
	return cl.db.LPush(cl.ctx, key, value)
}

// LPushX 跟LPush的区别是，仅当列表存在的时候才插入数据
func (cl *ListOperation) LPushX(key string, value interface{}) *redis.IntCmd {
	return cl.db.LPushX(cl.ctx, key, value)
}

// RPop 从列表的右边删除第一个数据，并返回删除的数据
func (cl *ListOperation) RPop(key string) *redis.StringCmd {
	return cl.db.RPop(cl.ctx, key)
}

// RPush 从列表右边插入数据
func (cl *ListOperation) RPush(key string, value interface{}) *redis.IntCmd {
	return cl.db.RPush(cl.ctx, key, value)
}

// RPushX 跟RPush的区别是，仅当列表存在的时候才插入数据
func (cl *ListOperation) RPushX(key string, value interface{}) *redis.IntCmd {
	return cl.db.RPushX(cl.ctx, key, value)
}

// LPop 从列表左边删除第一个数据，并返回删除的数据
func (cl *ListOperation) LPop(key string) *redis.StringCmd {
	return cl.db.LPop(cl.ctx, key)
}

// Len 返回列表的大小
func (cl *ListOperation) Len(key string) *redis.IntCmd {
	return cl.db.LLen(cl.ctx, key)
}

// Range 返回列表的一个范围内的数据，也可以返回全部数据
func (cl *ListOperation) Range(key string, start, stop int64) *redis.StringSliceCmd {
	return cl.db.LRange(cl.ctx, key, start, stop)
}

// RangeAli 返回key全部数据
func (cl *ListOperation) RangeAli(key string) *redis.StringSliceCmd {
	return cl.db.LRange(cl.ctx, key, 0, -1)
}

// Rem 删除key中的数据
func (cl *ListOperation) Rem(key string, count int64, value interface{}) *redis.IntCmd {
	return cl.db.LRem(cl.ctx, key, count, value)
}

// Index 根据索引坐标，查询key中的数据
func (cl *ListOperation) Index(key string, index int64) *redis.StringCmd {
	return cl.db.LIndex(cl.ctx, key, index)
}

// Insert 在指定位置插入数据
func (cl *ListOperation) Insert(key, op string, pivot, value interface{}) *redis.IntCmd {
	return cl.db.LInsert(cl.ctx, key, op, pivot, value)
}
