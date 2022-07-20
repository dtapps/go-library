package wechatpayopen

import (
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type PayPartnerTransactionsH5Response struct {
	H5Url string `json:"h5_url"` // 支付跳转链接
}

type PayPartnerTransactionsH5Result struct {
	Result   PayPartnerTransactionsH5Response // 结果
	Body     []byte                           // 内容
	Http     gorequest.Response               // 请求
	Err      error                            // 错误
	ApiError ApiError                         // 接口错误
}

func newPayPartnerTransactionsH5Result(result PayPartnerTransactionsH5Response, body []byte, http gorequest.Response, err error, apiError ApiError) *PayPartnerTransactionsH5Result {
	return &PayPartnerTransactionsH5Result{Result: result, Body: body, Http: http, Err: err, ApiError: apiError}
}

// PayPartnerTransactionsH5 H5下单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_3_1.shtml
func (c *Client) PayPartnerTransactionsH5(notMustParams ...gorequest.Params) *PayPartnerTransactionsH5Result {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sp_appid", c.config.SpAppid)   // 服务商应用ID
	params.Set("sp_mchid", c.config.SpMchId)   // 服务商户号
	params.Set("sub_appid", c.config.SubAppid) // 子商户应用ID
	params.Set("sub_mchid", c.config.SubMchId) // 子商户号
	// 请求
	request, err := c.request(apiUrl+"/v3/pay/partner/transactions/h5", params, http.MethodPost)
	if err != nil {
		return newPayPartnerTransactionsH5Result(PayPartnerTransactionsH5Response{}, request.ResponseBody, request, err, ApiError{})
	}
	// 结果
	var response PayPartnerTransactionsH5Response
	err = json.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = json.Unmarshal(request.ResponseBody, &apiError)
	return newPayPartnerTransactionsH5Result(response, request.ResponseBody, request, err, apiError)
}
