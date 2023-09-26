package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
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
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/submit_audit?access_token="+authorizerAccessToken, params, http.MethodPost)
	if err != nil {
		return newWxaSubmitAuditResult(WxaSubmitAuditResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaSubmitAuditResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaSubmitAuditResult(response, request.ResponseBody, request), err
}
