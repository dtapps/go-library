package dorm

import (
	"errors"
	"fmt"
	"github.com/go-rel/mysql"
	"github.com/go-rel/rel"
	_ "github.com/go-sql-driver/mysql"
)

func NewRelMysqlClient(config *ConfigRelClient) (*RelClient, error) {

	var err error
	c := &RelClient{config: config}

	adapter, err := mysql.Open(c.config.Dns)
	defer adapter.Close()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	repo := rel.New(adapter)

	c.Db = &repo

	return c, nil
}
