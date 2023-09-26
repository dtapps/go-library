package wechatpayopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type PayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult struct {
	Body []byte             // 内容
	Http gorequest.Response // 请求
}

func newPayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult(body []byte, http gorequest.Response) *PayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult {
	return &PayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult{Body: body, Http: http}
}

// PayPartnerTransactionsOutTradeNoOutTradeNoClosePost 关闭订单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_3.shtml
func (c *Client) PayPartnerTransactionsOutTradeNoOutTradeNoClosePost(ctx context.Context, outTradeNo string, notMustParams ...gorequest.Params) (*PayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult, ApiError, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sp_mchid", c.GetSpMchId())   // 服务商户号
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/pay/partner/transactions/out-trade-no/"+outTradeNo+"/close", params, http.MethodPost)
	if err != nil {
		return newPayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult(request.ResponseBody, request), ApiError{}, err
	}
	// 错误
	var apiError ApiError
	err = gojson.Unmarshal(request.ResponseBody, &apiError)
	return newPayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult(request.ResponseBody, request), apiError, err
}
