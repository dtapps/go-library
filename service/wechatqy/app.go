package wechatqy

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gohttp"
	"go.dtapp.net/library/utils/gomongo"
	"net/http"
)

type App struct {
	Key   string
	Mongo gomongo.App // 日志数据库
}

func (app *App) request(url string, params map[string]interface{}) (body []byte, err error) {
	// 请求参数
	paramsStr, err := json.Marshal(params)
	// 请求
	postJson, err := gohttp.PostJson(url, paramsStr)
	// 日志
	go app.mongoLog(url, params, http.MethodPost, postJson)
	return postJson.Body, err
}
