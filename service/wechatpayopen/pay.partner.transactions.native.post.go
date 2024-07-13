package wechatpayopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PayPartnerTransactionsNativePostResponse struct {
	CodeUrl string `json:"code_url"` // 二维码链接
}

type PayPartnerTransactionsNativePostResult struct {
	Result PayPartnerTransactionsNativePostResponse // 结果
	Body   []byte                                   // 内容
	Http   gorequest.Response                       // 请求
}

func newPayPartnerTransactionsNativePostResult(result PayPartnerTransactionsNativePostResponse, body []byte, http gorequest.Response) *PayPartnerTransactionsNativePostResult {
	return &PayPartnerTransactionsNativePostResult{Result: result, Body: body, Http: http}
}

// PayPartnerTransactionsNativePost Native下单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_1.shtml
func (c *Client) PayPartnerTransactionsNativePost(ctx context.Context, notMustParams ...gorequest.Params) (*PayPartnerTransactionsNativePostResult, ApiError, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "v3/pay/partner/transactions/native")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sp_appid", c.GetSpAppid())   // 服务商应用ID
	params.Set("sp_mchid", c.GetSpMchId())   // 服务商户号
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号

	// 请求
	var response PayPartnerTransactionsNativePostResponse
	var apiError ApiError
	request, err := c.request(ctx, "v3/pay/partner/transactions/native", params, http.MethodPost, &response, &apiError)
	return newPayPartnerTransactionsNativePostResult(response, request.ResponseBody, request), apiError, err
}
