package dorm

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
	"path"
	"time"
)

// GormClientFun *GormClient 驱动
type GormClientFun func() *GormClient

// GormClientTableFun *GormClient 驱动
// string 表名
type GormClientTableFun func() (*GormClient, string)

type GormClientConfig struct {
	Dns                    string // 地址
	LogStatus              bool   // 日志 - 状态
	LogPath                string // 日志 - 路径
	LogSlow                int64  // 日志 - 慢SQL阈值
	LogLevel               string // 日志 - 级别
	ConnSetMaxIdle         int    // 连接 - 设置空闲连接池中连接的最大数量
	ConnSetMaxOpen         int    // 连接 - 设置打开数据库连接的最大数量
	ConnSetConnMaxLifetime int64  // 连接 - 设置了连接可复用的最大时间
}

// GormClient
// https://gorm.io/zh_CN/docs/index.html
type GormClient struct {
	Db     *gorm.DB          // 驱动
	config *GormClientConfig // 配置
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
