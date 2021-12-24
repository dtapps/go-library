package dingdanxia

import (
	"net/http"

	"github.com/dtapps/go-library/utils/gohttp"
)

type App struct {
	ApiKey string
}

func (app *App) request(url string, params map[string]interface{}, method string) ([]byte, error) {
	// 公共参数
	params["apikey"] = app.ApiKey
	// 请求
	if method == http.MethodGet {
		get, err := gohttp.Get(url, params)
		return get.Body, err
	} else {
		postJson, err := gohttp.PostForm(url, params)
		return postJson.Body, err
	}
}
