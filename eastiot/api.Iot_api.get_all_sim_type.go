package eastiot

import (
	"encoding/json"
	"go.dtapp.net/library/gorequest"
	"net/http"
)

type IotApiGetAllSimTypeResponse struct {
	Code int `json:"code"`
	Data []struct {
		Type   int    `json:"type"`   // 卡类型
		Name   string `json:"name"`   // 类型名
		MOrder int    `json:"mOrder"` // 是否支持单次充值多个流量包，0:不支持 1:支持
	} `json:"data"`
	Msg string `json:"msg"`
}

type IotApiGetAllSimTypeResult struct {
	Result IotApiGetAllSimTypeResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
	Err    error                       // 错误
}

func NewIotApiGetAllSimTypeResult(result IotApiGetAllSimTypeResponse, body []byte, http gorequest.Response, err error) *IotApiGetAllSimTypeResult {
	return &IotApiGetAllSimTypeResult{Result: result, Body: body, Http: http, Err: err}
}

// IotApiGetAllSimType 卡类型列表查询
// https://www.showdoc.com.cn/916774523755909/4858492092033167
func (app *App) IotApiGetAllSimType() *IotApiGetAllSimTypeResult {
	// 请求
	request, err := app.request("http://m2m.eastiot.net/Api/IotApi/getAllSimType", map[string]interface{}{}, http.MethodPost)
	// 定义
	var response IotApiGetAllSimTypeResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewIotApiGetAllSimTypeResult(response, request.ResponseBody, request, err)
}
