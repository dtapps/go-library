package wechatunion

import (
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/goredis"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

const (
	UnionUrl = "https://api.weixin.qq.com/union"
)

// App 微信小程序联盟
type App struct {
	appId        string         // 小程序唯一凭证，即 appId
	appSecret    string         // 小程序唯一凭证密钥，即 appSecret
	accessToken  string         // 接口调用凭证
	pid          string         // 推广位PID
	Redis        goredis.App    // 缓存数据库服务
	pgsql        *gorm.DB       // pgsql数据库
	client       *gorequest.App // 请求客户端
	log          *golog.Api     // 日志服务
	logTableName string         // 日志表名
	logStatus    bool           // 日志状态
}

func NewApp(appId string, appSecret string, pid string, redis goredis.App, pgsql *gorm.DB) *App {
	app := &App{appId: appId, appSecret: appSecret, pid: pid, Redis: redis}
	app.client = gorequest.NewHttp()
	if pgsql != nil {
		app.pgsql = pgsql
		app.logStatus = true
		app.logTableName = "wechatunion"
		app.log = golog.NewApi(&golog.ApiConfig{
			Db:        pgsql,
			TableName: app.logTableName,
		})
	}
	return app
}

// 请求
func (app *App) request(url string, params map[string]interface{}, method string) (resp gorequest.Response, err error) {

	// 创建请求
	client := app.client

	// 设置请求地址
	client.SetUri(url)

	// 设置请求方式
	client.SetMethod(method)

	// 设置FORM格式
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
