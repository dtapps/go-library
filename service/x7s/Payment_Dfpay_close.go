package x7s

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type PaymentDfpayCloseResponse struct {
	Code   int    `json:"code"`             // 0=错误 1=成功
	Msg    string `json:"msg"`              // 描述
	Status int    `json:"status,omitempty"` // 状态 -1=删除 0=可用 90=关闭
}

type PaymentDfpayCloseResult struct {
	Result PaymentDfpayCloseResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newPaymentDfpayCloseResult(result PaymentDfpayCloseResponse, body []byte, http gorequest.Response) *PaymentDfpayCloseResult {
	return &PaymentDfpayCloseResult{Result: result, Body: body, Http: http}
}

// PaymentDfpayClose 取消订单
// https://gys.x7s.com/Home_Index_documenta.html#doc9
func (c *Client) PaymentDfpayClose(ctx context.Context, partnerOrderNo string, notMustParams ...*gorequest.Params) (*PaymentDfpayCloseResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("partner_id", c.GetPartnerID())     // 供应商ID
	params.Set("partner_order_no", partnerOrderNo) // 供应商订单号

	// 响应
	var response PaymentDfpayCloseResponse

	// 请求
	request, err := c.request(ctx, "Payment_Dfpay_close", params, &response)
	return newPaymentDfpayCloseResult(response, request.ResponseBody, request), err
}
