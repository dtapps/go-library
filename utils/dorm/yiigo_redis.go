package dorm

import (
	"github.com/shenghui0779/yiigo"
	"time"
)

func NewYiiGoRedisClient(config *ConfigYiiGoClient) (*YiiGoClient, error) {

	c := &YiiGoClient{config: config}

	yiigo.Init(
		yiigo.WithRedis(yiigo.Default, &yiigo.RedisConfig{
			Addr: c.config.Addr,
			Options: &yiigo.RedisOptions{
				ConnTimeout:  10 * time.Second,
				ReadTimeout:  10 * time.Second,
				WriteTimeout: 10 * time.Second,
				PoolSize:     10,
				IdleTimeout:  5 * time.Minute,
			},
		}),
	)

	return c, nil
}
