package gopostgresql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Db  *gorm.DB
	Dns string
}

func (app *App) InitClient() {
	var err error

	app.Db, err = gorm.Open(postgres.Open(app.Dns), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("连接数据库失败：%v", err))
	}

	_, err = app.Db.DB()
	if err != nil {
		panic(fmt.Sprintf("连接数据库服务器失败：%v", err))
	}
}
