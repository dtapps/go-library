package eastiot

import (
	"encoding/json"
	"net/http"
)

type IotApiRechargeSimResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type IotApiRechargeSimResult struct {
	Result IotApiRechargeSimResponse // 结果
	Body   []byte                    // 内容
	Err    error                     // 错误
}

func NewIotApiRechargeSimResult(result IotApiRechargeSimResponse, body []byte, err error) *IotApiRechargeSimResult {
	return &IotApiRechargeSimResult{Result: result, Body: body, Err: err}
}

// IotApiRechargeSim 单卡流量充值
// https://www.showdoc.com.cn/916774523755909/4880284631482420
func (app *App) IotApiRechargeSim(notMustParams ...Params) *IotApiRechargeSimResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request("http://m2m.eastiot.net/Api/IotApi/rechargeSim", params, http.MethodPost)
	// 定义
	var response IotApiRechargeSimResponse
	err = json.Unmarshal(body, &response)
	return NewIotApiRechargeSimResult(response, body, err)
}
