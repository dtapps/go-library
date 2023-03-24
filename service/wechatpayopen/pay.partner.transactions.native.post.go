package wechatpayopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type PayPartnerTransactionsNativePostResponse struct {
	CodeUrl string `json:"code_url"` // 二维码链接
}

type PayPartnerTransactionsNativePostResult struct {
	Result   PayPartnerTransactionsNativePostResponse // 结果
	Body     []byte                                   // 内容
	Http     gorequest.Response                       // 请求
	Err      error                                    // 错误
	ApiError ApiError                                 // 接口错误
}

func newPayPartnerTransactionsNativePostResult(result PayPartnerTransactionsNativePostResponse, body []byte, http gorequest.Response, err error, apiError ApiError) *PayPartnerTransactionsNativePostResult {
	return &PayPartnerTransactionsNativePostResult{Result: result, Body: body, Http: http, Err: err, ApiError: apiError}
}

// PayPartnerTransactionsNativePost Native下单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_1.shtml
func (c *Client) PayPartnerTransactionsNativePost(ctx context.Context, notMustParams ...gorequest.Params) *PayPartnerTransactionsNativePostResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sp_appid", c.GetSpAppid())   // 服务商应用ID
	params.Set("sp_mchid", c.GetSpMchId())   // 服务商户号
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/pay/partner/transactions/native", params, http.MethodPost)
	if err != nil {
		return newPayPartnerTransactionsNativePostResult(PayPartnerTransactionsNativePostResponse{}, request.ResponseBody, request, err, ApiError{})
	}
	// 定义
	var response PayPartnerTransactionsNativePostResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = gojson.Unmarshal(request.ResponseBody, &apiError)
	return newPayPartnerTransactionsNativePostResult(response, request.ResponseBody, request, err, apiError)
}
