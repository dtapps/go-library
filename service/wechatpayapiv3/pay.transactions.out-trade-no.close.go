package wechatpayapiv3

import (
	"context"
	"fmt"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// PayTransactionsOutTradeNoClose 关闭订单
// https://pay.weixin.qq.com/doc/v3/merchant/4012791901
func (c *Client) PayTransactionsOutTradeNoClose(ctx context.Context, OutTradeNo string, notMustParams ...*gorequest.Params) (apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("mchid", c.GetMchId()) // 【商户号】商户下单时传入的商户号

	// 	请求
	err = c.DoRequest(ctx, fmt.Sprintf("/v3/pay/transactions/out-trade-no/%s/close", OutTradeNo), params, http.MethodPost, nil, &apiError)
	return
}
