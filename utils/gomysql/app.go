package gomysql

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path"
	"time"
)

type App struct {
	Db  *gorm.DB // 驱动
	Dns string   // 地址
	Log bool     // 日志
}

type writer struct{}

func (w writer) Printf(format string, args ...interface{}) {

	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/mysql"
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

func (app *App) InitClient() {

	log.Printf("gomysql：%+v\n", app)

	var err error

	if app.Log == true {
		app.Db, err = gorm.Open(mysql.Open(app.Dns), &gorm.Config{
			Logger: logger.New(
				writer{},
				logger.Config{
					SlowThreshold:             time.Second, // 慢 SQL 阈值
					LogLevel:                  logger.Info, // 日志级别
					IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
					Colorful:                  false,       // 禁用彩色打印
				},
			),
		})
	} else {
		app.Db, err = gorm.Open(mysql.Open(app.Dns), &gorm.Config{})
	}

	if err != nil {
		panic(fmt.Sprintf("连接数据库失败：%v", err))
	}

	sqlDB, err := app.Db.DB()
	if err != nil {
		panic(fmt.Sprintf("连接数据库服务器失败：%v", err))
	}
	sqlDB.SetMaxIdleConns(10)                   // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)                  // 设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(time.Second * 600) // 设置了连接可复用的最大时间。

}
