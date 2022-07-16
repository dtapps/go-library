package dorm

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lesismal/sqlw"
)

func NewSqlWMysqlClient(config *ConfigSqlWClient) (*SqlWClient, error) {

	var err error
	c := &SqlWClient{config: config}

	c.Db, err = sqlw.Open("mysql", c.config.Dns, "db")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	return c, nil
}
