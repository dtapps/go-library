package gocache

import (
	"github.com/bradfitz/gomemcache/memcache"
)

// MemcachedConfig 配置
type MemcachedConfig struct {
	Dns string           // 连接地址，可选
	Db  *memcache.Client // 驱动，可选
}

// Memcached https://github.com/bradfitz/gomemcache
type Memcached struct {
	db *memcache.Client // 驱动
}

// NewMemcached 实例化
func NewMemcached(config *MemcachedConfig) *Memcached {
	if config.Dns == "" {
		return &Memcached{db: config.Db}
	} else {
		mc := memcache.New(config.Dns)
		if mc == nil {
			panic("连接失败")
		}
		return &Memcached{db: mc}
	}
}

// Set 插入数据
func (m *Memcached) Set(key string, value []byte) error {
	return m.db.Set(&memcache.Item{Key: key, Value: value})
}

// Get 获取单个数据
func (m *Memcached) Get(key string) (string, error) {
	it, err := m.db.Get(key)
	if err == memcache.ErrCacheMiss {
		return "", memcache.ErrCacheMiss
	}
	if it.Key == key {
		return string(it.Value), nil
	}
	return "", memcache.ErrCacheMiss
}

// Del 删除单个数据
func (m *Memcached) Del(key string) error {
	return m.db.Delete(key)
}
