package kashangwl

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type ApiCustomerResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Id      int    `json:"id"`      // 商家编号
		Name    string `json:"name"`    // 商家名称
		Balance string `json:"balance"` // 余额
	} `json:"data"`
}

type ApiCustomerResult struct {
	Result ApiCustomerResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newApiCustomerResult(result ApiCustomerResponse, body []byte, http gorequest.Response) *ApiCustomerResult {
	return &ApiCustomerResult{Result: result, Body: body, Http: http}
}

// ApiCustomer 获取商家信息
// customer_id = 商家编号
// http://doc.cqmeihu.cn/sales/merchant-info.html
func (c *Client) ApiCustomer(ctx context.Context, customerID int64, notMustParams ...gorequest.Params) (*ApiCustomerResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "api/customer")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("customer_id", customerID) // 商家编号

	// 请求
	var response ApiCustomerResponse
	request, err := c.request(ctx, "api/customer", params, &response)
	return newApiCustomerResult(response, request.ResponseBody, request), err
}
