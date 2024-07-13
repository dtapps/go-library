package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WxaMsgSecCheckResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Result  struct {
		Suggest string `json:"suggest"`
		Label   int    `json:"label"`
	} `json:"result"`
	Detail []struct {
		Strategy string `json:"strategy"`
		Errcode  int    `json:"errcode"`
		Suggest  string `json:"suggest"`
		Label    int    `json:"label"`
		Prob     int    `json:"prob,omitempty"`
		Level    int    `json:"level,omitempty"`
		Keyword  string `json:"keyword,omitempty"`
	} `json:"detail"`
	TraceId string `json:"trace_id"`
}

type WxaMsgSecCheckResult struct {
	Result WxaMsgSecCheckResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newWxaMsgSecCheckResult(result WxaMsgSecCheckResponse, body []byte, http gorequest.Response) *WxaMsgSecCheckResult {
	return &WxaMsgSecCheckResult{Result: result, Body: body, Http: http}
}

// WxaMsgSecCheck 文本内容安全识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (c *Client) WxaMsgSecCheck(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*WxaMsgSecCheckResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxa/msg_sec_check")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response WxaMsgSecCheckResponse
	request, err := c.request(ctx, span, "wxa/msg_sec_check?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newWxaMsgSecCheckResult(response, request.ResponseBody, request), err
}
