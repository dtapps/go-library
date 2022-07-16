package dorm

import (
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

func NewXormSqlite3Client(config *ConfigXormClient) (*XormClient, error) {

	var err error
	c := &XormClient{config: config}

	c.Db, err = xorm.NewEngine("sqlite3", c.config.Dns)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	return c, nil
}
