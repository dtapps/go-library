package x7s

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type PaymentDfpayAddorderResponse struct {
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

type PaymentDfpayAddorderResult struct {
	Result PaymentDfpayAddorderResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
}

func newPaymentDfpayAddorderResult(result PaymentDfpayAddorderResponse, body []byte, http gorequest.Response) *PaymentDfpayAddorderResult {
	return &PaymentDfpayAddorderResult{Result: result, Body: body, Http: http}
}

// PaymentDfpayAddorder 统一下单接口
// https://gys.x7s.com/Home_Index_documenta.html#doc6
func (c *Client) PaymentDfpayAddorder(ctx context.Context, partnerOrderNo string, Type int, account string, amount float64, notifyUrl string, notMustParams ...gorequest.Params) (*PaymentDfpayAddorderResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("partner_id", c.GetPartnerID())     // 供应商ID
	params.Set("partner_order_no", partnerOrderNo) // 供应商订单号
	params.Set("type", Type)                       // 类型 17-新奥燃气，16-云缴费电费，1000-国网电费，13-淘宝电费，31-话费
	params.Set("account", account)                 // 账号 电费户号、话费手机号
	params.Set("amount", amount)                   // 充值金额(元)
	params.Set("notify_url", notifyUrl)            // 回调通知地址

	// 响应
	var response PaymentDfpayAddorderResponse

	// 请求
	request, err := c.request(ctx, "Payment_Dfpay_addorder", params, &response)
	return newPaymentDfpayAddorderResult(response, request.ResponseBody, request), err
}
