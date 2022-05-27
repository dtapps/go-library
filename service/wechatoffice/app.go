package wechatoffice

import (
	"encoding/json"
	"errors"
	"go.dtapp.net/library/utils/gohttp"
	"go.dtapp.net/library/utils/gomongo"
	"go.dtapp.net/library/utils/goredis"
	"net/http"
)

// App 微信公众号服务
type App struct {
	AppId       string      // 小程序唯一凭证，即 AppID
	AppSecret   string      // 小程序唯一凭证密钥，即 AppSecret
	AccessToken string      // 接口调用凭证
	JsapiTicket string      // 签名凭证
	Redis       goredis.App // 缓存数据库服务
	Mongo       gomongo.App // 日志数据库
}

func (app *App) request(url string, params map[string]interface{}, method string) (resp []byte, err error) {
	switch method {
	case http.MethodGet:
		get, err := gohttp.Get(url, params)
		// 日志
		go app.mongoLog(url, params, method, get)
		return get.Body, err
	case http.MethodPost:
		// 请求参数
		paramsStr, err := json.Marshal(params)
		postJson, err := gohttp.PostJson(url, paramsStr)
		// 日志
		go app.mongoLog(url, params, method, postJson)
		return postJson.Body, err
	default:
		return nil, errors.New("请求类型不支持")
	}
}
