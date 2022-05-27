package wechatpayapiv3

import (
	"fmt"
	"go.dtapp.net/library/utils/gohttp"
	"net/http"
)

type PayTransactionsOutTradeNoCloseResult struct {
	Body []byte          // 内容
	Http gohttp.Response // 请求
	Err  error           // 错误
}

func NewPayTransactionsOutTradeNoCloseResult(body []byte, http gohttp.Response, err error) *PayTransactionsOutTradeNoCloseResult {
	return &PayTransactionsOutTradeNoCloseResult{Body: body, Http: http, Err: err}
}

// PayTransactionsOutTradeNoClose 关闭订单API https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_3.shtml
func (app *App) PayTransactionsOutTradeNoClose(OutTradeNo string) *PayTransactionsOutTradeNoCloseResult {
	// 参数
	params := NewParams()
	params["mchid"] = app.MchId
	// 	请求
	request, err := app.request(fmt.Sprintf("https://api.mch.weixin.qq.com/v3/pay/transactions/out-trade-no/%s/close", OutTradeNo), params, http.MethodPost, false)
	return NewPayTransactionsOutTradeNoCloseResult(request.Body, request, err)
}
