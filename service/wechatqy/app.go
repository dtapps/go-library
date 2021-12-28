package wechatqy

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
)

type App struct {
	Key    string
	ZapLog *zap.Logger // 日志服务
}

func (app *App) request(url string, params map[string]interface{}) (body []byte, err error) {
	// 请求参数
	paramsStr, err := json.Marshal(params)
	// 请求
	postJson, err := gohttp.PostJson(url, paramsStr)
	// 日志
	if app.ZapLog != nil {
		app.ZapLog.Sugar().Info(fmt.Sprintf("%s", postJson.Body))
	}
	return postJson.Body, err
}
