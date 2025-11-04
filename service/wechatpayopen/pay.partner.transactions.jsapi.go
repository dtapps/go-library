package wechatpayopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type PayPartnerTransactionsJsapiResponse struct {
	PrepayId string `json:"prepay_id"`
}

// PayPartnerTransactionsJsapi JSAPI下单
// https://pay.weixin.qq.com/doc/v3/partner/4012738519
func (c *Client) PayPartnerTransactionsJsapi(ctx context.Context, notMustParams ...*gorequest.Params) (response PayPartnerTransactionsJsapiResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sp_appid", c.GetSpAppid())   // 服务商应用ID
	params.Set("sp_mchid", c.GetSpMchId())   // 服务商户号
	params.Set("sub_appid", c.GetSubAppid()) // 子商户应用ID
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号

	// 请求
	err = c.request(ctx, "/v3/pay/partner/transactions/jsapi", params, http.MethodPost, &response, &apiError)
	return
}
