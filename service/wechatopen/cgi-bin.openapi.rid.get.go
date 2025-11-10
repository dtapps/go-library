package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type CgiBinOpenapiRidGetResponse struct {
	APIResponse // 错误
	Request     struct {
		InvokeTime   int    `json:"invoke_time"`   // 发起请求的时间戳
		CostInMs     int    `json:"cost_in_ms"`    // 请求毫秒级耗时
		RequestUrl   string `json:"request_url"`   // 请求的URL参数
		RequestBody  string `json:"request_body"`  // post请求的请求参数
		ResponseBody string `json:"response_body"` // 接口请求返回参数
		ClientIp     string `json:"client_ip"`     // 接口请求的客户端ip
	} `json:"request"`
}

// CgiBinOpenapiRidGet 查询rid信息
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/openapi/getRidInfo.html
func (c *Client) CgiBinOpenapiRidGet(ctx context.Context, authorizerAccessToken, rid string, notMustParams ...*gorequest.Params) (response CgiBinOpenapiRidGetResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	if rid != "" {
		params.Set("rid", rid)
	}

	// 请求
	err = c.request(ctx, "cgi-bin/openapi/rid/get?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}
