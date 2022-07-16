package dorm

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func NewXormMysqlClient(config *ConfigXormClient) (*XormClient, error) {

	var err error
	c := &XormClient{config: config}

	c.Db, err = xorm.NewEngine("mysql", c.config.Dns)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	return c, nil
}
