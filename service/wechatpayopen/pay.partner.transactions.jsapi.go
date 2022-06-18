package wechatpayopen

import (
	"encoding/json"
	gorequest2 "go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PayPartnerTransactionsJsapiResponse struct {
	PrepayId string `json:"prepay_id"`
}

type PayPartnerTransactionsJsapiResult struct {
	Result PayPartnerTransactionsJsapiResponse // 结果
	Body   []byte                              // 内容
	Http   gorequest2.Response                 // 请求
	Err    error                               // 错误
}

func NewPayPartnerTransactionsJsapiResult(result PayPartnerTransactionsJsapiResponse, body []byte, http gorequest2.Response, err error) *PayPartnerTransactionsJsapiResult {
	return &PayPartnerTransactionsJsapiResult{Result: result, Body: body, Http: http, Err: err}
}

// PayPartnerTransactionsJsapi JSAPI下单
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_5_1.shtml
func (app *App) PayPartnerTransactionsJsapi(notMustParams ...gorequest2.Params) *PayPartnerTransactionsJsapiResult {
	// 参数
	params := gorequest2.NewParamsWith(notMustParams...)
	params.Set("sp_appid", app.spAppid)   // 服务商应用ID
	params.Set("sp_mchid", app.spMchId)   // 服务商户号
	params.Set("sub_appid", app.subAppid) // 子商户应用ID
	params.Set("sub_mchid", app.subMchId) // 子商户号
	// 请求
	request, err := app.request("https://api.mch.weixin.qq.com/v3/pay/partner/transactions/jsapi", params, http.MethodPost)
	if err != nil {
		return NewPayPartnerTransactionsJsapiResult(PayPartnerTransactionsJsapiResponse{}, request.ResponseBody, request, err)
	}
	// 定义
	var response PayPartnerTransactionsJsapiResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewPayPartnerTransactionsJsapiResult(response, request.ResponseBody, request, err)
}
