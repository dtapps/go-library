package dorm

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
	"path"
	"time"
)

type ConfigGormClient struct {
	Dns    string // 地址
	Log    bool   // 日志
	LogUrl string // 日志路径
}

type GormClient struct {
	Db     *gorm.DB          // 驱动
	config *ConfigGormClient // 配置
}

func (c *GormClient) GetDb() *gorm.DB {
	return c.Db
}

type writer struct{}

// 日志路径
var logsUrl = ""

func (w writer) Printf(format string, args ...interface{}) {

	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + logsUrl
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	logFileName := "gorm." + now.Format("2006-01-02") + ".log"

	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	// 实例化
	l := logrus.New()

	// 设置输出
	l.Out = src

	// 设置日志格式 JSONFormatter=json TextFormatter=text
	l.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	l.Println(args...)
}
