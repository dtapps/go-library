package gomongodb

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// App 实例
type App struct {
	Mgo        *mongo.Client
	User       string // 用户名
	Password   string // 密码
	Host       string // 地址
	Port       int    // 端口
	Database   string // 数据库
	Collection string // 表名
}

// InitClient 初始化连接
func (app *App) InitClient() {
	var err error
	// 连接到MongoDB
	app.Mgo, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d",
		app.User,
		app.Password,
		app.Host,
		app.Port,
	)))
	if err != nil {
		panic(errors.New(fmt.Sprintf("gomongodb connect error：%v", err)))
	}
	// 检查连接
	err = app.Mgo.Ping(context.TODO(), nil)
	if err != nil {
		panic(errors.New(fmt.Sprintf("gomongodb ping error：%v", err)))
	}
	return
}

// Close 关闭
func (app *App) Close() {
	err := app.Mgo.Disconnect(context.TODO())
	if err != nil {
		panic(errors.New(fmt.Sprintf("gomongodb close error：%v", err)))
	}
	return
}
