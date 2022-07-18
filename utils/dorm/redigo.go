package dorm

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type ConfigRedigoClient struct {
	Addr     string // 地址
	Password string // 密码
	DB       int    // 数据库
	PoolSize int    // 连接池大小
}

// RedigoClient
// https://Redigo.uptrace.dev/
type RedigoClient struct {
	Db     *redis.Conn         // 驱动
	config *ConfigRedigoClient // 配置
}

func NewRedigoClient(config *ConfigRedigoClient) (*RedigoClient, error) {

	c := &RedigoClient{}
	c.config = config
	if c.config.PoolSize == 0 {
		c.config.PoolSize = 100
	}

	db, err := redis.Dial("tcp", c.config.Addr, redis.DialPassword(c.config.Password), redis.DialDatabase(c.config.DB))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}
	c.Db = &db

	return c, nil
}
