package wechatpayapiv2

import (
	"context"
	"encoding/xml"
	"github.com/dtapps/go-library/utils/gorandom"
	"github.com/dtapps/go-library/utils/gorequest"
)

type PayUnifiedOrderResponse struct {
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

	TradeType string `json:"trade_type,omitempty" xml:"trade_type,omitempty"` // 交易类型
	PrepayId  string `json:"prepay_id,omitempty" xml:"prepay_id,omitempty"`   // 预支付交易会话标识
	CodeUrl   string `json:"code_url,omitempty" xml:"code_url,omitempty"`     // 二维码链接
}

type PayUnifiedOrderResult struct {
	Result PayUnifiedOrderResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
}

func newPayUnifiedOrderResult(result PayUnifiedOrderResponse, body []byte, http gorequest.Response) *PayUnifiedOrderResult {
	return &PayUnifiedOrderResult{Result: result, Body: body, Http: http}
}

// PayUnifiedOrder
// 小程序支付 - 统一下单
// https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_1
func (c *Client) PayUnifiedOrder(ctx context.Context, notMustParams ...gorequest.Params) (*PayUnifiedOrderResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", c.GetAppId())                  // 小程序ID
	params.Set("mch_id", c.GetMchId())                 // 商户号
	params.Set("nonce_str", gorandom.Alphanumeric(32)) // 随机字符串
	// 签名
	params.Set("sign", c.getMd5Sign(params))
	// 	请求
	request, err := c.request(ctx, apiUrl+"/pay/unifiedorder", params, false, nil)
	if err != nil {
		return newPayUnifiedOrderResult(PayUnifiedOrderResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PayUnifiedOrderResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return newPayUnifiedOrderResult(response, request.ResponseBody, request), err
}
