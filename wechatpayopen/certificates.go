package wechatpayopen

import (
	"encoding/json"
	"go.dtapp.net/library/gorequest"
	"net/http"
	"time"
)

type CertificatesResponse struct {
	Data []struct {
		EffectiveTime      time.Time `json:"effective_time"` // 过期时间
		EncryptCertificate struct {
			Algorithm      string `json:"algorithm"`
			AssociatedData string `json:"associated_data"`
			Ciphertext     string `json:"ciphertext"`
			Nonce          string `json:"nonce"`
		} `json:"encrypt_certificate"` // 加密证书
		ExpireTime time.Time `json:"expire_time"` // 有效时间
		SerialNo   string    `json:"serial_no"`   // 序列号
	} `json:"data"`
}

type CertificatesResult struct {
	Result CertificatesResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
	Err    error                // 错误
}

func NewCertificatesResult(result CertificatesResponse, body []byte, http gorequest.Response, err error) *CertificatesResult {
	return &CertificatesResult{Result: result, Body: body, Http: http, Err: err}
}

// Certificates 获取平台证书列表
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/wechatpay5_1.shtml
func (app *App) Certificates() *CertificatesResult {
	// 请求
	request, err := app.request("https://api.mch.weixin.qq.com/v3/certificates", map[string]interface{}{}, http.MethodGet)
	if err != nil {
		return NewCertificatesResult(CertificatesResponse{}, request.ResponseBody, request, err)
	}
	// 定义
	var response CertificatesResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewCertificatesResult(response, request.ResponseBody, request, err)
}
