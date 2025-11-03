package wechatpayopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type PayPartnerTransactionsH5Response struct {
	H5Url string `json:"h5_url"` // 支付跳转链接
}

// PayPartnerTransactionsH5 H5下单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_3_1.shtml
func (c *Client) PayPartnerTransactionsH5(ctx context.Context, notMustParams ...*gorequest.Params) (response PayPartnerTransactionsH5Response, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sp_appid", c.GetSpAppid())   // 服务商应用ID
	params.Set("sp_mchid", c.GetSpMchId())   // 服务商户号
	params.Set("sub_appid", c.GetSubAppid()) // 子商户应用ID
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号

	// 请求
	err = c.request(ctx, "v3/pay/partner/transactions/h5", params, http.MethodPost, &response, &apiError)
	return
}
