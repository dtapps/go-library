package wechatpayopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PayPartnerTransactionsH5Response struct {
	H5Url string `json:"h5_url"` // 支付跳转链接
}

type PayPartnerTransactionsH5Result struct {
	Result PayPartnerTransactionsH5Response // 结果
	Body   []byte                           // 内容
	Http   gorequest.Response               // 请求
}

func newPayPartnerTransactionsH5Result(result PayPartnerTransactionsH5Response, body []byte, http gorequest.Response) *PayPartnerTransactionsH5Result {
	return &PayPartnerTransactionsH5Result{Result: result, Body: body, Http: http}
}

// PayPartnerTransactionsH5 H5下单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_3_1.shtml
func (c *Client) PayPartnerTransactionsH5(ctx context.Context, notMustParams ...gorequest.Params) (*PayPartnerTransactionsH5Result, ApiError, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "v3/pay/partner/transactions/h5")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sp_appid", c.GetSpAppid())   // 服务商应用ID
	params.Set("sp_mchid", c.GetSpMchId())   // 服务商户号
	params.Set("sub_appid", c.GetSubAppid()) // 子商户应用ID
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号

	// 请求
	var response PayPartnerTransactionsH5Response
	var apiError ApiError
	request, err := c.request(ctx, "v3/pay/partner/transactions/h5", params, http.MethodPost, &response, &apiError)
	return newPayPartnerTransactionsH5Result(response, request.ResponseBody, request), apiError, err
}
