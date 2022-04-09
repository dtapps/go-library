package dingtalk

import (
	"encoding/json"
	"errors"
	"github.com/dtapps/go-library/utils/gohttp"
	"github.com/dtapps/go-library/utils/gomongo"
	"net/http"
)

type App struct {
	Secret      string
	AccessToken string
	Mongo       gomongo.App // 日志数据库
}

func (app *App) request(url string, params map[string]interface{}, method string) ([]byte, error) {
	switch method {
	case http.MethodGet:
		// 请求
		get, err := gohttp.Get(url, params)
		// 日志
		go app.mongoLog(url, params, method, get)
		return get.Body, err
	case http.MethodPost:
		// 请求参数
		paramsMarshal, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		// 请求
		postJson, err := gohttp.PostJson(url, paramsMarshal)
		// 日志
		go app.mongoLog(url, params, method, postJson)
		return postJson.Body, err
	default:
		return nil, errors.New("请求类型不支持")
	}
}
