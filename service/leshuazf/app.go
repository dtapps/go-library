package leshuazf

import (
	"dtapps/dta/library/utils/gohttp"
	"dtapps/dta/library/utils/gomongo"
	"dtapps/dta/library/utils/gopostgresql"
	"errors"
	"gitee.com/dtapps/go-library/utils/gorandom"
	"gitee.com/dtapps/go-library/utils/gotime"
	"log"
	"net/http"
)

// App 乐刷
type App struct {
	AgentId     string // 服务商编号，由乐刷分配的接入方唯一标识，明文传输。
	Environment string //  环境
	KeyAgent    string
	Pgsql       gopostgresql.App // 日志数据库
	Mongo       gomongo.App      // 日志数据库
}

func (app *App) request(url string, params map[string]interface{}, method string) (resp []byte, err error) {
	// 环境
	if app.Environment == "test" {
		url = "http://t-saas-mch.lepass.cn/" + url
	} else {
		url = "https://saas-mch.leshuazf.com/" + url
	}
	// 参数
	params["agentId"] = app.AgentId                                                            // 服务商编号，由乐刷分配的接入方唯一标识，明文传输。
	params["version"] = "2.0"                                                                  // 目前固定值2.0
	params["reqSerialNo"] = gotime.Current().SetFormat("20060102150405") + gorandom.Numeric(5) // 请求流水号(yyyyMMddHHmmssSSSXXXXX，其中 XXXXX为5位顺序号,禁止使用UUID等无意义数据)
	params["sign"] = app.getSign(params)

	log.Println(app)
	log.Println(url)
	log.Println(params)

	switch method {
	case http.MethodGet:
		// 请求
		get, err := gohttp.Get(url, params)
		return get.Body, err
	case http.MethodPost:
		postJson, err := gohttp.PostForm(url, params)
		return postJson.Body, err
	default:
		return nil, errors.New("请求类型不支持")
	}
}
