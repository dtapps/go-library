package wechatpayapiv2

import (
	"context"
	"encoding/xml"
	"github.com/dtapps/go-library/utils/gorandom"
	"github.com/dtapps/go-library/utils/gorequest"
)

type TransfersResponse struct {
	ReturnCode     string `json:"return_code" xml:"return_code"`                       // 返回状态码
	ReturnMsg      string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`     // 返回信息
	MchAppid       string `json:"mch_appid" xml:"mch_appid"`                           // 商户appid
	Mchid          string `json:"mchid" xml:"mchid"`                                   // 商户号
	DeviceInfo     string `json:"device_info,omitempty" xml:"device_info,omitempty"`   // 设备号
	NonceStr       string `json:"nonce_str" xml:"nonce_str"`                           // 随机字符串
	ResultCode     string `json:"result_code" xml:"result_code"`                       // 业务结果
	ErrCode        string `json:"err_code,omitempty" xml:"err_code,omitempty"`         // 错误代码
	ErrCodeDes     string `json:"err_code_des,omitempty" xml:"err_code_des,omitempty"` // 错误代码描述
	PartnerTradeNo string `json:"partner_trade_no" xml:"partner_trade_no"`             // 商户订单号
	PaymentNo      string `json:"payment_no" xml:"payment_no"`                         // 微信付款单号
	PaymentTime    string `json:"payment_time" xml:"payment_time"`                     // 付款成功时间
}

type TransfersResult struct {
	Result TransfersResponse  // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
	Err    error              // 错误
}

func newTransfersResult(result TransfersResponse, body []byte, http gorequest.Response, err error) *TransfersResult {
	return &TransfersResult{Result: result, Body: body, Http: http, Err: err}
}

// Transfers
// 付款到零钱 - 付款
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2
func (c *Client) Transfers(ctx context.Context, partnerTradeNo, openid string, amount int64, desc string) *TransfersResult {
	cert, err := c.P12ToPem()
	// 参数
	params := NewParams()
	params.Set("mch_appid", c.GetAppId())
	params.Set("mchid", c.GetMchId())
	params.Set("nonce_str", gorandom.Alphanumeric(32))
	params.Set("partner_trade_no", partnerTradeNo)
	params.Set("openid", openid)
	params.Set("check_name", "NO_CHECK")
	params.Set("amount", amount)
	params.Set("desc", desc)
	// 签名
	params.Set("sign", c.getMd5Sign(params))
	// 	请求
	request, err := c.request(ctx, apiUrl+"/mmpaymkttransfers/promotion/transfers", params, cert)
	// 定义
	var response TransfersResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return newTransfersResult(response, request.ResponseBody, request, err)
}
