package bigcache_cache

import (
	"context"
	"encoding/json"
	"github.com/allegro/bigcache/v3"
	"time"
)

// GetStringFunc String缓存结构
type GetStringFunc func() (string, error)

// GetAnyFunc Any缓存结构
type GetAnyFunc func() (any, error)

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
	data, err := bc.Get(ctx, key)
	if err == nil {
		return string(data), nil
	}

	// 不存在调用 GetStringFunc
	value, err := callback()
	if err != nil {
		return value, err
	}

	// 设置缓存值
	return value, bc.Set(ctx, key, []byte(value))
}

// GetAny 缓存操作
func (bc *CacheClient) GetAny(ctx context.Context, key string, resultValue any, callback GetAnyFunc) (err error) {
	// 尝试从缓存中获取值
	data, err := bc.Get(ctx, key)
	if err == nil {
		return json.Unmarshal(data, resultValue)
	}

	// 不存在调用 GetAnyFunc
	value, err := callback()
	if err != nil {
		return err
	}

	// 序列化
	marshal, err := json.Marshal(value)
	if err != nil {
		return err
	}
	// 设置缓存值
	err = bc.Set(ctx, key, marshal)
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
	return bc.Set(ctx, key, marshal)
}

// GetAnyKey 获取缓存值
func (bc *CacheClient) GetAnyKey(ctx context.Context, key string, result any) error {
	data, err := bc.Get(ctx, key)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
}

// Get 获取缓存
func (bc *CacheClient) Get(ctx context.Context, key string) ([]byte, error) {
	return bc.operation.Get(key)
}

// Set 设置缓存
func (bc *CacheClient) Set(ctx context.Context, key string, entry []byte) error {
	return bc.operation.Set(key, entry)
}

// Delete 删除缓存
func (bc *CacheClient) Delete(ctx context.Context, key string) error {
	return bc.operation.Delete(key)
}

func (bc *CacheClient) GetWithInfo(ctx context.Context, key string) ([]byte, bigcache.Response, error) {
	return bc.operation.GetWithInfo(key)
}

func (bc *CacheClient) Append(ctx context.Context, key string, entry []byte) error {
	return bc.operation.Append(key, entry)
}

func (bc *CacheClient) Reset(ctx context.Context) error {
	return bc.operation.Reset()
}

func (bc *CacheClient) ResetStats(ctx context.Context) error {
	return bc.operation.ResetStats()
}

func (bc *CacheClient) Len(ctx context.Context) int {
	return bc.operation.Len()
}

func (bc *CacheClient) Capacity(ctx context.Context) int {
	return bc.operation.Capacity()
}

func (bc *CacheClient) Stats(ctx context.Context) bigcache.Stats {
	return bc.operation.Stats()
}

func (bc *CacheClient) KeyMetadata(ctx context.Context, key string) bigcache.Metadata {
	return bc.operation.KeyMetadata(key)
}

func (bc *CacheClient) Iterator(ctx context.Context) *bigcache.EntryInfoIterator {
	return bc.operation.Iterator()
}
