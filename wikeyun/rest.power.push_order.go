package wikeyun

import (
	"encoding/json"
	"go.dtapp.net/library/gorequest"
)

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
	Http   gorequest.Response         // 请求
	Err    error                      // 错误
}

func NewRestPowerPushOrderResult(result RestPowerPushOrderResponse, body []byte, http gorequest.Response, err error) *RestPowerPushOrderResult {
	return &RestPowerPushOrderResult{Result: result, Body: body, Http: http, Err: err}
}

// RestPowerPushOrder 电费充值API
// https://open.wikeyun.cn/#/apiDocument/9/document/311
func (app *App) RestPowerPushOrder(notMustParams ...Params) *RestPowerPushOrderResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	params.Set("store_id", app.storeId) // 店铺ID
	// 请求
	request, err := app.request("https://router.wikeyun.cn/rest/Power/pushOrder", params)
	// 定义
	var response RestPowerPushOrderResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewRestPowerPushOrderResult(response, request.ResponseBody, request, err)
}
