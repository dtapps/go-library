package bigcache_cache

import (
	"context"
	"encoding/json"
	"github.com/allegro/bigcache/v3"
	"time"
)

// GetStringFunc String缓存结构
type GetStringFunc func() string

// GetAnyFunc Any缓存结构
type GetAnyFunc func() any

// CacheClient 缓存客户端结构 https://github.com/allegro/bigcache
type CacheClient struct {
	defaultExpiration time.Duration      // 过期时间
	operation         *bigcache.BigCache // 操作
}

// NewClient 实例化
func NewClient(ctx context.Context, expiration time.Duration) (*CacheClient, error) {
	client, err := bigcache.New(ctx, bigcache.DefaultConfig(expiration))
	return &CacheClient{
		defaultExpiration: expiration,
		operation:         client,
	}, err
}

// GetString 缓存操作
func (bc *CacheClient) GetString(ctx context.Context, key string, callback GetStringFunc) (ret string, err error) {
	// 尝试从缓存中获取值
	data, err := bc.operation.Get(key)
	if err == nil {
		return string(data), nil
	}

	// 不存在调用 GetStringFunc
	value := callback()
	// 设置缓存值
	err = bc.operation.Set(key, []byte(value))
	return value, err
}

// GetAny 缓存操作
func (bc *CacheClient) GetAny(ctx context.Context, key string, resultValue any, callback GetAnyFunc) (err error) {
	// 尝试从缓存中获取值
	data, err := bc.operation.Get(key)
	if err == nil {
		return json.Unmarshal(data, resultValue)
	}

	// 不存在调用 GetAnyFunc
	value := callback()
	// 序列化
	marshal, err := json.Marshal(value)
	if err != nil {
		return err
	}
	// 设置缓存值
	err = bc.operation.Set(key, marshal)
	if err != nil {
		return err
	}
	return json.Unmarshal(marshal, resultValue)
}

// SetAnyKey 设置缓存值
func (bc *CacheClient) SetAnyKey(ctx context.Context, key string, value any) error {
	marshal, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return bc.operation.Set(key, marshal)
}

// GetAnyKey 获取缓存值
func (bc *CacheClient) GetAnyKey(ctx context.Context, key string, result any) error {
	data, err := bc.operation.Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
}
