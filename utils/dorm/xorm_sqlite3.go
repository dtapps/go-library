package dorm

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

const XormSqlite3Driver = "sqlite3"

func NewXormSqlite3Client(ctx context.Context, config *XormClientConfig) (*XormClient, error) {

	var err error
	c := &XormClient{config: config}

	c.db, err = xorm.NewEngine(XormSqlite3Driver, c.config.Dns)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	return c, nil
}
