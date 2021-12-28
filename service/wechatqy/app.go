package wechatqy

import (
	"encoding/json"
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
	"gopkg.in/dtapps/go-library.v3/utils/golog"
)

type App struct {
	Key    string
	ZapLog golog.App // 日志服务
}

func (app *App) request(url string, params map[string]interface{}) (body []byte, err error) {
	// 请求参数
	paramsStr, err := json.Marshal(params)
	// 请求
	postJson, err := gohttp.PostJson(url, paramsStr)
	// 日志
	if app.ZapLog.Logger != nil {
		app.ZapLog.LogName = "wechatqy.log"
		app.ZapLog.Logger.Sugar().Info(postJson)
	}
	return postJson.Body, err
}
