package tianyancha

import (
	"dtapps/dta/library/utils/gohttp"
	"encoding/json"
	"net/http"
)

type App struct{}

func (app *App) request(url string, params map[string]interface{}, method string) (resp []byte, err error) {
	// 请求
	if method == http.MethodGet {
		get, err := gohttp.Get(url, params)
		return get.Body, err
	} else {
		paramsStr, err := json.Marshal(params)
		postJson, err := gohttp.PostJson(url, paramsStr)
		return postJson.Body, err
	}
}
