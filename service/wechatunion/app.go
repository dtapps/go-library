package wechatunion

import (
	"dtapps/dta/library/utils/gohttp"
	"dtapps/dta/library/utils/gomongo"
	"dtapps/dta/library/utils/goredis"
	"encoding/json"
	"errors"
	"net/http"
)

// App 微信小程序联盟
type App struct {
	AppId       string      // 小程序唯一凭证，即 AppID
	AppSecret   string      // 小程序唯一凭证密钥，即 AppSecret
	AccessToken string      // 接口调用凭证
	Pid         string      // 推广位PID
	Redis       goredis.App // 缓存数据库服务
	Mongo       gomongo.App // 日志数据库
}

const (
	UnionUrl = "https://api.weixin.qq.com/union"
)

// 请求
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
