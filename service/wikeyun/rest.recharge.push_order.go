package wikeyun

import (
	"encoding/json"
)

type RestRechargePushOrderResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderNumber string `json:"order_number"`
	} `json:"data"`
}

type RestRechargePushOrderResult struct {
	Result RestRechargePushOrderResponse // 结果
	Body   []byte                        // 内容
	Err    error                         // 错误
}

func NewRestRechargePushOrderResult(result RestRechargePushOrderResponse, body []byte, err error) *RestRechargePushOrderResult {
	return &RestRechargePushOrderResult{Result: result, Body: body, Err: err}
}

// RestRechargePushOrder 充值请求业务参数
func (app *App) RestRechargePushOrder(notMustParams ...Params) *RestRechargePushOrderResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request("https://router.wikeyun.cn/rest/Recharge/pushOrder", params)
	// 定义
	var response RestRechargePushOrderResponse
	err = json.Unmarshal(body, &response)
	return NewRestRechargePushOrderResult(response, body, err)
}
