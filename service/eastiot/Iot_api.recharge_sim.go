package eastiot

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type IotApiRechargeSimResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type IotApiRechargeSimResult struct {
	Result IotApiRechargeSimResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
	Err    error                     // 错误
}

func NewIotApiRechargeSimResult(result IotApiRechargeSimResponse, body []byte, http gorequest.Response, err error) *IotApiRechargeSimResult {
	return &IotApiRechargeSimResult{Result: result, Body: body, Http: http, Err: err}
}

// IotApiRechargeSim 单卡流量充值
// https://www.showdoc.com.cn/916774523755909/4880284631482420
func (app *App) IotApiRechargeSim(notMustParams ...Params) *IotApiRechargeSimResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request("http://m2m.eastiot.net/Api/IotApi/rechargeSim", params, http.MethodPost)
	// 定义
	var response IotApiRechargeSimResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewIotApiRechargeSimResult(response, request.ResponseBody, request, err)
}
