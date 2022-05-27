package wechatpayapiv3

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PayTransactionsJsapiResponse struct {
	PrepayId string `json:"prepay_id"`
}

type PayTransactionsJsapiResult struct {
	Result PayTransactionsJsapiResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
	Err    error                        // 错误
}

func NewPayTransactionsJsapiResult(result PayTransactionsJsapiResponse, body []byte, http gorequest.Response, err error) *PayTransactionsJsapiResult {
	return &PayTransactionsJsapiResult{Result: result, Body: body, Http: http, Err: err}
}

// PayTransactionsJsapi 小程序 JSAPI下单
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_5_1.shtml
func (app *App) PayTransactionsJsapi(notMustParams ...Params) *PayTransactionsJsapiResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request("https://api.mch.weixin.qq.com/v3/pay/transactions/jsapi", params, http.MethodPost, true)
	if err != nil {
		return NewPayTransactionsJsapiResult(PayTransactionsJsapiResponse{}, request.ResponseBody, request, err)
	}
	// 定义
	var response PayTransactionsJsapiResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewPayTransactionsJsapiResult(response, request.ResponseBody, request, err)
}
