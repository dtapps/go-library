package dorm

import (
	"errors"
	"fmt"
	"github.com/dtapps/go-library/utils/gotime"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func NewGormPostgresClient(config *ConfigGormClient) (*GormClient, error) {

	c := &GormClient{}
	c.config = config

	// 判断路径
	if c.config.LogUrl == "" {
		logsUrl = "/logs/postgresql"
	}

	var err error

	if c.config.Log == true {
		c.Db, err = gorm.Open(postgres.Open(c.config.Dns), &gorm.Config{
			Logger: logger.New(
				writer{},
				logger.Config{
					SlowThreshold:             time.Second, // 慢 SQL 阈值
					LogLevel:                  logger.Info, // 日志级别
					IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
					Colorful:                  false,       // 禁用彩色打印
				},
			),
			NowFunc: func() time.Time {
				return gotime.Current().Now().Local()
			},
		})
	} else {
		c.Db, err = gorm.Open(postgres.Open(c.config.Dns), &gorm.Config{})
	}

	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	sqlDB, err := c.Db.DB()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("检查连接失败：%v", err))
	}

	sqlDB.SetMaxIdleConns(10)                   // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)                  // 设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(time.Second * 600) // 设置了连接可复用的最大时间。

	return c, nil
}

func NewGormPostgresqlClient(config *ConfigGormClient) (*GormClient, error) {

	c := &GormClient{}
	c.config = config

	// 判断路径
	if c.config.LogUrl == "" {
		logsUrl = "/logs/postgresql"
	}

	var err error

	if c.config.Log == true {
		c.Db, err = gorm.Open(postgres.Open(c.config.Dns), &gorm.Config{
			Logger: logger.New(
				writer{},
				logger.Config{
					SlowThreshold:             time.Second, // 慢 SQL 阈值
					LogLevel:                  logger.Info, // 日志级别
					IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
					Colorful:                  false,       // 禁用彩色打印
				},
			),
			NowFunc: func() time.Time {
				return gotime.Current().Now().Local()
			},
		})
	} else {
		c.Db, err = gorm.Open(postgres.Open(c.config.Dns), &gorm.Config{})
	}

	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	sqlDB, err := c.Db.DB()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("检查连接失败：%v", err))
	}

	sqlDB.SetMaxIdleConns(10)                   // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)                  // 设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(time.Second * 600) // 设置了连接可复用的最大时间。

	return c, nil
}
