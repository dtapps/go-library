package sendcloud

import (
	"dtapps/dta/library/utils/gohttp"
	"dtapps/dta/library/utils/gomongo"
	"errors"
	"net/http"
)

// App 公共请求参数
type App struct {
	ApiUser string      // API_USER
	ApiKey  string      // API_KEY
	Mongo   gomongo.App // 日志数据库
}

func (app *App) request(url string, params map[string]interface{}, method string) ([]byte, error) {
	// 公共参数
	//params["userid"] = app.UserID
	//params["pwd"] = app.Pwd
	//// 签名
	//params["userkey"] = gomd5.ToUpper(fmt.Sprintf("%s%s", app.signStr, app.Key))
	switch method {
	case http.MethodGet:
		// 请求
		get, err := gohttp.Get(url, params)
		return get.Body, err
	case http.MethodPost:
		// 请求
		postJson, err := gohttp.PostForm(url, params)
		return postJson.Body, err
	default:
		return nil, errors.New("请求类型不支持")
	}
}
