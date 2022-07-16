package dorm

import (
	"errors"
	"fmt"
	"github.com/lqs/sqlingo"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqLiNgoSqliteClient(config *ConfigSqLiNgoClient) (*SqLiNgoClient, error) {

	var err error
	c := &SqLiNgoClient{config: config}

	c.Db, err = sqlingo.Open("sqlite", c.config.Dns)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	return c, nil
}
