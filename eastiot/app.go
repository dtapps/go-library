package eastiot

import (
	"go.dtapp.net/library/golog"
	"go.dtapp.net/library/gomongo"
	"go.dtapp.net/library/gorequest"
	"gorm.io/gorm"
	"time"
)

type App struct {
	appId        string
	apiKey       string
	mongo        *gomongo.Client // 日志数据库
	pgsql        *gorm.DB        // pgsql数据库
	client       *gorequest.App  // 请求客户端
	log          *golog.Api      // 日志服务
	logTableName string          // 日志表名
	logStatus    bool            // 日志状态
}

func NewApp(appID string, apiKey string, pgsql *gorm.DB) *App {
	app := &App{appId: appID, apiKey: apiKey}
	app.client = gorequest.NewHttp()
	if pgsql != nil {
		app.pgsql = pgsql
		app.logStatus = true
		app.logTableName = "eastiot"
		app.log = golog.NewApi(&golog.ApiConfig{
			Db:        pgsql,
			TableName: app.logTableName,
		})
	}
	return app
}

func (app *App) request(url string, params map[string]interface{}, method string) (resp gorequest.Response, err error) {

	// 公共参数
	params["appId"] = app.appId
	params["timeStamp"] = time.Now().Unix()

	// 签名
	params["sign"] = app.getSign(app.apiKey, params)

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
	if app.mongo != nil && app.mongo.Db != nil {
		go app.mongoLog(request)
	}
	if app.logStatus == true {
		go app.postgresqlLog(request)
	}

	return request, err
}
