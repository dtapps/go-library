package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WxaSetWebViewDoMainResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type WxaSetWebViewDoMainResult struct {
	Result WxaSetWebViewDoMainResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
}

func newWxaSetWebViewDoMainResult(result WxaSetWebViewDoMainResponse, body []byte, http gorequest.Response) *WxaSetWebViewDoMainResult {
	return &WxaSetWebViewDoMainResult{Result: result, Body: body, Http: http}
}

// WxaSetWebViewDoMain 配置小程序业务域名
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/domain-management/modifyJumpDomain.html
func (c *Client) WxaSetWebViewDoMain(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*WxaSetWebViewDoMainResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxa/setwebviewdomain")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response WxaSetWebViewDoMainResponse
	request, err := c.request(ctx, span, "wxa/setwebviewdomain?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newWxaSetWebViewDoMainResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaSetWebViewDoMainResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 40001:
		return "获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的公众号调用接口"
	case 89019:
		return "业务域名无更改，无需重复设置"
	case 89020:
		return "尚未设置小程序业务域名，请先在第三方平台中设置小程序业务域名后在调用本接口"
	case 89021:
		return "请求保存的域名不是第三方平台中已设置的小程序业务域名或子域名"
	case 89029:
		return "业务域名数量超过限制，最多可以添加100个业务域名"
	case 89231:
		return "个人小程序不支持调用 setwebviewdomain 接口"
	default:
		return resp.Result.Errmsg
	}
}
