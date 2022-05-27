package wechatminiprogram

import (
	"go.dtapp.net/golog"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/goredis"
	"gorm.io/gorm"
)

const (
	WECHAT_API_URL = "https://api.weixin.qq.com"
	WECHAT_MP_URL  = "https://mp.weixin.qq.com"
	CGIUrl         = WECHAT_API_URL + "/cgi-bin"
	UnionUrl       = WECHAT_API_URL + "/union"
)

// App 微信小程序服务
type App struct {
	appId        string         // 小程序唯一凭证，即 appId
	appSecret    string         // 小程序唯一凭证密钥，即 appSecret
	accessToken  string         // 接口调用凭证
	jsapiTicket  string         // 签名凭证
	redis        goredis.App    // 缓存数据库
	db           *gorm.DB       // 令牌数据库
	pgsql        *gorm.DB       // pgsql数据库
	client       *gorequest.App // 请求客户端
	log          *golog.Api     // 日志服务
	logTableName string         // 日志表名
	logStatus    bool           // 日志状态
}

// NewApp 实例化
func NewApp(appId, appSecret string, redis goredis.App, db *gorm.DB, pgsql *gorm.DB) *App {
	app := &App{appId: appId, appSecret: appSecret, redis: redis, db: db}
	app.client = gorequest.NewHttp()
	if pgsql != nil {
		app.pgsql = pgsql
		app.logStatus = true
		app.logTableName = "wechatminiprogram"
		app.log = golog.NewApi(&golog.ApiConfig{
			Db:        pgsql,
			TableName: app.logTableName,
		})
	}
	return app
}

// Config 配置
func (app *App) Config(appId, appSecret string) *App {
	app.appId = appId
	app.appSecret = appSecret
	return app
}

// 请求接口
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

func (app *App) GetAppId() string {
	return app.appId
}
