package gomysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type App struct {
	Db  *gorm.DB
	Dns string // 地址
}

func (app *App) InitClient() {

	log.Printf("gomysql：%+v\n", app)

	var err error

	app.Db, err = gorm.Open(mysql.Open(app.Dns), &gorm.Config{})

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
