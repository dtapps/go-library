package dorm

import (
	"errors"
	"fmt"
	"github.com/daodao97/fly"
	_ "github.com/go-sql-driver/mysql"
)

func NewFlyMysqlClient(config *ConfigFlyClient) (*FlyClient, error) {

	var err error
	c := &FlyClient{config: config}

	err = fly.Init(map[string]*fly.Config{
		"default": {DSN: c.config.Dns},
	})

	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	return c, nil
}
