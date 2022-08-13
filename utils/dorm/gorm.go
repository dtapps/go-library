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
	Dns string // 地址
	Log struct {
		Status        bool   // 状态
		Path          string // 路径
		Slow          int64  // 慢SQL阈值
		Level         string // 级别
		NotFoundError bool   // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful      bool   // 禁用彩色打印
	} // 日志
	Conn struct {
		SetMaxIdle         int   // 设置空闲连接池中连接的最大数量
		SetMaxOpen         int   // 设置打开数据库连接的最大数量
		SetConnMaxLifetime int64 // 设置了连接可复用的最大时间
	} // 连接
}

// GormClient
// https://gorm.io/
type GormClient struct {
	Db     *gorm.DB          // 驱动
	config *ConfigGormClient // 配置
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
