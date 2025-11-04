package wechatpayopen

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

// 通过私钥对字符串以 SHA256WithRSA 算法生成签名信息
func (c *Client) signSHA256WithRSA(msg string, key *rsa.PrivateKey) (string, error) {
	if key == nil {
		return "", errors.New("private key is nil")
	}
	h := sha256.New()
	h.Write([]byte(msg))
	digest := h.Sum(nil)

	// 注意：这里应该用 SignPKCS1v15，但传的是 digest，且 hash 参数必须匹配
	// 实际上，更稳妥的是用 SignPSS 或直接用 crypto.Signer
	// 但微信要求 PKCS#1 v1.5，所以可以这样：

	sign, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, digest)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}
