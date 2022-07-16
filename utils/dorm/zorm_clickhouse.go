package dorm

import (
	"errors"
	"fmt"
	"gitee.com/chunanyong/zorm"
	_ "github.com/mailru/go-clickhouse/v2"
)

func NewZormClickhouseClient(config *ConfigZormClient) (*ZormClient, error) {

	var err error
	c := &ZormClient{config: config}

	c.Db, err = zorm.NewDBDao(&zorm.DataSourceConfig{
		DSN:                   c.config.Dns,
		DriverName:            "chhttp",     // 数据库驱动名称
		DBType:                "clickhouse", // 数据库类型
		PrintSQL:              true,         // 是否打印sql
		MaxOpenConns:          0,            // 数据库最大连接数,默认50
		MaxIdleConns:          0,            // 数据库最大空闲连接数,默认50
		ConnMaxLifetimeSecond: 0,            // 连接存活秒时间. 默认600(10分钟)后连接被销毁重建.
		// 避免数据库主动断开连接,造成死连接.MySQL默认wait_timeout 28800秒(8小时)
		DefaultTxOptions: nil, // 事务隔离级别的默认配置,默认为nil
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	return c, nil
}
