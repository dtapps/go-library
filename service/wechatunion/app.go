package wechatunion

import (
	"encoding/json"
	"gopkg.in/dtapps/go-library.v2/utils/gohttp"
	"net/http"
)

type App struct {
	AppId       string // 小程序唯一凭证，即 AppID
	AppSecret   string // 小程序唯一凭证密钥，即 AppSecret
	AccessToken string // 接口调用凭证
}

func (app *App) request(url string, params map[string]interface{}, method string) (resp []byte, err error) {
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
