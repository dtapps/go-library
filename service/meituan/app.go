package meituan

import (
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
	"net/http"
)

// App 美团联盟
type App struct {
	Secret string      // 秘钥
	AppKey string      // 渠道标记
	ZapLog *zap.Logger // 日志服务
}

func (app *App) request(url string, params map[string]interface{}, method string) (resp []byte, err error) {
	switch method {
	case http.MethodGet:
		// 请求
		get, err := gohttp.Get(url, params)
		// 日志
		if app.ZapLog != nil {
			app.ZapLog.Sugar().Info(get)
		}
		return get.Body, err
	case http.MethodPost:
		// 请求
		paramsStr, err := json.Marshal(params)
		postJson, err := gohttp.PostJson(url, paramsStr)
		// 日志
		if app.ZapLog != nil {
			app.ZapLog.Sugar().Info(postJson)
		}
		return postJson.Body, err
	default:
		return nil, errors.New("请求类型不支持")
	}
}
