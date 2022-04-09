package wechatminiprogram

import (
	"encoding/json"
	"errors"
	"github.com/dtapps/go-library/utils/gohttp"
	"github.com/dtapps/go-library/utils/gomongo"
	"github.com/dtapps/go-library/utils/goredis"
	"net/http"
)

// App 微信小程序服务
type App struct {
	AppId       string      // 小程序唯一凭证，即 AppID
	AppSecret   string      // 小程序唯一凭证密钥，即 AppSecret
	AccessToken string      // 接口调用凭证
	JsapiTicket string      // 签名凭证
	Redis       goredis.App // 缓存数据库服务
	Mongo       gomongo.App // 日志数据库
}

const (
	WECHAT_API_URL = "https://api.weixin.qq.com"
	WECHAT_MP_URL  = "https://mp.weixin.qq.com"
	CGIUrl         = WECHAT_API_URL + "/cgi-bin"
	UnionUrl       = WECHAT_API_URL + "/union"
)

// 请求
func (app *App) request(url string, params map[string]interface{}, method string) (resp gohttp.Response, err error) {
	switch method {
	case http.MethodGet:
		get, err := gohttp.Get(url, params)
		// 日志
		go app.mongoLog(url, params, method, get)
		return get, err
	case http.MethodPost:
		// 请求参数
		paramsStr, err := json.Marshal(params)
		postJson, err := gohttp.PostJson(url, paramsStr)
		// 日志
		go app.mongoLog(url, params, method, postJson)
		return postJson, err
	default:
		return resp, errors.New("请求类型不支持")
	}
}
