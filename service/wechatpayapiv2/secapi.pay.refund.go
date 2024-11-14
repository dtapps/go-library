package wechatpayapiv2

import (
	"context"
	"go.dtapp.net/library/utils/gorandom"
	"go.dtapp.net/library/utils/gorequest"
)

type SecApiPayRefundResponse struct {
	ReturnCode string `json:"return_code" xml:"return_code"`                   // 返回状态码
	ReturnMsg  string `json:"return_msg,omitempty" xml:"return_msg,omitempty"` // 返回信息

	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`   // 业务结果
	ErrCode    string `json:"err_code,omitempty" xml:"err_code,omitempty"`         // 错误代码
	ErrCodeDes string `json:"err_code_des,omitempty" xml:"err_code_des,omitempty"` // 错误代码描述

	Appid      string `json:"appid,omitempty" xml:"appid,omitempty"`             // 小程序ID
	MchId      string `json:"mch_id,omitempty" xml:"mch_id,omitempty"`           // 商户号
	DeviceInfo string `json:"device_info,omitempty" xml:"device_info,omitempty"` // 设备号
	NonceStr   string `json:"nonce_str,omitempty" xml:"nonce_str,omitempty"`     // 随机字符串
	Sign       string `json:"sign,omitempty" xml:"sign,omitempty"`               // 签名

	TotalRefundCount    int    `json:"total_refund_count,omitempty" xml:"total_refund_count,omitempty"`       // 订单总退款次数
	TransactionId       string `json:"transaction_id,omitempty" xml:"transaction_id,omitempty"`               // 微信订单号
	OutTradeNo          string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`                   // 商户订单号
	OutRefundNo         string `json:"out_refund_no,omitempty" xml:"out_refund_no,omitempty"`                 // 商户退款单号
	RefundId            string `json:"refund_id,omitempty" xml:"refund_id,omitempty"`                         // 微信退款单号
	RefundFee           int    `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`                       // 退款金额
	SettlementRefundFee int    `json:"settlement_refund_fee,omitempty" xml:"settlement_refund_fee,omitempty"` // 退款金额
	TotalFee            int    `json:"total_fee,omitempty" xml:"total_fee,omitempty"`                         // 标价金额
	SettlementTotalFee  int    `json:"settlement_total_fee,omitempty" xml:"settlement_total_fee,omitempty"`   // 应结订单金额
	FeeType             string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`                           // 货币种类
	CashFee             int    `json:"cash_fee,omitempty" xml:"cash_fee,omitempty"`                           // 现金支付金额
	CashFeeType         string `json:"cash_fee_type,omitempty" xml:"cash_fee_type,omitempty"`                 // 现金支付币种
	CashRefundFee       int    `json:"cash_refund_fee,omitempty" xml:"cash_refund_fee,omitempty"`             // 现金退款金额
	CouponType          string `json:"coupon_type,omitempty" xml:"coupon_type,omitempty"`                     // 代金券类型
	CouponRefundFee     int    `json:"coupon_refund_fee,omitempty" xml:"coupon_refund_fee,omitempty"`         // 代金券退款总金额
	CouponRefundCount   int    `json:"coupon_refund_count,omitempty" xml:"coupon_refund_count,omitempty"`     // 退款代金券使用数量
	CouponRefundId      string `json:"coupon_refund_id,omitempty" xml:"coupon_refund_id,omitempty"`           // 退款代金券ID
}

type SecApiPayRefundResult struct {
	Result SecApiPayRefundResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
}

func newSecApiPayRefundResult(result SecApiPayRefundResponse, body []byte, http gorequest.Response) *SecApiPayRefundResult {
	return &SecApiPayRefundResult{Result: result, Body: body, Http: http}
}

// SecApiPayRefund
// 小程序支付 - 申请退款
// 需要证书
// https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_4
func (c *Client) SecApiPayRefund(ctx context.Context, notMustParams ...*gorequest.Params) (*SecApiPayRefundResult, error) {

	// 证书
	cert, err := c.P12ToPem()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", c.GetAppId())                  // 小程序ID
	params.Set("mch_id", c.GetMchId())                 // 商户号
	params.Set("nonce_str", gorandom.Alphanumeric(32)) // 随机字符串

	// 签名
	params.Set("sign", c.getMd5Sign(params))

	// 	请求
	var response SecApiPayRefundResponse
	request, err := c.request(ctx, "secapi/pay/refund", params, true, cert, &response)
	return newSecApiPayRefundResult(response, request.ResponseBody, request), err
}
