package wechatopen

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
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
	Http   gorequest.Response           // 请求
	Err    error                        // 错误
}

func NewWxaGetWxaCodeUnLimitResult(result WxaGetWxaCodeUnLimitResponse, body []byte, http gorequest.Response, err error) *WxaGetWxaCodeUnLimitResult {
	return &WxaGetWxaCodeUnLimitResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaGetWxaCodeUnLimit 获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (app *App) WxaGetWxaCodeUnLimit(notMustParams ...Params) *WxaGetWxaCodeUnLimitResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s", app.GetAuthorizerAccessToken()), params, http.MethodPost)
	// 定义
	var response WxaGetWxaCodeUnLimitResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewWxaGetWxaCodeUnLimitResult(response, request.ResponseBody, request, err)
}

// ErrcodeInfo 错误描述
func (resp *WxaGetWxaCodeUnLimitResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 45009:
		return "调用分钟频率受限(目前5000次/分钟，会调整)，如需大量小程序码，建议预生成"
	case 41030:
		return "page 不合法（页面不存在或者小程序没有发布、根路径前加 /或者携带参数）"
	case 40097:
		return "env_version 不合法"
	}
	return "系统繁忙"
}
