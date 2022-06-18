package dingdanxia

import (
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gomongo"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

type App struct {
	apiKey       string          // 密钥
	mongo        *gomongo.Client // 日志数据库
	pgsql        *gorm.DB        // pgsql数据库
	client       *gorequest.App  // 请求客户端
	log          *golog.Api      // 日志服务
	logTableName string          // 日志表名
	logStatus    bool            // 日志状态
}

// NewApp 创建实例
func NewApp(apiKey string, pgsql *gorm.DB) *App {
	app := &App{apiKey: apiKey}
	app.client = gorequest.NewHttp()
	if pgsql != nil {
		app.pgsql = pgsql
		app.logStatus = true
		app.logTableName = "dingdanxia"
		app.log = golog.NewApi(&golog.ApiConfig{
			Db:        pgsql,
			TableName: app.logTableName,
		})
	}
	return app
}

// 请求接口
func (app *App) request(url string, params map[string]interface{}, method string) (resp gorequest.Response, err error) {

	// 公共参数
	params["apikey"] = app.apiKey

	// 创建请求
	client := app.client

	// 设置请求地址
	client.SetUri(url)

	// 设置方式
	client.SetMethod(method)

	// 设置格式
	client.SetContentTypeForm()

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Request()
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if app.logStatus == true {
		go app.postgresqlLog(request)
	}

	return request, err
}
