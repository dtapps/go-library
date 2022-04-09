package wikeyun

import "encoding/json"

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
	Err    error                  // 错误
}

func NewRestPowerQueryResult(result RestPowerQueryResponse, body []byte, err error) *RestPowerQueryResult {
	return &RestPowerQueryResult{Result: result, Body: body, Err: err}
}

func (app *App) RestPowerQuery(orderNumber string) *RestPowerQueryResult {
	// 参数
	param := NewParams()
	param.Set("order_number", orderNumber) // 官方订单号
	params := app.NewParamsWith(param)
	// 请求
	body, err := app.request("https://router.wikeyun.cn/rest/Power/query", params)
	// 定义
	var response RestPowerQueryResponse
	err = json.Unmarshal(body, &response)
	return NewRestPowerQueryResult(response, body, err)
}
