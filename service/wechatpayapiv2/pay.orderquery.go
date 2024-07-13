package wechatpayapiv2

import (
	"context"
	"go.dtapp.net/library/utils/gorandom"
	"go.dtapp.net/library/utils/gorequest"
)

type PayOrderQueryResponse struct {
	ReturnCode string `json:"return_code" xml:"return_code"`                   // 返回状态码
	ReturnMsg  string `json:"return_msg,omitempty" xml:"return_msg,omitempty"` // 返回信息

	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`   // 业务结果
	ErrCode    string `json:"err_code,omitempty" xml:"err_code,omitempty"`         // 错误代码
	ErrCodeDes string `json:"err_code_des,omitempty" xml:"err_code_des,omitempty"` // 错误代码描述

	Appid    string `json:"appid,omitempty" xml:"appid,omitempty"`         // 小程序ID
	MchId    string `json:"mch_id,omitempty" xml:"mch_id,omitempty"`       // 商户号
	NonceStr string `json:"nonce_str,omitempty" xml:"nonce_str,omitempty"` // 随机字符串
	Sign     string `json:"sign,omitempty" xml:"sign,omitempty"`           // 签名

	DeviceInfo         string `json:"device_info,omitempty" xml:"device_info,omitempty"`                   // 设备号
	Openid             string `json:"openid,omitempty" xml:"openid,omitempty"`                             // 用户标识
	IsSubscribe        string `json:"is_subscribe,omitempty" xml:"is_subscribe,omitempty"`                 // 是否关注公众账号
	TradeType          string `json:"trade_type,omitempty" xml:"trade_type,omitempty"`                     // 交易类型
	TradeState         string `json:"trade_state,omitempty" xml:"trade_state,omitempty"`                   // 交易状态
	BankType           string `json:"bank_type,omitempty" xml:"bank_type,omitempty"`                       // 付款银行
	TotalFee           int    `json:"total_fee,omitempty" xml:"total_fee,omitempty"`                       // 标价金额
	SettlementTotalFee int    `json:"settlement_total_fee,omitempty" xml:"settlement_total_fee,omitempty"` // 应结订单金额
	FeeType            string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`                         // 标价币种
	CashFee            int    `json:"cash_fee,omitempty" xml:"cash_fee,omitempty"`                         // 现金支付金额
	CashFeeType        string `json:"cash_fee_type,omitempty" xml:"cash_fee_type,omitempty"`               // 现金支付币种
	CouponFee          int    `json:"coupon_fee,omitempty" xml:"coupon_fee,omitempty"`                     // 代金券金额
	CouponCount        int    `json:"coupon_count,omitempty" xml:"coupon_count,omitempty"`                 // 代金券使用数量
	CouponType         string `json:"coupon_type,omitempty" xml:"coupon_type,omitempty"`                   // 代金券类型
	CouponId           string `json:"coupon_id,omitempty" xml:"coupon_id,omitempty"`                       // 代金券ID
	TransactionId      string `json:"transaction_id,omitempty" xml:"transaction_id,omitempty"`             // 微信支付订单号
	OutTradeNo         string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`                 // 商户订单号
	Attach             string `json:"attach,omitempty" xml:"attach,omitempty"`                             // 附加数据
	TimeEnd            string `json:"time_end,omitempty" xml:"time_end,omitempty"`                         // 支付完成时间
	TradeStateDesc     string `json:"trade_state_desc,omitempty" xml:"trade_state_desc,omitempty"`         // 交易状态描述
}

type PayOrderQueryResult struct {
	Result PayOrderQueryResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newPayOrderQueryResult(result PayOrderQueryResponse, body []byte, http gorequest.Response) *PayOrderQueryResult {
	return &PayOrderQueryResult{Result: result, Body: body, Http: http}
}

// PayOrderQuery
// 小程序支付 - 查询订单
// https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_2
func (c *Client) PayOrderQuery(ctx context.Context, transactionId, outTradeNo string, notMustParams ...gorequest.Params) (*PayOrderQueryResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "pay/orderquery")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", c.GetAppId())  // 小程序ID
	params.Set("mch_id", c.GetMchId()) // 商户号
	if transactionId != "" {
		params.Set("transaction_id", transactionId) // 微信订单号
	}
	if outTradeNo != "" {
		params.Set("out_trade_no", outTradeNo) // 商户订单号
	}
	params.Set("nonce_str", gorandom.Alphanumeric(32)) // 随机字符串

	// 签名
	params.Set("sign", c.getMd5Sign(params))

	// 	请求
	var response PayOrderQueryResponse
	request, err := c.request(ctx, "pay/orderquery", params, false, nil, &response)
	return newPayOrderQueryResult(response, request.ResponseBody, request), err
}
