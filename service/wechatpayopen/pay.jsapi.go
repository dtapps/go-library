package wechatpayopen

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gorandom"
	"time"
)

// GetJsApi 入参
type GetJsApi struct {
	Package string `json:"package"`
}

// GetJsApiResult 返回参数
type GetJsApiResult struct {
	AppId     string `json:"app_id"`     // 应用ID
	TimeStamp string `json:"time_stamp"` // 时间戳
	NonceStr  string `json:"nonce_str"`  // 随机字符串
	Package   string `json:"package"`    // 订单详情扩展字符串
	SignType  string `json:"sign_type"`  // 签名方式
	PaySign   string `json:"pay_sign"`   // 签名
}

// GetJsApi JSAPI调起支付API https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_4.shtml
func (c *Client) GetJsApi(ctx context.Context, param GetJsApi) (result GetJsApiResult, err error) {

	// sign params
	timeStamp := time.Now().Unix()
	nonce := gorandom.Alphanumeric(32)

	result.AppId = c.GetSubAppid()
	result.TimeStamp = fmt.Sprintf("%v", timeStamp) // 时间戳
	result.NonceStr = nonce                         // 随机字符串
	result.Package = param.Package                  // 订单详情扩展字符串

	// 签名
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n", c.GetSubAppid(), fmt.Sprintf("%v", timeStamp), nonce, param.Package)

	signBytes, err := c.signPKCS1v15(message, []byte(c.GetMchSslKey()))
	if err != nil {
		return result, err
	}

	sign := c.base64EncodeStr(signBytes)
	result.PaySign = sign   // 签名
	result.SignType = "RSA" // 签名方式
	return result, nil
}
