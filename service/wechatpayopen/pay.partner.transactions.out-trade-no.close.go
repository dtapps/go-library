package wechatpayopen

import (
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PayPartnerTransactionsOutTradeNoCloseResult struct {
	Body []byte             // 内容
	Http gorequest.Response // 请求
	Err  error              // 错误
}

func newPayPartnerTransactionsOutTradeNoCloseResult(body []byte, http gorequest.Response, err error) *PayPartnerTransactionsOutTradeNoCloseResult {
	return &PayPartnerTransactionsOutTradeNoCloseResult{Body: body, Http: http, Err: err}
}

// PayPartnerTransactionsOutTradeNoClose 关闭订单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_5_3.shtml
func (c *Client) PayPartnerTransactionsOutTradeNoClose(outTradeNo string) *PayPartnerTransactionsOutTradeNoCloseResult {
	// 参数
	params := gorequest.NewParams()
	params.Set("sp_mchid", c.config.SpMchId)   // 服务商户号
	params.Set("sub_mchid", c.config.SubMchId) // 子商户号
	// 请求
	request, err := c.request(fmt.Sprintf(apiUrl+"/v3/pay/partner/transactions/out-trade-no/%s/close", outTradeNo), params, http.MethodPost)
	if err != nil {
		return newPayPartnerTransactionsOutTradeNoCloseResult(request.ResponseBody, request, err)
	}
	return newPayPartnerTransactionsOutTradeNoCloseResult(request.ResponseBody, request, err)
}
