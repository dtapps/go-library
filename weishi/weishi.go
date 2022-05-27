package weishi

import (
	"errors"
	"go.dtapp.net/library/golog"
	"go.dtapp.net/library/gorequest"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type WeiShi struct {
	ua           string         // 用户代理
	pgsql        *gorm.DB       // pgsql数据库
	client       *gorequest.App // 请求客户端
	log          *golog.Api     // 日志服务
	logTableName string         // 日志表名
	logStatus    bool           // 日志状态
}

func NewWeiShi(pgsql *gorm.DB) *WeiShi {
	ws := &WeiShi{ua: "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"}
	ws.client = gorequest.NewHttp()
	if pgsql != nil {
		ws.pgsql = pgsql
		ws.logStatus = true
		ws.logTableName = "weishi"
		ws.log = golog.NewApi(&golog.ApiConfig{
			Db:        pgsql,
			TableName: ws.logTableName,
		})
	}
	return ws
}

func (ws *WeiShi) request(url string) (resp gorequest.Response, err error) {

	// 创建请求
	client := ws.client

	// 设置请求地址
	client.SetUri(url)

	// 设置用户代理
	client.SetUserAgent(ws.ua)

	// 发起请求
	request, err := client.Get()
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if ws.logStatus == true {
		go ws.postgresqlLog(request)
	}

	return request, err
}

func (ws *WeiShi) urlJudge(str string) string {
	if strings.Index(str, "weishi.qq.com") != -1 {
		return str
	}
	return ""
}

func (ws *WeiShi) request302(url string) (string, error) {
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
