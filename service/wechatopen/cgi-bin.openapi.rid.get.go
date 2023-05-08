package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type CgiBinOpenapiRidGetResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Request struct {
		InvokeTime   int    `json:"invoke_time"`   // 发起请求的时间戳
		CostInMs     int    `json:"cost_in_ms"`    // 请求毫秒级耗时
		RequestUrl   string `json:"request_url"`   // 请求的URL参数
		RequestBody  string `json:"request_body"`  // post请求的请求参数
		ResponseBody string `json:"response_body"` // 接口请求返回参数
		ClientIp     string `json:"client_ip"`     // 接口请求的客户端ip
	} `json:"request"`
}

type CgiBinOpenapiRidGetResult struct {
	Result CgiBinOpenapiRidGetResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
}

func newCgiBinOpenapiRidGetResult(result CgiBinOpenapiRidGetResponse, body []byte, http gorequest.Response) *CgiBinOpenapiRidGetResult {
	return &CgiBinOpenapiRidGetResult{Result: result, Body: body, Http: http}
}

// CgiBinOpenapiRidGet 查询rid信息
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/openapi/getRidInfo.html
func (c *Client) CgiBinOpenapiRidGet(ctx context.Context, rid string, notMustParams ...gorequest.Params) (*CgiBinOpenapiRidGetResult, error) {
	// 检查
	if err := c.checkAuthorizerConfig(ctx); err != nil {
		return newCgiBinOpenapiRidGetResult(CgiBinOpenapiRidGetResponse{}, []byte{}, gorequest.Response{}), err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	if rid != "" {
		params.Set("rid", rid)
	}
	// 请求
	request, err := c.request(ctx, apiUrl+"/cgi-bin/openapi/rid/get?access_token="+GetAuthorizerAccessToken(ctx, c), params, http.MethodPost)
	if err != nil {
		return newCgiBinOpenapiRidGetResult(CgiBinOpenapiRidGetResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response CgiBinOpenapiRidGetResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newCgiBinOpenapiRidGetResult(response, request.ResponseBody, request), err
}
