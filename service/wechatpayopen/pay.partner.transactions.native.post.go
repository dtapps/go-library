package wechatpayopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type PayPartnerTransactionsNativePostResponse struct {
	CodeUrl string `json:"code_url"` // 二维码链接
}

// PayPartnerTransactionsNativePost Native下单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_1.shtml
func (c *Client) PayPartnerTransactionsNativePost(ctx context.Context, notMustParams ...*gorequest.Params) (response PayPartnerTransactionsNativePostResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sp_appid", c.GetSpAppid())   // 服务商应用ID
	params.Set("sp_mchid", c.GetSpMchId())   // 服务商户号
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号

	// 请求
	err = c.request(ctx, "/v3/pay/partner/transactions/native", params, http.MethodPost, &response, &apiError)
	return
}
