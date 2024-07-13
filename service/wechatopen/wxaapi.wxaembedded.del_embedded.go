package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WxaApiWxaembeddedDelAuthorizeResponse struct {
	Errcode int    `json:"errcode"` // 返回码
	Errmsg  string `json:"errmsg"`  // 返回码信息
}

type WxaApiWxaembeddedDelAuthorizeResult struct {
	Result WxaApiWxaembeddedDelAuthorizeResponse // 结果
	Body   []byte                                // 内容
	Http   gorequest.Response                    // 请求
}

func newWxaApiWxaembeddedDelAuthorizeResult(result WxaApiWxaembeddedDelAuthorizeResponse, body []byte, http gorequest.Response) *WxaApiWxaembeddedDelAuthorizeResult {
	return &WxaApiWxaembeddedDelAuthorizeResult{Result: result, Body: body, Http: http}
}

// WxaApiWxaembeddedDelAuthorize 取消授权小程序
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/embedded-management/deleteAuthorizedEmbedded.html
func (c *Client) WxaApiWxaembeddedDelAuthorize(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*WxaApiWxaembeddedDelAuthorizeResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxaapi/wxaembedded/del_authorize")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response WxaApiWxaembeddedDelAuthorizeResponse
	request, err := c.request(ctx, span, "wxaapi/wxaembedded/del_authorize?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newWxaApiWxaembeddedDelAuthorizeResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaApiWxaembeddedDelAuthorizeResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 89416:
		return "取消半屏小程序授权 appid 参数为空"
	case 89431:
		return "不支持此类型小程序"
	case 89432:
		return "不是小程序"
	default:
		return resp.Result.Errmsg
	}
}
