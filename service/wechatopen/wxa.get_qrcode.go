package wechatopen

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WxaGetQrcodeResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type WxaGetQrcodeResult struct {
	Result WxaGetQrcodeResponse // 结果
	Body   []byte               // 内容
	Err    error                // 错误
}

func NewWxaGetQrcodeResult(result WxaGetQrcodeResponse, body []byte, err error) *WxaGetQrcodeResult {
	return &WxaGetQrcodeResult{Result: result, Body: body, Err: err}
}

// WxaGetQrcode 获取体验版二维码
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_qrcode.html
func (app *App) WxaGetQrcode(path string) *WxaGetQrcodeResult {
	app.authorizerAccessToken = app.GetAuthorizerAccessToken()
	// 参数
	params := NewParams()
	if path != "" {
		params["path"] = path // 指定二维码扫码后直接进入指定页面并可同时带上参数）
	}
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/get_qrcode?access_token=%s", app.authorizerAccessToken), params, http.MethodGet)
	// 定义
	var response WxaGetQrcodeResponse
	err = json.Unmarshal(body, &response)
	return NewWxaGetQrcodeResult(response, body, err)
}
