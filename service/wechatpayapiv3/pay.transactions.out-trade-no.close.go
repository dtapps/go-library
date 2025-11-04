package wechatpayapiv3

import (
	"context"
	"fmt"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// PayTransactionsOutTradeNoClose 关闭订单API https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_3.shtml
func (c *Client) PayTransactionsOutTradeNoClose(ctx context.Context, OutTradeNo string, notMustParams ...*gorequest.Params) (apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("mchid", c.GetMchId())

	// 	请求
	err = c.DoRequest(ctx, fmt.Sprintf("/v3/pay/transactions/out-trade-no/%s/close", OutTradeNo), params, http.MethodPost, false, nil)
	return
}
