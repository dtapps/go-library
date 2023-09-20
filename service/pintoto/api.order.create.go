package pintoto

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type ApiOrderCreateResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type ApiOrderCreateResult struct {
	Result ApiOrderCreateResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
	Err    error                  // 错误
}

func newApiOrderCreateResult(result ApiOrderCreateResponse, body []byte, http gorequest.Response, err error) *ApiOrderCreateResult {
	return &ApiOrderCreateResult{Result: result, Body: body, Http: http, Err: err}
}

// ApiOrderCreate 下单api https://www.showdoc.com.cn/1154868044931571/5891022916496848
func (c *Client) ApiOrderCreate(ctx context.Context, notMustParams ...*gorequest.Params) *ApiOrderCreateResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/api/order/create", params)
	// 定义
	var response ApiOrderCreateResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newApiOrderCreateResult(response, request.ResponseBody, request, err)
}
