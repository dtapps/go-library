package wechatpayapiv2

import (
	"context"
	"github.com/gin-gonic/gin"
)

// PayUnifiedOrderNotifyGinRequest 小程序支付 - 统一下单 - 回调通知 - 请求参数
type PayUnifiedOrderNotifyGinRequest struct {
	ReturnCode         string `form:"return_code" json:"return_code" xml:"return_code" uri:"return_code" binding:"required"`                                      // 返回状态码
	ReturnMsg          string `form:"return_msg" json:"return_msg" xml:"return_msg" uri:"return_msg" binding:"omitempty"`                                         // 返回信息
	Appid              string `form:"appid" json:"appid" xml:"appid" uri:"appid" binding:"omitempty"`                                                             // 小程序ID
	MchId              string `form:"mch_id" json:"mch_id" xml:"mch_id" uri:"mch_id" binding:"omitempty"`                                                         // 商户号
	DeviceInfo         string `form:"device_info" json:"device_info" xml:"device_info" uri:"device_info" binding:"omitempty"`                                     // 设备号
	NonceStr           string `form:"nonce_str" json:"nonce_str" xml:"nonce_str" uri:"nonce_str" binding:"omitempty"`                                             // 随机字符串
	Sign               string `form:"sign" json:"sign" xml:"sign" uri:"sign" binding:"omitempty"`                                                                 // 签名
	SignType           string `form:"sign_type" json:"sign_type" xml:"sign_type" uri:"sign_type" binding:"omitempty"`                                             // 签名类型
	ResultCode         string `form:"result_code" json:"result_code" xml:"result_code" uri:"result_code" binding:"omitempty"`                                     // 业务结果
	ErrCode            string `form:"err_code" json:"err_code" xml:"err_code" uri:"err_code" binding:"omitempty"`                                                 // 错误代码
	ErrCodeDes         string `form:"err_code_des" json:"err_code_des" xml:"err_code_des" uri:"err_code_des" binding:"omitempty"`                                 // 错误代码描述
	Openid             string `form:"openid" json:"openid" xml:"openid" uri:"openid" binding:"omitempty"`                                                         // 用户标识
	IsSubscribe        string `form:"is_subscribe" json:"is_subscribe" xml:"is_subscribe" uri:"is_subscribe" binding:"omitempty"`                                 // 是否关注公众账号
	TradeType          string `form:"trade_type" json:"trade_type" xml:"trade_type" uri:"trade_type" binding:"omitempty"`                                         // 交易类型
	BankType           string `form:"bank_type" json:"bank_type" xml:"bank_type" uri:"bank_type" binding:"omitempty"`                                             // 付款银行
	TotalFee           int    `form:"total_fee" json:"total_fee" xml:"total_fee" uri:"total_fee" binding:"omitempty"`                                             // 订单金额
	SettlementTotalFee int    `form:"settlement_total_fee" json:"settlement_total_fee" xml:"settlement_total_fee" uri:"settlement_total_fee" binding:"omitempty"` // 应结订单金额
	FeeType            string `form:"fee_type" json:"fee_type" xml:"fee_type" uri:"fee_type" binding:"omitempty"`                                                 // 货币种类
	CashFee            int    `form:"cash_fee" json:"cash_fee" xml:"cash_fee" uri:"cash_fee" binding:"omitempty"`                                                 // 现金支付金额
	CashFeeType        string `form:"cash_fee_type" json:"cash_fee_type" xml:"cash_fee_type" uri:"cash_fee_type" binding:"omitempty"`                             // 现金支付货币类型
	CouponFee          string `form:"coupon_fee" json:"coupon_fee" xml:"coupon_fee" uri:"coupon_fee" binding:"omitempty"`                                         // 总代金券金额
	CouponCount        int    `form:"coupon_count" json:"coupon_count" xml:"coupon_count" uri:"coupon_count" binding:"omitempty"`                                 // 代金券使用数量
	CouponType         string `form:"coupon_type" json:"coupon_type" xml:"coupon_type" uri:"coupon_type" binding:"omitempty"`                                     // 代金券类型
	CouponId           string `form:"coupon_id" json:"coupon_id" xml:"coupon_id" uri:"coupon_id" binding:"omitempty"`                                             // 代金券ID
	TransactionId      string `form:"transaction_id" json:"transaction_id" xml:"transaction_id" uri:"transaction_id" binding:"omitempty"`                         // 微信支付订单号
	OutTradeNo         string `form:"out_trade_no" json:"out_trade_no" xml:"out_trade_no" uri:"out_trade_no" binding:"omitempty"`                                 // 商户订单号
	Attach             string `form:"attach" json:"attach" xml:"attach" uri:"attach" binding:"omitempty"`                                                         // 商家数据包
	TimeEnd            string `form:"time_end" json:"time_end" xml:"time_end" uri:"time_end" binding:"omitempty"`                                                 // 支付完成时间
}

// PayUnifiedOrderNotifyGin 小程序支付 - 统一下单 - 回调通知
// https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_7&index=8
func (c *Client) PayUnifiedOrderNotifyGin(ctx context.Context, ginCtx *gin.Context) (validateJson PayUnifiedOrderNotifyGinRequest, err error) {

	// 解析
	err = ginCtx.ShouldBind(&validateJson)

	return validateJson, err
}
