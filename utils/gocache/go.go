package gocache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

// GoConfig 配置
type GoConfig struct {
	DefaultExpiration time.Duration // 默认过期时间
	DefaultClear      time.Duration // 清理过期数据
}

// Go https://github.com/patrickmn/go-cache
type Go struct {
	config *GoConfig
	db     *cache.Cache // 驱动
}

// NewGo 实例化
func NewGo(config *GoConfig) *Go {
	c := &Go{config: config}
	c.db = cache.New(c.config.DefaultExpiration, c.config.DefaultClear)
	return c
}

// Set 插入数据 并设置过期时间
func (c *Go) Set(key string, value interface{}, expirationTime time.Duration) {
	c.db.Set(key, value, expirationTime)
}

// Get 获取单个数据
func (c *Go) Get(key string) (interface{}, bool) {
	return c.db.Get(key)
}

// SetDefault 插入数据 并设置为默认过期时间
func (c *Go) SetDefault(key string, value interface{}) {
	c.db.Set(key, value, c.config.DefaultExpiration)
}
