package wechatopen

import (
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/goredis"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

// App 微信公众号服务
type App struct {
	componentAccessToken   string // 第三方平台 access_token
	componentVerifyTicket  string // 微信后台推送的 ticket
	preAuthCode            string // 预授权码
	authorizerAccessToken  string // 接口调用令牌
	authorizerRefreshToken string // 刷新令牌
	authorizerAppid        string // 授权方 appid
	componentAppId         string // 第三方平台 appid
	componentAppSecret     string // 第三方平台 app_secret
	messageToken           string
	messageKey             string
	redis                  goredis.App    // 缓存数据库
	pgsql                  *gorm.DB       // pgsql数据库
	client                 *gorequest.App // 请求客户端
	log                    *golog.Api     // 日志服务
	logTableName           string         // 日志表名
	logStatus              bool           // 日志状态
}

// NewApp 实例化
func NewApp(componentAppId string, componentAppSecret string, messageToken string, messageKey string, redis goredis.App, pgsql *gorm.DB) *App {
	app := &App{componentAppId: componentAppId, componentAppSecret: componentAppSecret, messageToken: messageToken, messageKey: messageKey, redis: redis}
	app.client = gorequest.NewHttp()
	if pgsql != nil {
		app.pgsql = pgsql
		app.logStatus = true
		app.logTableName = "wechatopen"
		app.log = golog.NewApi(&golog.ApiConfig{
			Db:        pgsql,
			TableName: app.logTableName,
		})
	}
	return app
}

// Config 配置
func (app *App) Config(componentAppId, componentAppSecret string) *App {
	app.componentAppId = componentAppId
	app.componentAppSecret = componentAppSecret
	return app
}

// ConfigAuthorizer 配置第三方
func (app *App) ConfigAuthorizer(authorizerAppid string) *App {
	app.authorizerAppid = authorizerAppid
	return app
}

func (app *App) request(url string, params map[string]interface{}, method string) (resp gorequest.Response, err error) {

	// 创建请求
	client := app.client

	// 设置请求地址
	client.SetUri(url)

	// 设置方式
	client.SetMethod(method)

	// 设置格式
	client.SetContentTypeJson()

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
