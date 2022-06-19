package kashangwl

import (
	"encoding/json"
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
	Err    error               // 错误
}

func newApiCustomerResult(result ApiCustomerResponse, body []byte, http gorequest.Response, err error) *ApiCustomerResult {
	return &ApiCustomerResult{Result: result, Body: body, Http: http, Err: err}
}

// ApiCustomer 获取商家信息
// http://doc.cqmeihu.cn/sales/merchant-info.html
func (c *Client) ApiCustomer() *ApiCustomerResult {
	// 请求
	request, err := c.request(apiUrl+"/api/customer", map[string]interface{}{})
	// 定义
	var response ApiCustomerResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newApiCustomerResult(response, request.ResponseBody, request, err)
}
