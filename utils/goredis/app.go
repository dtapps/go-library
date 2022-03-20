package goredis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

// App 实例
type App struct {
	Db       *redis.Client
	Addr     string // 地址
	Password string // 密码
	DB       int    // 数据库
	PoolSize int    // 连接池大小
}

// InitClient 初始化连接
func (app *App) InitClient() {
	if app.PoolSize == 0 {
		app.PoolSize = 100
	}
	log.Printf("goredis：%+v\n", app)
	app.Db = redis.NewClient(&redis.Options{
		Addr:     app.Addr,     // 地址
		Password: app.Password, // 密码
		DB:       app.DB,       // 数据库
		PoolSize: app.PoolSize, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := app.Db.Ping(ctx).Result()
	if err != nil {
		panic(errors.New(fmt.Sprintf("goredis ping error：%v", err)))
	}
	return
}
