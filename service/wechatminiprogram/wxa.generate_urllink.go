package wechatminiprogram

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WxaGenerateUrlLinkResponse struct {
	Errcode  int         `json:"errcode"`
	Errmsg   string      `json:"errmsg"`
	Openlink interface{} `json:"openlink"`
}

type WxaGenerateUrlLinkResult struct {
	Result WxaGenerateUrlLinkResponse // 结果
	Body   []byte                     // 内容
	Err    error                      // 错误
}

func NewWxaGenerateUrlLinkResult(result WxaGenerateUrlLinkResponse, body []byte, err error) *WxaGenerateUrlLinkResult {
	return &WxaGenerateUrlLinkResult{Result: result, Body: body, Err: err}
}

// WxaGenerateUrlLink 获取小程序 URL Link，适用于短信、邮件、网页、微信内等拉起小程序的业务场景。通过该接口，可以选择生成到期失效和永久有效的小程序链接，有数量限制，目前仅针对国内非个人主体的小程序开放
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html
func (app *App) WxaGenerateUrlLink(notMustParams ...Params) *WxaGenerateUrlLinkResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/generate_urllink?access_token=%s", app.AccessToken), params, http.MethodPost)
	// 定义
	var response WxaGenerateUrlLinkResponse
	err = json.Unmarshal(body, &response)
	return NewWxaGenerateUrlLinkResult(response, body, err)
}
