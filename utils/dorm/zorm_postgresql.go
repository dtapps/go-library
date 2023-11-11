package dorm

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/chunanyong/zorm"
	_ "github.com/bmizerany/pq"
)

const ZormPostgresqlDriver = "postgresql"

func NewZormPostgresqlClient(ctx context.Context, config *ZormClientConfig) (*ZormClient, error) {

	var err error
	c := &ZormClient{config: config}

	dbConfig := &zorm.DataSourceConfig{
		// DSN 数据库的连接字符串,parseTime=true会自动转换为time格式,默认查询出来的是[]byte数组.&loc=Local用于设置时区
		DSN: c.config.Dns,
		// sql.Open(DriverName,DSN) DriverName就是驱动的sql.Open第一个字符串参数,根据驱动实际情况获取
		DriverName:            ZormPostgresDriver,   // 数据库驱动名称
		Dialect:               ZormPostgresqlDriver, // 数据库类型
		MaxOpenConns:          0,                    // 数据库最大连接数,默认50
		MaxIdleConns:          0,                    // 数据库最大空闲连接数,默认50
		ConnMaxLifetimeSecond: 0,                    // 连接存活秒时间. 默认600(10分钟)后连接被销毁重建.
		// 避免数据库主动断开连接,造成死连接.MySQL默认wait_timeout 28800秒(8小时)
		DefaultTxOptions: nil, // 事务隔离级别的默认配置,默认为nil
	}
	c.db, err = zorm.NewDBDao(dbConfig)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	return c, nil
}
