package dorm

import (
	"errors"
	"fmt"
	"xorm.io/xorm"
)

type ConfigXormClient struct {
	Dns string // 地址
}

type XormClient struct {
	Db     *xorm.Engine      // 驱动
	config *ConfigXormClient // 配置
}

func NewXormMysqlClient(config *ConfigXormClient) (*XormClient, error) {

	var err error
	c := &XormClient{config: config}

	c.Db, err = xorm.NewEngine("mysql", c.config.Dns)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	return c, nil
}
