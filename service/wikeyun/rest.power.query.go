package wikeyun

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
)

type RestPowerQueryResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		OrderNumber   string `json:"order_number"`
		OrderNo       string `json:"order_no"`
		CardId        string `json:"card_id"`
		Amount        int    `json:"amount"`
		CostPrice     string `json:"cost_price"`
		Fanli         string `json:"fanli"`
		Status        int    `json:"status"`
		ArrivedAmount int64  `json:"arrived_amount"`
	} `json:"data"`
}

type RestPowerQueryResult struct {
	Result RestPowerQueryResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
	Err    error                  // 错误
}

func NewRestPowerQueryResult(result RestPowerQueryResponse, body []byte, http gorequest.Response, err error) *RestPowerQueryResult {
	return &RestPowerQueryResult{Result: result, Body: body, Http: http, Err: err}
}

// RestPowerQuery 电费订单查询
// https://open.wikeyun.cn/#/apiDocument/9/document/313
func (c *Client) RestPowerQuery(orderNumber string) *RestPowerQueryResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("order_number", orderNumber) // 平台单号
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(apiUrl+"/rest/Power/query", params)
	// 定义
	var response RestPowerQueryResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewRestPowerQueryResult(response, request.ResponseBody, request, err)
}
