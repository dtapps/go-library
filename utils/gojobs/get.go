package gojobs

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"strings"
)

// GetDb 获取数据库驱动
func (c *Client) GetDb() *gorm.DB {
	return c.gormConfig.client
}

// GetGormDb 获取数据库驱动
func (c *Client) GetGormDb() *gorm.DB {
	return c.gormConfig.client
}

// GetRedisDb 获取缓存数据库驱动
func (c *Client) GetRedisDb() *redis.Client {
	return c.redisConfig.client
}

// GetCurrentIp 获取当前IP
func (c *Client) GetCurrentIp() string {
	return c.config.systemOutsideIP
}

// GetSubscribeAddress 获取订阅地址
func (c *Client) GetSubscribeAddress() string {
	var builder strings.Builder
	builder.WriteString(c.redisConfig.cornKeyPrefix)
	builder.WriteString("_")
	builder.WriteString(c.redisConfig.cornKeyCustom)
	return builder.String()
}
