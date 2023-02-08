package dorm

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// Subscribe 订阅channel
func (r *RedisClient) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return r.db.Subscribe(ctx, channels...)
}

// PSubscribe 订阅channel支持通配符匹配
func (r *RedisClient) PSubscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return r.db.PSubscribe(ctx, channels...)
}

// Publish 将信息发送到指定的channel
func (r *RedisClient) Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd {
	return r.db.Publish(ctx, channel, message)
}

// PubSubChannels 查询活跃的channel
func (r *RedisClient) PubSubChannels(ctx context.Context, pattern string) *redis.StringSliceCmd {
	return r.db.PubSubChannels(ctx, pattern)
}

// PubSubNumSub 查询指定的channel有多少个订阅者
func (r *RedisClient) PubSubNumSub(ctx context.Context, channels ...string) *redis.MapStringIntCmd {
	return r.db.PubSubNumSub(ctx, channels...)
}
