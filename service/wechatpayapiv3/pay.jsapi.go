package wechatpayapiv3

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gorandom"
	"time"
)

// GetJsApi 入参
type GetJsApi struct {
	Package string `json:"package"`
}

// GetJsApiResult 返回参数
type GetJsApiResult struct {
	AppId     string // 应用ID
	TimeStamp string // 时间戳
	NonceStr  string // 随机字符串
	Package   string // 订单详情扩展字符串
	SignType  string // 签名方式
	PaySign   string // 签名
}

// GetJsApi JSAPI调起支付API https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_4.shtml
func (c *Client) GetJsApi(ctx context.Context, param GetJsApi) (result GetJsApiResult, err error) {

	// sign params
	timeStamp := time.Now().Unix()
	nonce := gorandom.Alphanumeric(32)

	result.AppId = c.GetAppId()
	result.TimeStamp = fmt.Sprintf("%v", timeStamp)
	result.NonceStr = nonce
	result.Package = param.Package

	// 签名
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n", c.GetAppId(), fmt.Sprintf("%v", timeStamp), nonce, param.Package)

	signBytes, err := c.signPKCS1v15(message, []byte(c.GetMchSslKey()))
	if err != nil {
		return result, err
	}

	sign := c.base64EncodeStr(signBytes)
	result.PaySign = sign
	result.SignType = "RSA"
	return result, nil
}
