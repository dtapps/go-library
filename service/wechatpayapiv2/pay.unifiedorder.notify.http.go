package wechatpayapiv2

import (
	"context"
	"encoding/xml"
	"net/http"
)

// PayUnifiedOrderNotifyHttpRequest 小程序支付 - 统一下单 - 回调通知 - 请求参数
type PayUnifiedOrderNotifyHttpRequest struct {
	ReturnCode         string `json:"return_code" xml:"return_code"`                                       // 返回状态码
	ReturnMsg          string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`                     // 返回信息
	Appid              string `json:"appid,omitempty" xml:"appid,omitempty"`                               // 小程序ID
	MchId              string `json:"mch_id,omitempty" xml:"mch_id,omitempty"`                             // 商户号
	DeviceInfo         string `json:"device_info,omitempty" xml:"device_info,omitempty"`                   // 设备号
	NonceStr           string `json:"nonce_str,omitempty" xml:"nonce_str,omitempty"`                       // 随机字符串
	Sign               string `json:"sign,omitempty" xml:"sign,omitempty"`                                 // 签名
	SignType           string `json:"sign_type,omitempty" xml:"sign_type,omitempty"`                       // 签名类型
	ResultCode         string `json:"result_code,omitempty" xml:"result_code,omitempty"`                   // 业务结果
	ErrCode            string `json:"err_code,omitempty" xml:"err_code,omitempty"`                         // 错误代码
	ErrCodeDes         string `json:"err_code_des,omitempty" xml:"err_code_des,omitempty"`                 // 错误代码描述
	Openid             string `json:"openid,omitempty" xml:"openid,omitempty"`                             // 用户标识
	IsSubscribe        string `json:"is_subscribe,omitempty" xml:"is_subscribe,omitempty"`                 // 是否关注公众账号
	TradeType          string `json:"trade_type,omitempty" xml:"trade_type,omitempty"`                     // 交易类型
	BankType           string `json:"bank_type,omitempty" xml:"bank_type,omitempty"`                       // 付款银行
	TotalFee           int    `json:"total_fee,omitempty" xml:"total_fee,omitempty"`                       // 订单金额
	SettlementTotalFee int    `json:"settlement_total_fee,omitempty" xml:"settlement_total_fee,omitempty"` // 应结订单金额
	FeeType            string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`                         // 货币种类
	CashFee            int    `json:"cash_fee,omitempty" xml:"cash_fee,omitempty"`                         // 现金支付金额
	CashFeeType        string `json:"cash_fee_type,omitempty" xml:"cash_fee_type,omitempty"`               // 现金支付货币类型
	CouponFee          string `json:"coupon_fee,omitempty" xml:"coupon_fee,omitempty"`                     // 总代金券金额
	CouponCount        int    `json:"coupon_count,omitempty" xml:"coupon_count,omitempty"`                 // 代金券使用数量
	CouponType         string `json:"coupon_type,omitempty" xml:"coupon_type,omitempty"`                   // 代金券类型
	CouponId           string `json:"coupon_id,omitempty" xml:"coupon_id,omitempty"`                       // 代金券ID
	TransactionId      string `json:"transaction_id,omitempty" xml:"transaction_id,omitempty"`             // 微信支付订单号
	OutTradeNo         string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`                 // 商户订单号
	Attach             string `json:"attach,omitempty" xml:"attach,omitempty"`                             // 商家数据包
	TimeEnd            string `json:"time_end,omitempty" xml:"time_end,omitempty"`                         // 支付完成时间
}

// PayUnifiedOrderNotifyHttp 小程序支付 - 统一下单 - 回调通知
// https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_7&index=8
func (c *Client) PayUnifiedOrderNotifyHttp(ctx context.Context, w http.ResponseWriter, r *http.Request) (validateXml PayUnifiedOrderNotifyHttpRequest, err error) {
	err = xml.NewDecoder(r.Body).Decode(&validateXml)
	return validateXml, err
}
