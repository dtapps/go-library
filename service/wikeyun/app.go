package wikeyun

import (
	"fmt"
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
	"gopkg.in/dtapps/go-library.v3/utils/golog"
)

type App struct {
	StoreID   int
	AppKey    int
	AppSecret string
	ClientIP  string
	ZapLog    golog.App // 日志服务
}

func (app *App) request(url string, params map[string]interface{}) (resp []byte, err error) {
	// 签名
	sign := app.sign(params)
	// 请求
	requestUrl := fmt.Sprintf("%s?app_key=%d&timestamp=%s&client=%s&format=%s&v=%s&sign=%s", url, app.AppKey, sign.Timestamp, sign.Client, sign.Format, sign.V, sign.Sign)
	postForm, err := gohttp.PostForm(requestUrl, params)
	// 日志
	if app.ZapLog.Logger != nil {
		app.ZapLog.LogName = "wikeyun.log"
		app.ZapLog.Logger.Sugar().Info(postForm)
	}
	return postForm.Body, err
}
