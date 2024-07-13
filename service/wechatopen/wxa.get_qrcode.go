package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WxaGetQrcodeResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type WxaGetQrcodeResult struct {
	Result WxaGetQrcodeResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
}

func newWxaGetQrcodeResult(result WxaGetQrcodeResponse, body []byte, http gorequest.Response) *WxaGetQrcodeResult {
	return &WxaGetQrcodeResult{Result: result, Body: body, Http: http}
}

// WxaGetQrcode 获取体验版二维码
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_qrcode.html
func (c *Client) WxaGetQrcode(ctx context.Context, authorizerAccessToken, path string, notMustParams ...gorequest.Params) (*WxaGetQrcodeResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxa/get_qrcode")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	if path != "" {
		params.Set("path", path) // 指定二维码扫码后直接进入指定页面并可同时带上参数）
	}

	// 请求
	var response WxaGetQrcodeResponse
	request, err := c.request(ctx, span, "wxa/get_qrcode?access_token="+authorizerAccessToken, params, http.MethodGet, &response)

	// 判断内容是否为图片
	//if request.HeaderIsImg() == false {
	//	err = json.Unmarshal(request.ResponseBody, &response)
	//}
	return newWxaGetQrcodeResult(response, request.ResponseBody, request), err
}
