package gojobs

import (
	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// GetDb 获取数据库驱动
func (c *Client) GetDb() *gorm.DB {
	return c.gormClient.GetDb()
}

// GetMongoDb 获取数据库驱动
func (c *Client) GetMongoDb() *mongo.Client {
	return c.mongoClient.GetDb()
}

// GetRedis 获取缓存数据库驱动
func (c *Client) GetRedis() *redis.Client {
	return c.cache.redisClient.GetDb()
}

// GetCurrentIp 获取当前ip
func (c *Client) GetCurrentIp() string {
	return c.config.systemOutsideIp
}

// GetSubscribeAddress 获取订阅地址
func (c *Client) GetSubscribeAddress() string {
	return c.cache.cornKeyPrefix + "_" + c.cache.cornKeyCustom
}
