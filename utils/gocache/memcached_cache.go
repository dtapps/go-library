package gocache

import (
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
)

// MemcachedCache https://github.com/bradfitz/gomemcache
type MemcachedCache struct {
	db              *Memcached       // 驱动
	GetterString    GttStringFunc    // 不存在的操作
	GetterInterface GttInterfaceFunc // 不存在的操作
}

// NewCache 实例化
func (m *Memcached) NewCache() *MemcachedCache {
	return &MemcachedCache{db: m}
}

// GetString 缓存操作
func (mc *MemcachedCache) GetString(key string) (ret string) {

	f := func() string {
		return mc.GetterString()
	}

	// 如果不存在，则调用GetterString
	ret, err := mc.db.Get(key)

	if err == memcache.ErrCacheMiss {
		mc.db.Set(key, []byte(f()))
		ret, _ = mc.db.Get(key)
	}

	return
}

// GetInterface 缓存操作
func (mc *MemcachedCache) GetInterface(key string, result interface{}) {

	f := func() string {
		marshal, _ := json.Marshal(mc.GetterInterface())
		return string(marshal)
	}

	// 如果不存在，则调用GetterInterface
	ret, err := mc.db.Get(key)

	if err == memcache.ErrCacheMiss {
		mc.db.Set(key, []byte(f()))
		ret, _ = mc.db.Get(key)
	}

	err = json.Unmarshal([]byte(ret), result)

	return
}
