package wikeyun

import "encoding/json"

type RestPowerPushOrderResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		OrderNumber string `json:"order_number"`
	} `json:"data"`
}

type RestPowerPushOrderResult struct {
	Result RestPowerPushOrderResponse // 结果
	Body   []byte                     // 内容
	Err    error                      // 错误
}

func NewRestPowerPushOrderResult(result RestPowerPushOrderResponse, body []byte, err error) *RestPowerPushOrderResult {
	return &RestPowerPushOrderResult{Result: result, Body: body, Err: err}
}

// RestPowerPushOrder 充值下单
func (app *App) RestPowerPushOrder(notMustParams ...Params) *RestPowerPushOrderResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request("https://router.wikeyun.cn/rest/Power/pushOrder", params)
	// 定义
	var response RestPowerPushOrderResponse
	err = json.Unmarshal(body, &response)
	return NewRestPowerPushOrderResult(response, body, err)
}
