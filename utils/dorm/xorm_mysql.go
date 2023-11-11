package dorm

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

const XormMysqlDriver = "mysql"

func NewXormMysqlClient(config *XormClientConfig) (*XormClient, error) {

	var err error
	c := &XormClient{config: config}

	c.db, err = xorm.NewEngine(XormMysqlDriver, c.config.Dns)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	return c, nil
}
