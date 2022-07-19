package dorm

import "github.com/bradfitz/gomemcache/memcache"

type ConfigMemcacheClient struct {
	Dns string
}

type MemcacheClient struct {
	Db     *memcache.Client      // 驱动
	config *ConfigMemcacheClient // 配置
}

func NewMemcacheClient(config *ConfigMemcacheClient) (*MemcacheClient, error) {

	c := &MemcacheClient{config: config}

	c.Db = memcache.New(c.config.Dns)
	if c.Db != nil {
		panic("memcache New failed")
	}

	return c, nil
}
