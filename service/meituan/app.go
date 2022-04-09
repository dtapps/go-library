package meituan

import (
	"dtapps/dta/library/utils/gohttp"
	"dtapps/dta/library/utils/gomongo"
	"encoding/json"
	"errors"
	"net/http"
)

// App 美团联盟
type App struct {
	Secret string      // 秘钥
	AppKey string      // 渠道标记
	Mongo  gomongo.App // 日志数据库
}

func (app *App) request(url string, params map[string]interface{}, method string) (resp []byte, err error) {
	switch method {
	case http.MethodGet:
		// 请求
		get, err := gohttp.Get(url, params)
		// 日志
		go app.mongoLog(url, params, method, get)
		return get.Body, err
	case http.MethodPost:
		// 请求
		paramsStr, err := json.Marshal(params)
		postJson, err := gohttp.PostJson(url, paramsStr)
		// 日志
		go app.mongoLog(url, params, method, postJson)
		return postJson.Body, err
	default:
		return nil, errors.New("请求类型不支持")
	}
}
