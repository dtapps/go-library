package kuaishou

import (
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gomongo"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

type App struct {
	ua           string
	mongo        *gomongo.Client // 日志数据库
	pgsql        *gorm.DB        // pgsql数据库
	client       *gorequest.App  // 请求客户端
	log          *golog.Api      // 日志服务
	logTableName string          // 日志表名
	logStatus    bool            // 日志状态
}

func NewApp(pgsql *gorm.DB) *App {
	app := &App{ua: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36"}
	app.client = gorequest.NewHttp()
	if pgsql != nil {
		app.pgsql = pgsql
		app.logStatus = true
		app.logTableName = "kuaishou"
		app.log = golog.NewApi(&golog.ApiConfig{
			Db:        pgsql,
			TableName: app.logTableName,
		})
	}
	return app
}
