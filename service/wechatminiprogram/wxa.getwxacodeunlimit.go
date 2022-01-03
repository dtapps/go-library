package wechatminiprogram

import (
	"encoding/json"
	"fmt"
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
	"net/http"
)

type WxaGetWxaCodeUnLimitResponse struct {
	Errcode     int         `json:"errcode"`
	Errmsg      string      `json:"errmsg"`
	ContentType string      `json:"contentType"`
	Buffer      interface{} `json:"buffer"`
}

type WxaGetWxaCodeUnLimitResult struct {
	Result WxaGetWxaCodeUnLimitResponse // 结果
	Body   []byte                       // 内容
	Http   gohttp.Response              // 请求
	Err    error                        // 错误
}

func NewWxaGetWxaCodeUnLimitResult(result WxaGetWxaCodeUnLimitResponse, body []byte, http gohttp.Response, err error) *WxaGetWxaCodeUnLimitResult {
	return &WxaGetWxaCodeUnLimitResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaGetWxaCodeUnLimit 获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (app *App) WxaGetWxaCodeUnLimit(notMustParams ...Params) *WxaGetWxaCodeUnLimitResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s", app.AccessToken), params, http.MethodPost)
	// 定义
	var response WxaGetWxaCodeUnLimitResponse
	err = json.Unmarshal(request.Body, &response)
	return NewWxaGetWxaCodeUnLimitResult(response, request.Body, request, err)
}
