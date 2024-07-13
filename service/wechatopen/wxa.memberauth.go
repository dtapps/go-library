package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WxaMemberAuthResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
	Members []struct {
		Userstr string `json:"userstr"` // 人员对应的唯一字符串
	} `json:"members"` // 人员信息列表
}

type WxaMemberAuthResult struct {
	Result WxaMemberAuthResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newWxaMemberAuthResult(result WxaMemberAuthResponse, body []byte, http gorequest.Response) *WxaMemberAuthResult {
	return &WxaMemberAuthResult{Result: result, Body: body, Http: http}
}

// WxaMemberAuth 获取体验者列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/memberauth.html
func (c *Client) WxaMemberAuth(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*WxaMemberAuthResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxa/memberauth")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("action", "get_experiencer")

	// 请求
	var response WxaMemberAuthResponse
	request, err := c.request(ctx, span, "wxa/memberauth?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newWxaMemberAuthResult(response, request.ResponseBody, request), err
}
