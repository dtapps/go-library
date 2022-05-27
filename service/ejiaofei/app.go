package ejiaofei

import (
	"errors"
	"fmt"
	"go.dtapp.net/library/utils/gohttp"
	"go.dtapp.net/library/utils/gomd5"
	"go.dtapp.net/library/utils/gomongo"
	"net/http"
)

type App struct {
	UserID  string
	Pwd     string
	Key     string
	signStr string
	Mongo   gomongo.App // 日志数据库
}

func (app *App) request(url string, params map[string]interface{}, method string) ([]byte, error) {
	// 公共参数
	params["userid"] = app.UserID
	params["pwd"] = app.Pwd
	// 签名
	params["userkey"] = gomd5.ToUpper(fmt.Sprintf("%s%s", app.signStr, app.Key))
	switch method {
	case http.MethodGet:
		// 请求
		get, err := gohttp.Get(url, params)
		// 日志
		go app.mongoLog(url, params, method, get)
		return get.Body, err
	case http.MethodPost:
		// 请求
		postJson, err := gohttp.PostForm(url, params)
		// 日志
		go app.mongoLog(url, params, method, postJson)
		return postJson.Body, err
	default:
		return nil, errors.New("请求类型不支持")
	}
}
