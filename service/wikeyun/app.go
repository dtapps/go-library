package wikeyun

import (
	"fmt"
	"go.dtapp.net/library/utils/goip"
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

type App struct {
	storeId      int            // 店铺ID
	appKey       int            // key
	appSecret    string         // secret
	clientIp     string         // Ip
	pgsql        *gorm.DB       // pgsql数据库
	client       *gorequest.App // 请求客户端
	log          *golog.Api     // 日志服务
	logTableName string         // 日志表名
	logStatus    bool           // 日志状态
}

// NewApp 创建实例
func NewApp(storeId, appKey int, appSecret string, pgsql *gorm.DB) *App {
	app := &App{storeId: storeId, appKey: appKey, appSecret: appSecret}
	app.client = gorequest.NewHttp()
	if pgsql != nil {
		app.pgsql = pgsql
		app.logStatus = true
		app.logTableName = "wikeyun"
		app.log = golog.NewApi(&golog.ApiConfig{
			Db:        pgsql,
			TableName: app.logTableName,
		})
	}
	xip := goip.GetOutsideIp()
	if xip != "" && xip != "0.0.0.0" {
		app.clientIp = xip
	}
	return app
}

// 请求接口
func (app *App) request(url string, params map[string]interface{}) (resp gorequest.Response, err error) {

	// 签名
	sign := app.sign(params)

	// 创建请求
	client := app.client

	// 设置请求地址
	client.SetUri(fmt.Sprintf("%s?app_key=%d&timestamp=%s&client=%s&format=%s&v=%s&sign=%s", url, app.appKey, sign.Timestamp, sign.Client, sign.Format, sign.V, sign.Sign))

	// 设置FORM格式
	client.SetContentTypeForm()

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
