package wechatopen

import (
	"encoding/json"
	"fmt"
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
	Err    error                  // 错误
}

func newWxaSubmitAuditResult(result WxaSubmitAuditResponse, body []byte, http gorequest.Response, err error) *WxaSubmitAuditResult {
	return &WxaSubmitAuditResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaSubmitAudit 提交审核
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/submit_audit.html
func (c *Client) WxaSubmitAudit(notMustParams ...gorequest.Params) *WxaSubmitAuditResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(fmt.Sprintf(apiUrl+"/wxa/submit_audit?access_token=%s", c.GetAuthorizerAccessToken()), params, http.MethodPost)
	// 定义
	var response WxaSubmitAuditResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newWxaSubmitAuditResult(response, request.ResponseBody, request, err)
}
