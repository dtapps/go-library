package wechatpayopen

import (
	"context"
	"fmt"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// PayPartnerTransactionsOutTradeNoOutTradeNoClosePost 关闭订单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_3.shtml
func (c *Client) PayPartnerTransactionsOutTradeNoOutTradeNoClosePost(ctx context.Context, outTradeNo string, notMustParams ...*gorequest.Params) (apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sp_mchid", c.GetSpMchId())   // 服务商户号
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号

	// 请求
	err = c.request(ctx, fmt.Sprintf("v3/pay/partner/transactions/out-trade-no/%s/close", outTradeNo), params, http.MethodPost, nil, &apiError)
	return
}
