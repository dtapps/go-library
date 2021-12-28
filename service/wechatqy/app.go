package wechatqy

import (
	"encoding/json"
	"gopkg.in/dtapps/go-library.v2/utils/gohttp"
)

type App struct {
	Key string
}

func (app *App) request(url string, params map[string]interface{}) (body []byte, err error) {
	// 请求参数
	paramsStr, err := json.Marshal(params)
	// 请求
	postJson, err := gohttp.PostJson(url, paramsStr)
	return postJson.Body, err
}
