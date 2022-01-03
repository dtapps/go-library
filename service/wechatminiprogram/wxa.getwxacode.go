package wechatminiprogram

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WxaGetWxaCodeResponse struct {
	Errcode     int         `json:"errcode"`
	Errmsg      string      `json:"errmsg"`
	ContentType string      `json:"contentType"`
	Buffer      interface{} `json:"buffer"`
}

type WxaGetWxaCodeResult struct {
	Result WxaGetWxaCodeResponse // 结果
	Body   []byte                // 内容
	Err    error                 // 错误
}

func NewWxaGetWxaCodeResult(result WxaGetWxaCodeResponse, body []byte, err error) *WxaGetWxaCodeResult {
	return &WxaGetWxaCodeResult{Result: result, Body: body, Err: err}
}

// WxaGetWxaCode 获取小程序码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html
func (app *App) WxaGetWxaCode(notMustParams ...Params) *WxaGetWxaCodeResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacode?access_token=%s", app.AccessToken), params, http.MethodPost)
	// 定义
	var response WxaGetWxaCodeResponse
	err = json.Unmarshal(body, &response)
	return NewWxaGetWxaCodeResult(response, body, err)
}
