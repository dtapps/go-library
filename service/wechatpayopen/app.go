package wechatpayopen

import (
	"go.dtapp.net/golog"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

// App 微信支付服务器
type App struct {
	spAppid        string         // 服务商应用ID
	spMchId        string         // 服务商户号
	subAppid       string         // 子商户应用ID
	subMchId       string         // 子商户号
	apiV2          string         // APIv2密钥
	apiV3          string         // APIv3密钥
	serialNo       string         // 序列号
	mchSslSerialNo string         // pem 证书号
	mchSslCer      string         // pem 内容
	mchSslKey      string         // pem key 内容
	pgsql          *gorm.DB       // pgsql数据库
	client         *gorequest.App // 请求客户端
	log            *golog.Api     // 日志服务
	logTableName   string         // 日志表名
	logStatus      bool           // 日志状态
}

// NewApp 实例化
func NewApp(spAppid, spMchId, subAppid, subMchId, apiV2, apiV3, serialNo, mchSslSerialNo, mchSslCer, mchSslKey string, pgsql *gorm.DB) *App {
	app := &App{spAppid: spAppid, spMchId: spMchId, subAppid: subAppid, subMchId: subMchId, apiV2: apiV2, apiV3: apiV3, serialNo: serialNo, mchSslSerialNo: mchSslSerialNo, mchSslCer: mchSslCer, mchSslKey: mchSslKey}
	app.client = gorequest.NewHttp()
	if pgsql != nil {
		app.pgsql = pgsql
		app.logStatus = true
		app.logTableName = "wechatpayopen"
		app.log = golog.NewApi(&golog.ApiConfig{
			Db:        pgsql,
			TableName: app.logTableName,
		})
	}
	return app
}

// NewAppConfig 实例化
func (app *App) NewAppConfig(subAppid, subMchId string) *App {
	app.subAppid = subAppid
	app.subMchId = subMchId
	return app
}

func (app *App) request(url string, params map[string]interface{}, method string) (resp gorequest.Response, err error) {

	// 认证
	authorization, err := app.authorization(method, params, url)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 创建请求
	client := app.client

	// 设置请求地址
	client.SetUri(url)

	// 设置方式
	client.SetMethod(method)

	// 设置JSON格式
	client.SetContentTypeJson()

	// 设置参数
	client.SetParams(params)

	// 设置头部
	client.SetHeader("Authorization", authorization)
	client.SetHeader("Accept", "application/json")
	client.SetHeader("Accept-Language", "zh-CN")

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
