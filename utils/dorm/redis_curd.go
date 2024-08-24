package dorm

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

// Set 设置一个key的值
func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.db.Set(ctx, key, value, expiration)
}

// Get 查询key的值
func (r *RedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.db.Get(ctx, key)
}

// GetSet 设置一个key的值，并返回这个key的旧值
func (r *RedisClient) GetSet(ctx context.Context, key string, value interface{}) *redis.StringCmd {
	return r.db.GetSet(ctx, key, value)
}

// SetNX 如果key不存在，则设置这个key的值
func (r *RedisClient) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	return r.db.SetNX(ctx, key, value, expiration)
}

// MGet 批量查询key的值
func (r *RedisClient) MGet(ctx context.Context, keys ...string) *redis.SliceCmd {
	return r.db.MGet(ctx, keys...)
}

// MSet 批量设置key的值
// MSet(map[string]interface{}{"key1": "value1", "key2": "value2"})
func (r *RedisClient) MSet(ctx context.Context, values map[string]interface{}) *redis.StatusCmd {
	return r.db.MSet(ctx, values)
}

// Incr 针对一个key的数值进行递增操作
func (r *RedisClient) Incr(ctx context.Context, key string) *redis.IntCmd {
	return r.db.Incr(ctx, key)
}

// IncrBy 针对一个key的数值进行递增操作，指定每次递增多少
func (r *RedisClient) IncrBy(ctx context.Context, key string, value int64) *redis.IntCmd {
	return r.db.IncrBy(ctx, key, value)
}

// Decr 针对一个key的数值进行递减操作
func (r *RedisClient) Decr(ctx context.Context, key string) *redis.IntCmd {
	return r.db.Decr(ctx, key)
}

// DecrBy 针对一个key的数值进行递减操作，指定每次递减多少
func (r *RedisClient) DecrBy(ctx context.Context, key string, value int64) *redis.IntCmd {
	return r.db.DecrBy(ctx, key, value)
}

// Del 删除key操作，支持批量删除
func (r *RedisClient) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return r.db.Del(ctx, keys...)
}

// Keys 按前缀获取所有key名
func (r *RedisClient) Keys(ctx context.Context, prefix string) []string {
	values, _ := r.db.Keys(ctx, prefix).Result()
	keys := make([]string, 0, len(values))
	if len(values) <= 0 {
		return keys
	}
	for _, value := range values {
		keys = append(keys, value)
	}
	return keys
}

// KeysValue 按前缀获取所有key值
func (r *RedisClient) KeysValue(ctx context.Context, prefix string) *redis.SliceCmd {
	values, _ := r.db.Keys(ctx, prefix).Result()
	if len(values) <= 0 {
		return &redis.SliceCmd{}
	}
	keys := make([]string, 0, len(values))
	for _, value := range values {
		keys = append(keys, value)
	}
	return r.MGet(ctx, keys...)
}

// AddKeyToSet 将值添加到集合
func (r *RedisClient) AddKeyToSet(ctx context.Context, setKey string, value interface{}) error {
	return r.db.SAdd(ctx, setKey, value).Err()
}

// GetAllKeysInSet 获取集合的所有元素
func (r *RedisClient) GetAllKeysInSet(ctx context.Context, setKey string) ([]string, error) {
	return r.db.SMembers(ctx, setKey).Result()
}

// DoesKeyExistInSet 检查值是否存在于集合中
func (r *RedisClient) DoesKeyExistInSet(ctx context.Context, setKey string, targetKey interface{}) (bool, error) {
	return r.db.SIsMember(ctx, setKey, targetKey).Result()
}

// RemoveKeyFromSet 从集合中删除指定的元素
func (r *RedisClient) RemoveKeyFromSet(ctx context.Context, setKey string, targetKey interface{}) error {
	return r.db.SRem(ctx, setKey, targetKey).Err()
}

// DeleteKeysWithPrefix 根据前缀删除key
func (r *RedisClient) DeleteKeysWithPrefix(ctx context.Context, prefix string) error {
	// 获取所有符合给定模式的键
	keys, err := r.db.Keys(ctx, prefix+"*").Result()
	if err != nil {
		return err
	}

	// 删除所有匹配的键
	if len(keys) > 0 {
		_, err := r.db.Del(ctx, keys...).Result()
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteScanWithPrefix 根据前缀删除key
func (r *RedisClient) DeleteScanWithPrefix(ctx context.Context, prefix string) error {
	var cursor uint64
	for {
		// 使用 SCAN 迭代获取匹配模式的键
		keys, nextCursor, err := r.db.Scan(ctx, cursor, prefix+"*", 10).Result()
		if err != nil {
			return err
		}

		// 删除当前迭代返回的键
		if len(keys) > 0 {
			_, err := r.db.Del(ctx, keys...).Result()
			if err != nil {
				return err
			}
		}

		// 更新游标，继续迭代
		cursor = nextCursor

		// 如果迭代结束，退出循环
		if cursor == 0 {
			break
		}
	}

	return nil
}
