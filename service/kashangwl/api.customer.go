package kashangwl

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
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
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("customer_id", customerID) // 商家编号
	// 请求
	request, err := c.request(ctx, apiUrl+"/api/customer", params)
	if err != nil {
		return newApiCustomerResult(ApiCustomerResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ApiCustomerResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newApiCustomerResult(response, request.ResponseBody, request), err
}
