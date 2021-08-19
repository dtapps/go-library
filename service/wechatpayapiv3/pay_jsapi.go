package wechatpayapiv3

import (
	"fmt"
	"gitee.com/dtapps/go-library/utils/random"
	"io/ioutil"
	"os"
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
func (app *App) GetJsApi(param GetJsApi) (params GetJsApiResult, err error) {

	// sign params
	timeStamp := time.Now().Unix()
	nonce := random.Alphanumeric(32)

	params.AppId = app.AppId
	params.TimeStamp = fmt.Sprintf("%v", timeStamp)
	params.NonceStr = nonce
	params.Package = param.Package

	// 签名
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n", app.AppId, fmt.Sprintf("%v", timeStamp), nonce, param.Package)
	open, err := os.Open(app.MchPrivateKey)
	if err != nil {
		return params, err
	}
	defer open.Close()
	privateKey, err := ioutil.ReadAll(open)
	if err != nil {
		return params, err
	}
	signBytes, err := app.signPKCS1v15(message, privateKey)
	if err != nil {
		return params, err
	}

	sign := app.base64EncodeStr(signBytes)
	params.PaySign = sign
	params.SignType = "RSA"
	return params, nil
}
