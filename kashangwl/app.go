package kashangwl

import (
	"go.dtapp.net/library/golog"
	"go.dtapp.net/library/gorequest"
	"gorm.io/gorm"
	"time"
)

// App 卡商网服务
type App struct {
	customerId   int            // 商家编号
	customerKey  string         // 商家密钥
	pgsql        *gorm.DB       // pgsql数据库
	client       *gorequest.App // 请求客户端
	log          *golog.Api     // 日志服务
	logTableName string         // 日志表名
	logStatus    bool           // 日志状态
}

func NewApp(customerId int, customerKey string, pgsql *gorm.DB) *App {
	app := &App{customerId: customerId, customerKey: customerKey}
	app.client = gorequest.NewHttp()
	if pgsql != nil {
		app.pgsql = pgsql
		app.logStatus = true
		app.logTableName = "kashangwl"
		app.log = golog.NewApi(&golog.ApiConfig{
			Db:        pgsql,
			TableName: app.logTableName,
		})
	}
	return app
}

func (app *App) request(url string, params map[string]interface{}) (resp gorequest.Response, err error) {

	// 公共参数
	params["timestamp"] = time.Now().UnixNano() / 1e6
	params["customer_id"] = app.customerId

	// 签名参数
	params["sign"] = app.getSign(app.customerKey, params)

	// 创建请求
	client := app.client

	// 设置请求地址
	client.SetUri(url)

	// 设置格式
	client.SetContentTypeJson()

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Post()
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if app.logStatus == true {
		go app.postgresqlLog(request)
	}

	return request, err
}
