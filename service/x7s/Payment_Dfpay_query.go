package x7s

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type PaymentDfpayQueryResponse struct {
	Code int    `json:"code"` // 0=错误 1=成功
	Msg  string `json:"msg"`  // 描述
	Data struct {
		Status         string `json:"status"`           // 状态 -1=删除 0=可用 90=关闭
		PartnerId      string `json:"partner_id"`       // 供应商ID
		PartnerOrderNo string `json:"partner_order_no"` // 供应商订单号
		Account        string `json:"account"`          // 账号
		Amount         string `json:"amount"`           // 金额
		ChargeAmount   string `json:"charge_amount"`    // 收费金额
	} `json:"data,omitempty"`
}

type PaymentDfpayQueryResult struct {
	Result PaymentDfpayQueryResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newPaymentDfpayQueryResult(result PaymentDfpayQueryResponse, body []byte, http gorequest.Response) *PaymentDfpayQueryResult {
	return &PaymentDfpayQueryResult{Result: result, Body: body, Http: http}
}

// PaymentDfpayQuery 核销查询接口
// https://gys.x7s.com/Home_Index_documenta.html#doc8
func (c *Client) PaymentDfpayQuery(ctx context.Context, partnerOrderNo string, notMustParams ...gorequest.Params) (*PaymentDfpayQueryResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("partner_id", c.GetPartnerID())     // 供应商ID
	params.Set("partner_order_no", partnerOrderNo) // 供应商订单号

	// 响应
	var response PaymentDfpayQueryResponse

	// 请求
	request, err := c.request(ctx, "Payment_Dfpay_query", params, &response)
	return newPaymentDfpayQueryResult(response, request.ResponseBody, request), err
}
