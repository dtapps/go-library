package wechatpayapiv3

import (
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gomongo"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
	"net/http"
)

// App 微信支付直连商户
type App struct {
	appId          string          // 小程序或者公众号唯一凭证
	appSecret      string          // 小程序或者公众号唯一凭证密钥
	mchId          string          // 微信支付的商户id
	aesKey         string          // 私钥
	apiV3          string          // API v3密钥
	mchSslSerialNo string          // pem 证书号
	mchSslKey      string          // pem key 内容
	mongo          *gomongo.Client // 日志数据库
	pgsql          *gorm.DB        // pgsql数据库
	client         *gorequest.App  // 请求客户端
	log            *golog.Api      // 日志服务
	logTableName   string          // 日志表名
	logStatus      bool            // 日志状态
}

// NewApp 实例化
func NewApp(appId string, appSecret string, mchId string, aesKey string, apiV3 string, mchSslSerialNo string, mchSslKey string, pgsql *gorm.DB) *App {
	app := &App{appId: appId, appSecret: appSecret, mchId: mchId, aesKey: aesKey, apiV3: apiV3, mchSslSerialNo: mchSslSerialNo, mchSslKey: mchSslKey}
	app.client = gorequest.NewHttp()
	if pgsql != nil {
		app.pgsql = pgsql
		app.logStatus = true
		app.logTableName = "wechatpayapiv3"
		app.log = golog.NewApi(&golog.ApiConfig{
			Db:        pgsql,
			TableName: app.logTableName,
		})
	}
	return app
}

func (app *App) request(url string, params map[string]interface{}, method string, commonParams bool) (resp gorequest.Response, err error) {

	// 公共参数
	if method == http.MethodPost {
		if commonParams == true {
			params["appid"] = app.appId
			params["mchid"] = app.mchId
		}
	}

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
	client.SetHeader("Authorization", "WECHATPAY2-SHA256-RSA2048 "+authorization)
	client.SetHeader("Accept", "application/json")
	client.SetHeader("Accept-Language", "zh-CN")
	if url == "https://api.mch.weixin.qq.com/v3/merchant-service/complaints-v2" {
		client.SetHeader("Wechatpay-Serial", app.mchSslSerialNo)
	}

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
