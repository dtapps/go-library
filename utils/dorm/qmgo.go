package dorm

import (
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/qmgo"
)

type ConfigQMgoClient struct {
	Dns          string
	DatabaseName string // 库名
}

type QMgoClient struct {
	Db     *qmgo.Client
	config *ConfigQMgoClient // 配置
}

func NewQgoClient(config *ConfigQMgoClient) (*QMgoClient, error) {

	var err error
	c := &QMgoClient{config: config}

	c.Db, err = qmgo.NewClient(context.Background(), &qmgo.Config{Uri: c.config.Dns})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("检查连接失败：%v", err))
	}

	return c, nil
}
