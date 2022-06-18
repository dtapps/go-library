package wechatpayopen

import (
	"fmt"
	gorequest2 "go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PayPartnerTransactionsOutTradeNoCloseResult struct {
	Body []byte              // 内容
	Http gorequest2.Response // 请求
	Err  error               // 错误
}

func NewPayPartnerTransactionsOutTradeNoCloseResult(body []byte, http gorequest2.Response, err error) *PayPartnerTransactionsOutTradeNoCloseResult {
	return &PayPartnerTransactionsOutTradeNoCloseResult{Body: body, Http: http, Err: err}
}

// PayPartnerTransactionsOutTradeNoClose 关闭订单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_5_3.shtml
func (app *App) PayPartnerTransactionsOutTradeNoClose(outTradeNo string) *PayPartnerTransactionsOutTradeNoCloseResult {
	// 参数
	params := gorequest2.NewParams()
	params.Set("sp_mchid", app.spMchId)   // 服务商户号
	params.Set("sub_mchid", app.subMchId) // 子商户号
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.mch.weixin.qq.com/v3/pay/partner/transactions/out-trade-no/%s/close", outTradeNo), params, http.MethodPost)
	if err != nil {
		return NewPayPartnerTransactionsOutTradeNoCloseResult(request.ResponseBody, request, err)
	}
	return NewPayPartnerTransactionsOutTradeNoCloseResult(request.ResponseBody, request, err)
}
