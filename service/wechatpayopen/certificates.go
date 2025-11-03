package wechatpayopen

import (
	"context"
	"net/http"
	"time"

	"go.dtapp.net/library/utils/gorequest"
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

// Certificates 获取平台证书列表
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/wechatpay5_1.shtml
func (c *Client) Certificates(ctx context.Context, notMustParams ...*gorequest.Params) (response CertificatesResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "v3/certificates", params, http.MethodGet, &response, nil)
	return
}
