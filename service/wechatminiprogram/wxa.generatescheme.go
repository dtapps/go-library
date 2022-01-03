package wechatminiprogram

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WxaGenerateSchemeResponse struct {
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
	UrlLink interface{} `json:"url_link"`
}

type WxaGenerateSchemeResult struct {
	Result WxaGenerateSchemeResponse // 结果
	Body   []byte                    // 内容
	Err    error                     // 错误
}

func NewWxaGenerateSchemeResult(result WxaGenerateSchemeResponse, body []byte, err error) *WxaGenerateSchemeResult {
	return &WxaGenerateSchemeResult{Result: result, Body: body, Err: err}
}

// WxaGenerateScheme 获取小程序 scheme 码，适用于短信、邮件、外部网页、微信内等拉起小程序的业务场景。通过该接口，可以选择生成到期失效和永久有效的小程序码，有数量限制，目前仅针对国内非个人主体的小程序开放
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.generate.html
func (app *App) WxaGenerateScheme(notMustParams ...Params) *WxaGenerateSchemeResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/generatescheme?access_token=%s", app.AccessToken), params, http.MethodPost)
	// 定义
	var response WxaGenerateSchemeResponse
	err = json.Unmarshal(body, &response)
	return NewWxaGenerateSchemeResult(response, body, err)
}
