package eastiot

import (
	"errors"
	"go.uber.org/zap"
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
	"net/http"
	"time"
)

type App struct {
	AppID  string
	ApiKey string
	ZapLog *zap.Logger // 日志服务
}

func (app *App) request(url string, params map[string]interface{}, method string) ([]byte, error) {
	// 公共参数
	params["appId"] = app.AppID
	params["timeStamp"] = time.Now().Unix()
	// 签名
	params["sign"] = app.getSign(app.ApiKey, params)
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
		postJson, err := gohttp.PostForm(url, params)
		// 日志
		if app.ZapLog != nil {
			app.ZapLog.Sugar().Info(postJson)
		}
		return postJson.Body, err
	default:
		return nil, errors.New("请求类型不支持")
	}
}
