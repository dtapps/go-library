package wechatpayapiv2

import (
	"context"
	"go.dtapp.net/library/utils/gorandom"
	"go.dtapp.net/library/utils/gorequest"
)

type PayCloseOrderResponse struct {
	ReturnCode string `json:"return_code" xml:"return_code"`                   // 返回状态码
	ReturnMsg  string `json:"return_msg,omitempty" xml:"return_msg,omitempty"` // 返回信息

	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`   // 业务结果
	ResultMsg  string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`     // 业务结果描述
	ErrCode    string `json:"err_code,omitempty" xml:"err_code,omitempty"`         // 错误代码
	ErrCodeDes string `json:"err_code_des,omitempty" xml:"err_code_des,omitempty"` // 错误代码描述

	Appid    string `json:"appid,omitempty" xml:"appid,omitempty"`         // 小程序ID
	MchId    string `json:"mch_id,omitempty" xml:"mch_id,omitempty"`       // 商户号
	NonceStr string `json:"nonce_str,omitempty" xml:"nonce_str,omitempty"` // 随机字符串
	Sign     string `json:"sign,omitempty" xml:"sign,omitempty"`           // 签名
}

type PayCloseOrderResult struct {
	Result PayCloseOrderResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newPayCloseOrderResult(result PayCloseOrderResponse, body []byte, http gorequest.Response) *PayCloseOrderResult {
	return &PayCloseOrderResult{Result: result, Body: body, Http: http}
}

// PayCloseOrder
// 小程序支付 - 关闭订单
// https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_3
func (c *Client) PayCloseOrder(ctx context.Context, outTradeNo string, notMustParams ...gorequest.Params) (*PayCloseOrderResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", c.GetAppId())                  // 小程序ID
	params.Set("mch_id", c.GetMchId())                 // 商户号
	params.Set("out_trade_no", outTradeNo)             // 商户订单号
	params.Set("nonce_str", gorandom.Alphanumeric(32)) // 随机字符串

	// 签名
	params.Set("sign", c.getMd5Sign(params))

	// 	请求
	var response PayCloseOrderResponse
	request, err := c.request(ctx, "pay/closeorder", params, false, nil, &response)
	return newPayCloseOrderResult(response, request.ResponseBody, request), err
}
