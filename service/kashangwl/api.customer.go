package kashangwl

import "encoding/json"

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
	Err    error               // 错误
}

func NewApiCustomerResult(result ApiCustomerResponse, body []byte, err error) *ApiCustomerResult {
	return &ApiCustomerResult{Result: result, Body: body, Err: err}
}

// ApiCustomer 获取商家信息
// http://doc.cqmeihu.cn/sales/merchant-info.html
func (app *App) ApiCustomer() *ApiCustomerResult {
	// 请求
	body, err := app.request("http://www.kashangwl.com/api/customer", map[string]interface{}{})
	// 定义
	var response ApiCustomerResponse
	err = json.Unmarshal(body, &response)
	return NewApiCustomerResult(response, body, err)
}
