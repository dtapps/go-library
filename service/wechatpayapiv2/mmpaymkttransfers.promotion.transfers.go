package wechatpayapiv2

import (
	"context"

	"go.dtapp.net/library/utils/gorandom"
	"go.dtapp.net/library/utils/gorequest"
)

type TransfersResponse struct {
	ReturnCode string `json:"return_code" xml:"return_code"`                   // 返回状态码
	ReturnMsg  string `json:"return_msg,omitempty" xml:"return_msg,omitempty"` // 返回信息

	ResultCode string `json:"result_code" xml:"result_code"`                       // 业务结果
	ErrCode    string `json:"err_code,omitempty" xml:"err_code,omitempty"`         // 错误代码
	ErrCodeDes string `json:"err_code_des,omitempty" xml:"err_code_des,omitempty"` // 错误代码描述

	MchAppid       string `json:"mch_appid" xml:"mch_appid"`                         // 商户appid
	Mchid          string `json:"mchid" xml:"mchid"`                                 // 商户号
	DeviceInfo     string `json:"device_info,omitempty" xml:"device_info,omitempty"` // 设备号
	NonceStr       string `json:"nonce_str" xml:"nonce_str"`                         // 随机字符串
	PartnerTradeNo string `json:"partner_trade_no" xml:"partner_trade_no"`           // 商户订单号
	PaymentNo      string `json:"payment_no" xml:"payment_no"`                       // 微信付款单号
	PaymentTime    string `json:"payment_time" xml:"payment_time"`                   // 付款成功时间
}

// Transfers
// 付款到零钱 - 付款
// 需要证书
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2
func (c *Client) Transfers(ctx context.Context, notMustParams ...*gorequest.Params) (response TransfersResponse, err error) {

	// 证书
	cert, err := c.P12ToPem()
	if err != nil {
		return
	}

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("mch_appid", c.GetAppId())
	params.Set("mchid", c.GetMchId())
	params.Set("nonce_str", gorandom.Alphanumeric(32))

	// 签名
	params.Set("sign", c.getMd5Sign(params))

	// 	请求
	err = c.request(ctx, "mmpaymkttransfers/promotion/transfers", params, cert, &response)
	return
}
