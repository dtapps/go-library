package eastiot

import (
	"encoding/json"
	"go.dtapp.net/library/gorequest"
	"net/http"
)

type IotApiQueryUserBalanceResponse struct {
	Code int `json:"code"`
	Data struct {
		Balance float64 `json:"balance"`
	} `json:"data"`
	Msg string `json:"msg"`
}

type IotApiQueryUserBalanceResult struct {
	Result IotApiQueryUserBalanceResponse // 结果
	Body   []byte                         // 内容
	Http   gorequest.Response             // 请求
	Err    error                          // 错误
}

func NewIotApiQueryUserBalanceResult(result IotApiQueryUserBalanceResponse, body []byte, http gorequest.Response, err error) *IotApiQueryUserBalanceResult {
	return &IotApiQueryUserBalanceResult{Result: result, Body: body, Http: http, Err: err}
}

// IotApiQueryUserBalance 余额查询
// https://www.showdoc.com.cn/916774523755909/4857910459512420
func (app *App) IotApiQueryUserBalance() *IotApiQueryUserBalanceResult {
	// 请求
	request, err := app.request("http://m2m.eastiot.net/Api/IotApi/queryUserBalance", map[string]interface{}{}, http.MethodPost)
	// 定义
	var response IotApiQueryUserBalanceResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewIotApiQueryUserBalanceResult(response, request.ResponseBody, request, err)
}
