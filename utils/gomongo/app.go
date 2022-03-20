package gomongo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// App 实例
type App struct {
	Db         *mongo.Client
	User       string // 用户名
	Password   string // 密码
	Host       string // 地址
	Port       int    // 端口
	Dbname     string // 数据库
	Dns        string // 地址
	collection string // 表名
}

// InitClient 初始化连接
func (app *App) InitClient() {
	log.Printf("gomongo：%+v\n", app)
	var err error
	// 连接到MongoDB
	app.Db, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(app.Dns))
	if err != nil {
		panic(errors.New(fmt.Sprintf("gomongodb connect error：%v", err)))
	}
	// 检查连接
	err = app.Db.Ping(context.TODO(), nil)
	if err != nil {
		panic(errors.New(fmt.Sprintf("gomongodb ping error：%v", err)))
	}
	return
}

// Close 关闭
func (app *App) Close() {
	err := app.Db.Disconnect(context.TODO())
	if err != nil {
		panic(errors.New(fmt.Sprintf("gomongodb close error：%v", err)))
	}
	return
}
