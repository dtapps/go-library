package leshuazf

import (
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gomongo"
	"go.dtapp.net/library/utils/gorandom"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
	"gorm.io/gorm"
)

// App 乐刷
type App struct {
	agentId      string // 服务商编号，由乐刷分配的接入方唯一标识，明文传输。
	Environment  string //  环境
	KeyAgent     string
	mongo        *gomongo.Client // 日志数据库
	pgsql        *gorm.DB        // pgsql数据库
	client       *gorequest.App  // 请求客户端
	log          *golog.Api      // 日志服务
	logTableName string          // 日志表名
	logStatus    bool            // 日志状态
}

func NewApp(agentId string, environment string, keyAgent string, pgsql *gorm.DB) *App {
	app := &App{agentId: agentId, Environment: environment, KeyAgent: keyAgent}
	app.client = gorequest.NewHttp()
	if pgsql != nil {
		app.pgsql = pgsql
		app.logStatus = true
		app.logTableName = "leshuazf"
		app.log = golog.NewApi(&golog.ApiConfig{
			Db:        pgsql,
			TableName: app.logTableName,
		})
	}
	return app
}

func (app *App) request(url string, params map[string]interface{}, method string) (resp gorequest.Response, err error) {

	// 环境
	if app.Environment == "test" {
		url = "http://t-saas-mch.lepass.cn/" + url
	} else {
		url = "https://saas-mch.leshuazf.com/" + url
	}

	// 参数
	params["agentId"] = app.agentId                                                            // 服务商编号，由乐刷分配的接入方唯一标识，明文传输。
	params["version"] = "2.0"                                                                  // 目前固定值2.0
	params["reqSerialNo"] = gotime.Current().SetFormat("20060102150405") + gorandom.Numeric(5) // 请求流水号(yyyyMMddHHmmssSSSXXXXX，其中 XXXXX为5位顺序号,禁止使用UUID等无意义数据)
	params["sign"] = app.getSign(params)

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
