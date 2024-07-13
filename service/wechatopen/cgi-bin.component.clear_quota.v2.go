package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type ClearComponentQuotaByAppSecretResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type ClearComponentQuotaByAppSecretResult struct {
	Result ClearComponentQuotaByAppSecretResponse // 结果
	Body   []byte                                 // 内容
	Http   gorequest.Response                     // 请求
}

func newClearComponentQuotaByAppSecretResult(result ClearComponentQuotaByAppSecretResponse, body []byte, http gorequest.Response) *ClearComponentQuotaByAppSecretResult {
	return &ClearComponentQuotaByAppSecretResult{Result: result, Body: body, Http: http}
}

// ClearComponentQuotaByAppSecret 使用AppSecret重置第三方平台 API 调用次数
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/openapi/clearComponentQuotaByAppSecret.html
func (c *Client) ClearComponentQuotaByAppSecret(ctx context.Context, authorizerAccessToken, appid string, notMustParams ...gorequest.Params) (*ClearComponentQuotaByAppSecretResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "cgi-bin/component/clear_quota/v2")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", appid)
	params.Set("component_appid", c.GetComponentAppId())
	params.Set("appsecret", c.GetComponentAppSecret())

	// 请求
	var response ClearComponentQuotaByAppSecretResponse
	request, err := c.request(ctx, span, "cgi-bin/component/clear_quota/v2?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newClearComponentQuotaByAppSecretResult(response, request.ResponseBody, request), err
}
