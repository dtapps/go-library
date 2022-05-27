package gocache

import (
	"github.com/allegro/bigcache/v3"
)

// BigCache https://github.com/allegro/bigcache
type BigCache struct {
	db              *Big             // 驱动
	GetterInterface GttInterfaceFunc // 不存在的操作
}

// NewCache 实例化组件
func (c *Big) NewCache() *BigCache {
	return &BigCache{db: c}
}

// GetInterface 缓存操作
func (gc *BigCache) GetInterface(key string) (ret interface{}) {

	f := func() interface{} {
		return gc.GetterInterface()
	}

	// 如果不存在，则调用GetterInterface
	ret, err := gc.db.Get(key)

	if err == bigcache.ErrEntryNotFound {
		_ = gc.db.Set(key, f())
		ret, _ = gc.db.Get(key)
	}
	return
}
