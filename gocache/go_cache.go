package gocache

import (
	"time"
)

// GoCacheConfig 配置
type GoCacheConfig struct {
	expiration time.Duration // 过期时间
}

// GoCache https://github.com/patrickmn/go-cache
type GoCache struct {
	GoCacheConfig
	db              *Go              // 驱动
	GetterInterface GttInterfaceFunc // 不存在的操作
}

// NewCache 实例化
func (c *Go) NewCache(config *GoCacheConfig) *GoCache {
	app := &GoCache{}
	app.expiration = config.expiration
	app.db = c
	return app
}

// GetInterface 缓存操作
func (gc *GoCache) GetInterface(key string) (ret interface{}) {

	f := func() interface{} {
		return gc.GetterInterface()
	}

	// 如果不存在，则调用GetterInterface
	ret, found := gc.db.Get(key)

	if found == false {
		gc.db.Set(key, f(), gc.expiration)
		ret, _ = gc.db.Get(key)
	}

	return
}
