package pconline

import (
	"go.dtapp.net/library/golog"
	golog2 "go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gomongo"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

type App struct {
	mongo        *gomongo.Client // 日志数据库
	pgsql        *gorm.DB        // pgsql数据库
	client       *gorequest.App  // 请求客户端
	log          *golog2.Api     // 日志服务
	logTableName string          // 日志表名
	logStatus    bool            // 日志状态
}

func NewApp(pgsql *gorm.DB) *App {
	app := &App{}
	app.client = gorequest.NewHttp()
	if pgsql != nil {
		app.pgsql = pgsql
		app.logStatus = true
		app.logTableName = "pconline"
		app.log = golog.NewClientApi(pgsql, app.logTableName)
	}
	return app
}

func (app *App) request(url string) (resp gorequest.Response, err error) {

	// 创建请求
	client := app.client

	// 设置请求地址
	client.SetUri(url)

	// 发起请求
	request, err := client.Get()
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if app.mongo != nil && app.mongo.Db != nil {
		go app.mongoLog(request)
	}
	if app.logStatus == true {
		go app.postgresqlLog(request)
	}

	return request, err
}
