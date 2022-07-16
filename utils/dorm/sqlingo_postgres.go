package dorm

import (
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/lqs/sqlingo"
)

func NewSqLiNgoPostgresClient(config *ConfigSqLiNgoClient) (*SqLiNgoClient, error) {

	var err error
	c := &SqLiNgoClient{config: config}

	c.Db, err = sqlingo.Open("postgres", c.config.Dns)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	return c, nil
}
