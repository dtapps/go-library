package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WxaReleaseResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type WxaReleaseResult struct {
	Result WxaReleaseResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newWxaReleaseResult(result WxaReleaseResponse, body []byte, http gorequest.Response) *WxaReleaseResult {
	return &WxaReleaseResult{Result: result, Body: body, Http: http}
}

// WxaRelease 发布已通过审核的小程序
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/release.html
func (c *Client) WxaRelease(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*WxaReleaseResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxa/release")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response WxaReleaseResponse
	request, err := c.request(ctx, span, "wxa/release?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newWxaReleaseResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaReleaseResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 85019:
		return "没有审核版本"
	case 85020:
		return "审核状态未满足发布"
	default:
		return resp.Result.Errmsg
	}
}
