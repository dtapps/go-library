package dorm

import (
	"errors"
	"fmt"
	"github.com/dtapps/go-library/utils/gotime"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func NewGormMysqlClient(config *ConfigGormClient) (*GormClient, error) {

	var err error
	c := &GormClient{config: config}

	// 判断路径
	if c.config.Log.Path == "" {
		logsUrl = "/logs/mysql"
	} else {
		logsUrl = c.config.Log.Path
	}

	if c.config.Log.Status == true {
		var slowThreshold time.Duration
		var logLevel logger.LogLevel
		if c.config.Log.Slow == 0 {
			slowThreshold = 100 * time.Millisecond
		} else {
			slowThreshold = time.Duration(c.config.Log.Slow)
		}
		if c.config.Log.Level == "Error" {
			logLevel = logger.Error
		} else if c.config.Log.Level == "Warn" {
			logLevel = logger.Warn
		} else {
			logLevel = logger.Info
		}
		c.Db, err = gorm.Open(mysql.Open(c.config.Dns), &gorm.Config{
			Logger: logger.New(
				writer{},
				logger.Config{
					SlowThreshold:             slowThreshold,              // 慢SQL阈值
					LogLevel:                  logLevel,                   // 日志级别
					IgnoreRecordNotFoundError: c.config.Log.NotFoundError, // 忽略ErrRecordNotFound（记录未找到）错误
					Colorful:                  c.config.Log.Colorful,      // 禁用彩色打印
				},
			),
			NowFunc: func() time.Time {
				return gotime.Current().Now().Local()
			},
		})
	} else {
		c.Db, err = gorm.Open(mysql.Open(c.config.Dns), &gorm.Config{})
	}

	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	sqlDB, err := c.Db.DB()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("检查连接失败：%v", err))
	}

	// 设置空闲连接池中连接的最大数量
	if c.config.Conn.SetMaxIdle == 0 {
		sqlDB.SetMaxIdleConns(10)
	} else {
		sqlDB.SetMaxIdleConns(c.config.Conn.SetMaxIdle)
	}

	// 设置打开数据库连接的最大数量
	if c.config.Conn.SetMaxOpen == 0 {
		sqlDB.SetMaxOpenConns(100)
	} else {
		sqlDB.SetMaxOpenConns(c.config.Conn.SetMaxOpen)
	}

	// 设置了连接可复用的最大时间
	if c.config.Conn.SetConnMaxLifetime == 0 {
		sqlDB.SetConnMaxLifetime(time.Second * 600)
	} else {
		sqlDB.SetConnMaxLifetime(time.Duration(c.config.Conn.SetConnMaxLifetime))
	}

	return c, nil
}
