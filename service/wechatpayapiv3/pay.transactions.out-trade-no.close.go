package wechatpayapiv3

import (
	"fmt"
	"net/http"
)

// PayTransactionsOutTradeNoClose 关闭订单API https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_3.shtml
func (app *App) PayTransactionsOutTradeNoClose(OutTradeNo string) (result ErrResp, err error) {
	// 参数
	params := NewParams()
	params["mchid"] = app.MchId
	// 	请求
	_, result, err = app.request(fmt.Sprintf("https://api.mch.weixin.qq.com/v3/pay/transactions/out-trade-no/%s/close", OutTradeNo), params, http.MethodPost)
	if err != nil {
		return
	}
	return
}
