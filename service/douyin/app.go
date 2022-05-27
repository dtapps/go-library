package douyin

import (
	"errors"
	"go.dtapp.net/golog"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type App struct {
	ua           string         // 用户代理
	pgsql        *gorm.DB       // pgsql数据库
	client       *gorequest.App // 请求客户端
	log          *golog.Api     // 日志服务
	logTableName string         // 日志表名
	logStatus    bool           // 日志状态
}

func NewApp(pgsql *gorm.DB) *App {
	app := &App{ua: "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"}
	app.client = gorequest.NewHttp()
	if pgsql != nil {
		app.pgsql = pgsql
		app.logStatus = true
		app.logTableName = "douyin"
		app.log = golog.NewApi(&golog.ApiConfig{
			Db:        pgsql,
			TableName: app.logTableName,
		})
	}
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

	// 设置用户代理
	client.SetUserAgent(app.ua)

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

func (app *App) request302(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	client := new(http.Client)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("redirect")
	}

	response, err := client.Do(req)
	if err != nil {
		if response.StatusCode == http.StatusFound {
			location, err := response.Location()
			return location.String(), err
		} else {
			return "", err
		}
	}

	return "", nil
}

func (app *App) urlJudge(str string) string {
	if strings.Index(str, "douyin.com") != -1 || strings.Index(str, "iesdouyin.com") != -1 {
		return str
	}
	return ""
}
