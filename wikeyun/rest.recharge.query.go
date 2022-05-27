package wikeyun

import (
	"encoding/json"
	"go.dtapp.net/library/gorequest"
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
	Http   gorequest.Response        // 请求
	Err    error                     // 错误
}

func NewRestRechargeQueryResult(result RestRechargeQueryResponse, body []byte, http gorequest.Response, err error) *RestRechargeQueryResult {
	return &RestRechargeQueryResult{Result: result, Body: body, Http: http, Err: err}
}

// RestRechargeQuery 话费订单查询
// https://open.wikeyun.cn/#/apiDocument/9/document/299
func (app *App) RestRechargeQuery(orderNumber string) *RestRechargeQueryResult {
	// 参数
	param := NewParams()
	param.Set("order_number", orderNumber) // 平台订单号
	params := app.NewParamsWith(param)
	// 请求
	request, err := app.request("https://router.wikeyun.cn/rest/Recharge/query", params)
	// 定义
	var response RestRechargeQueryResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewRestRechargeQueryResult(response, request.ResponseBody, request, err)
}
