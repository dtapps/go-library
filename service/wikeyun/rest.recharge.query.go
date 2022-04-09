package wikeyun

import (
	"encoding/json"
)

type RestRechargeQueryResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		OrderNumber    string `json:"order_number"`
		OrderNo        string `json:"order_no"`
		Mobile         string `json:"mobile"`
		Amount         int    `json:"amount"`
		CostPrice      string `json:"cost_price"`
		Fanli          string `json:"fanli"`
		Status         int    `json:"status"`
		OrgOrderNumber string `json:"org_order_number"`
	} `json:"data"`
}

type RestRechargeQueryResult struct {
	Result RestRechargeQueryResponse // 结果
	Body   []byte                    // 内容
	Err    error                     // 错误
}

func NewRestRechargeQueryResult(result RestRechargeQueryResponse, body []byte, err error) *RestRechargeQueryResult {
	return &RestRechargeQueryResult{Result: result, Body: body, Err: err}
}

// RestRechargeQuery 查询接口
func (app *App) RestRechargeQuery(orderNumber string) *RestRechargeQueryResult {
	// 参数
	param := NewParams()
	param.Set("order_number", orderNumber) // 官方订单号
	params := app.NewParamsWith(param)
	// 请求
	body, err := app.request("https://router.wikeyun.cn/rest/Recharge/query", params)
	// 定义
	var response RestRechargeQueryResponse
	err = json.Unmarshal(body, &response)
	return NewRestRechargeQueryResult(response, body, err)
}
