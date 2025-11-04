package wechatpayapiv3

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"strconv"
	"time"

	"go.dtapp.net/library/utils/gorandom"
)

// 对消息的散列值进行数字签名
func (c *Client) signPKCS1v15(msg string, privateKey []byte) ([]byte, error) {

	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key decode error")
	}
	pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("parse private key error")
	}
	key, ok := pri.(*rsa.PrivateKey)
	if ok == false {
		return nil, errors.New("private key format error")
	}
	sign, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, c.haSha256(msg))
	if err != nil {
		return nil, errors.New("sign error")
	}
	return sign, nil
}

// base编码
func (c *Client) base64EncodeStr(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

// sha256加密
func (c *Client) haSha256(str string) []byte {
	h := sha256.New()
	h.Write([]byte(str))
	return h.Sum(nil)
}

// 报文解密
func (c *Client) decryptGCM(aesKey, nonceV, ciphertextV, additionalDataV string) ([]byte, error) {
	key := []byte(aesKey)
	nonce := []byte(nonceV)
	additionalData := []byte(additionalDataV)
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextV)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, additionalData)
	if err != nil {
		return nil, err
	}
	return plaintext, err
}

type SignParams struct {
	Method              string          // HTTP请求方法
	Body                map[string]any  // 请求参数
	Url                 string          // 请求路径
	PrivateKey          *rsa.PrivateKey // 商户私钥
	MchId               string          // 微信支付的商户id
	CertificateSerialNo string          // 证书序列号
}
type SignResult struct {
	Authorization string
	BodyBytes     []byte // 用于 SetBody
}

// 签名
// https://pay.weixin.qq.com/doc/v3/partner/4012365870
func Sign(param *SignParams) (resule *SignResult, err error) {

	// 构造请求体
	var bodyBytes []byte
	var bodyStr string
	if len(param.Body) > 0 {
		var err error
		bodyBytes, err = json.Marshal(param.Body)
		if err != nil {
			return nil, err
		}
		bodyStr = string(bodyBytes)
	} else {
		bodyStr = ""
	}

	// 解析 URL 路径（不含 query）
	urlPart, err := url.Parse(param.Url)
	if err != nil {
		return nil, err
	}
	urlPath := urlPart.Path

	// 时间戳和 nonce
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := gorandom.Alphanumeric(32)

	// 生成签名
	sig, err := generateSignature(param.Method, urlPath, timestamp, nonce, bodyStr, param.PrivateKey)
	if err != nil {
		slog.Error("[authorization] generateSignature failed", slog.Any("err", err))
		return nil, err
	}

	// 构造完整的 Authorization Header
	authz := fmt.Sprintf(
		`WECHATPAY2-SHA256-RSA2048 mchid="%s",nonce_str="%s",signature="%s",timestamp="%s",serial_no="%s"`,
		param.MchId,
		nonce,
		sig,
		timestamp,
		param.CertificateSerialNo,
	)

	return &SignResult{
		Authorization: authz,
		BodyBytes:     bodyBytes,
	}, nil
}

func generateSignature(method, urlPath, timestamp, nonce, body string, privateKey *rsa.PrivateKey) (string, error) {
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", method, urlPath, timestamp, nonce, body)

	h := sha256.New()
	h.Write([]byte(message))
	digest := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, digest)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}
