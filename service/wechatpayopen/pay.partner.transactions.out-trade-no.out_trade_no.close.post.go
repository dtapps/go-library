package wechatpayopen

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type PayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult struct {
	Body     []byte             // 内容
	Http     gorequest.Response // 请求
	Err      error              // 错误
	ApiError ApiError           // 接口错误
}

func newPayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult(body []byte, http gorequest.Response, err error, apiError ApiError) *PayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult {
	return &PayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult{Body: body, Http: http, Err: err, ApiError: apiError}
}

// PayPartnerTransactionsOutTradeNoOutTradeNoClosePost 关闭订单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_3.shtml
func (c *Client) PayPartnerTransactionsOutTradeNoOutTradeNoClosePost(ctx context.Context, outTradeNo string, notMustParams ...gorequest.Params) *PayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sp_mchid", c.GetSpMchId())   // 服务商户号
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/pay/partner/transactions/out-trade-no/"+outTradeNo+"/close", params, http.MethodPost)
	if err != nil {
		return newPayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult(request.ResponseBody, request, err, ApiError{})
	}
	// 错误
	var apiError ApiError
	err = json.Unmarshal(request.ResponseBody, &apiError)
	return newPayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult(request.ResponseBody, request, err, apiError)
}
