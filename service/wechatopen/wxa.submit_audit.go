package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type SubmitAuditResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Auditid int64  `json:"auditid"`
}

type SubmitAuditResult struct {
	Result SubmitAuditResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newSubmitAuditResult(result SubmitAuditResponse, body []byte, http gorequest.Response) *SubmitAuditResult {
	return &SubmitAuditResult{Result: result, Body: body, Http: http}
}

// SubmitAudit 提交代码审核
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/submitAudit.html
func (c *Client) SubmitAudit(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*SubmitAuditResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxa/submit_audit")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response SubmitAuditResponse
	request, err := c.request(ctx, span, "wxa/submit_audit?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newSubmitAuditResult(response, request.ResponseBody, request), err
}
