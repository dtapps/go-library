package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type BindTesterResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
	Userstr string `json:"userstr"` // 人员对应的唯一字符串
}

type BindTesterResult struct {
	Result BindTesterResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newBindTesterResult(result BindTesterResponse, body []byte, http gorequest.Response) *BindTesterResult {
	return &BindTesterResult{Result: result, Body: body, Http: http}
}

// BindTester 绑定体验者
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/member-management/bindTester.html
func (c *Client) BindTester(ctx context.Context, authorizerAccessToken, wechatid string, notMustParams ...gorequest.Params) (*BindTesterResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxa/bind_tester")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("wechatid", wechatid)

	// 请求
	var response BindTesterResponse
	request, err := c.request(ctx, span, "wxa/bind_tester?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newBindTesterResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *BindTesterResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 85001:
		return "微信号不存在或微信号设置为不可搜索"
	case 85002:
		return "小程序绑定的体验者数量达到上限"
	case 85003:
		return "微信号绑定的小程序体验者达到上限"
	case 85004:
		return "微信号已经绑定"
	default:
		return resp.Result.Errmsg
	}
}
