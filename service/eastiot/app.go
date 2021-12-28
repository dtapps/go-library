package eastiot

import (
	"gopkg.in/dtapps/go-library.v2/utils/gohttp"
	"net/http"
	"time"
)

type App struct {
	AppID  string
	ApiKey string
}

func (app *App) request(url string, params map[string]interface{}, method string) ([]byte, error) {
	// 公共参数
	params["appId"] = app.AppID
	params["timeStamp"] = time.Now().Unix()
	// 签名
	params["sign"] = app.getSign(app.ApiKey, params)
	// 请求
	if method == http.MethodGet {
		get, err := gohttp.Get(url, params)
		return get.Body, err
	} else {
		postJson, err := gohttp.PostForm(url, params)
		return postJson.Body, err
	}
}
