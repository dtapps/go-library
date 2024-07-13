package wechatpayopen

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PayPartnerTransactionsOutTradeNoCloseResult struct {
	Body []byte             // 内容
	Http gorequest.Response // 请求
}

func newPayPartnerTransactionsOutTradeNoCloseResult(body []byte, http gorequest.Response) *PayPartnerTransactionsOutTradeNoCloseResult {
	return &PayPartnerTransactionsOutTradeNoCloseResult{Body: body, Http: http}
}

// PayPartnerTransactionsOutTradeNoClose 关闭订单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_5_3.shtml
func (c *Client) PayPartnerTransactionsOutTradeNoClose(ctx context.Context, outTradeNo string, notMustParams ...gorequest.Params) (*PayPartnerTransactionsOutTradeNoCloseResult, ApiError, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, fmt.Sprintf("v3/pay/partner/transactions/out-trade-no/%s/close", outTradeNo))
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sp_mchid", c.GetSpMchId())   // 服务商户号
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号

	// 请求
	var apiError ApiError
	request, err := c.request(ctx, fmt.Sprintf("v3/pay/partner/transactions/out-trade-no/%s/close", outTradeNo), params, http.MethodPost, nil, &apiError)
	return newPayPartnerTransactionsOutTradeNoCloseResult(request.ResponseBody, request), apiError, err
}
