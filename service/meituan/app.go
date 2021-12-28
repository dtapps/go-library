package meituan

import (
	"encoding/json"
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
	"net/http"
)

// App 美团联盟
type App struct {
	Secret string // 秘钥
	AppKey string // 渠道标记
}

func (app *App) request(url string, params map[string]interface{}, method string) (resp []byte, err error) {

	// GET方式
	if method == http.MethodGet {
		get, err := gohttp.Get(url, params)
		return get.Body, err
	} else {
		// 请求参数
		paramsStr, err := json.Marshal(params)
		postJson, err := gohttp.PostJson(url, paramsStr)
		return postJson.Body, err
	}

}
