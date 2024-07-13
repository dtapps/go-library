package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type ClearQuotaResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type ClearQuotaResult struct {
	Result ClearQuotaResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newClearQuotaResult(result ClearQuotaResponse, body []byte, http gorequest.Response) *ClearQuotaResult {
	return &ClearQuotaResult{Result: result, Body: body, Http: http}
}

// ClearQuota 重置API调用次数
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/openapi/clearQuota.html
func (c *Client) ClearQuota(ctx context.Context, authorizerAccessToken, appid string, notMustParams ...gorequest.Params) (*ClearQuotaResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "cgi-bin/clear_quota")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", appid)

	// 请求
	var response ClearQuotaResponse
	request, err := c.request(ctx, span, "cgi-bin/clear_quota?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newClearQuotaResult(response, request.ResponseBody, request), err
}
