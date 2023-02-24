package dorm

import (
	"errors"
	"fmt"
	"gitee.com/chunanyong/zorm"
	_ "github.com/bmizerany/pq"
)

func NewZormPostgresqlClient(config *ConfigZormClient) (*ZormClient, error) {

	var err error
	c := &ZormClient{config: config}

	c.db, err = zorm.NewDBDao(&zorm.DataSourceConfig{
		DSN:        c.config.Dns,
		DriverName: "postgres",   // 数据库驱动名称
		Dialect:    "postgresql", // 数据库类型
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	return c, nil
}
