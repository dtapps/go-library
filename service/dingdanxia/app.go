package dingdanxia

import (
	"errors"
	"go.dtapp.net/library/utils/gohttp"
	"go.dtapp.net/library/utils/gomongo"
	"net/http"
)

type App struct {
	ApiKey string
	Mongo  gomongo.App // 日志数据库
}

func (app *App) request(url string, params map[string]interface{}, method string) (resp gohttp.Response, err error) {
	// 公共参数
	params["apikey"] = app.ApiKey
	switch method {
	case http.MethodGet:
		// 请求
		get, err := gohttp.Get(url, params)
		// 日志
		go app.mongoLog(url, params, method, get)
		return get, err
	case http.MethodPost:
		// 请求
		postJson, err := gohttp.PostForm(url, params)
		// 日志
		go app.mongoLog(url, params, method, postJson)
		return postJson, err
	default:
		return resp, errors.New("请求类型不支持")
	}
}
