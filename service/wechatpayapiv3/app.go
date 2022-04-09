package wechatpayapiv3

import (
	"encoding/json"
	"errors"
	"github.com/dtapps/go-library/utils/goheader"
	"github.com/dtapps/go-library/utils/gohttp"
	"github.com/dtapps/go-library/utils/gomongo"
	"net/http"
)

// App 微信支付服务
type App struct {
	AppId           string      // 小程序或者公众号唯一凭证
	AppSecret       string      // 小程序或者公众号唯一凭证密钥
	MchId           string      // 微信支付的商户id
	AesKey          string      // 私钥
	ApiV3           string      // API v3密钥
	PrivateSerialNo string      // 私钥证书号
	MchPrivateKey   string      // 商户私有证书内容 apiclient_key.pem
	Mongo           gomongo.App // 日志数据库
}

func (app *App) request(url string, params map[string]interface{}, method string, commonParams bool) (resp gohttp.Response, err error) {
	// 公共参数
	if method == http.MethodPost {
		if commonParams == true {
			params["appid"] = app.AppId
			params["mchid"] = app.MchId
		}
	}
	authorization, err := app.authorization(method, params, url)
	if err != nil {
		return gohttp.Response{}, err
	}

	headers := goheader.NewHeaders()
	headers.Set("Authorization", "WECHATPAY2-SHA256-RSA2048 "+authorization)
	headers.Set("Accept", "application/json")
	headers.Set("Accept-Language", "zh-CN")

	switch method {
	case http.MethodGet:
		// 请求
		getJson, err := gohttp.GetJsonHeader(url, params, headers)
		// 日志
		go app.mongoLog(url, params, method, getJson)
		return getJson, err
	case http.MethodPost:
		// 请求参数
		paramsStr, err := json.Marshal(params)
		postJson, err := gohttp.PostJsonHeader(url, paramsStr, headers)
		// 日志
		go app.mongoLog(url, params, method, postJson)
		return postJson, err
	default:
		return resp, errors.New("请求类型不支持")
	}
}
