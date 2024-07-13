package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WxaSubmitAuditResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Auditid int64  `json:"auditid"`
}

type WxaSubmitAuditResult struct {
	Result WxaSubmitAuditResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newWxaSubmitAuditResult(result WxaSubmitAuditResponse, body []byte, http gorequest.Response) *WxaSubmitAuditResult {
	return &WxaSubmitAuditResult{Result: result, Body: body, Http: http}
}

// WxaSubmitAudit 提交审核
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/submit_audit.html
func (c *Client) WxaSubmitAudit(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*WxaSubmitAuditResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxa/submit_audit")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response WxaSubmitAuditResponse
	request, err := c.request(ctx, span, "wxa/submit_audit?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newWxaSubmitAuditResult(response, request.ResponseBody, request), err
}
