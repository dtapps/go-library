package wechatpayapiv3

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type PayTransactionsJsapiResponse struct {
	PrepayId string `json:"prepay_id"`
}

// PayTransactionsJsapi JSAPI/小程序下单
// https://pay.weixin.qq.com/doc/v3/merchant/4012791897
func (c *Client) PayTransactionsJsapi(ctx context.Context, notMustParams ...*gorequest.Params) (response PayTransactionsJsapiResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", c.GetAppId()) // 【公众账号ID】是商户在微信开放平台（移动应用）或公众平台（公众号/小程序）上申请的一个唯一标识
	params.Set("mchid", c.GetMchId()) // 【商户号】是由微信支付系统生成并分配给每个商户的唯一标识符

	// 请求
	err = c.DoRequest(ctx, "v3/pay/transactions/jsapi", params, http.MethodPost, &response, &apiError)
	return
}
