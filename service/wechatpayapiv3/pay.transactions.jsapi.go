package wechatpayapiv3

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type PayTransactionsJsapiResponse struct {
	PrepayId string `json:"prepay_id"`
}

// PayTransactionsJsapi JSAPI下单
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_5_1.shtml
func (c *Client) PayTransactionsJsapi(ctx context.Context, notMustParams ...*gorequest.Params) (response PayTransactionsJsapiResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.DoRequest(ctx, "v3/pay/transactions/jsapi", params, http.MethodPost, true, &response)
	return
}
