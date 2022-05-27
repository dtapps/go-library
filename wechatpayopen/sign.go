package wechatpayopen

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
	"go.dtapp.net/library/gorandom"
	"net/url"
	"time"
)

// 对消息的散列值进行数字签名
func (app *App) signPKCS1v15(msg string, privateKey []byte) ([]byte, error) {

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

	sign, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, app.haSha256(msg))
	if err != nil {
		return nil, errors.New("sign error")
	}

	return sign, nil
}

// base编码
func (app *App) base64EncodeStr(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

// sha256加密
func (app *App) haSha256(str string) []byte {
	h := sha256.New()
	h.Write([]byte(str))
	return h.Sum(nil)
}

// 生成身份认证信息
// https://wechatpay-api.gitbook.io/wechatpay-api-v3/qian-ming-zhi-nan-1/qian-ming-sheng-cheng
func (app *App) authorization(method string, paramMap map[string]interface{}, rawUrl string) (token string, err error) {

	// 请求报文主体
	var signBody string
	if len(paramMap) != 0 {
		paramJsonBytes, err := json.Marshal(paramMap)
		if err != nil {
			return token, err
		}
		signBody = string(paramJsonBytes)
	}

	// URL
	urlPart, err := url.Parse(rawUrl)
	if err != nil {
		return token, err
	}
	canonicalUrl := urlPart.RequestURI()

	// 请求时间戳
	timestamp := time.Now().Unix()

	// 请求随机串
	nonce := gorandom.Alphanumeric(32)

	// 构造签名串
	message := fmt.Sprintf(SignatureMessageFormat, method, canonicalUrl, timestamp, nonce, signBody)

	sign, err := app.signSHA256WithRSA(message, app.getRsa([]byte(app.mchSslKey)))

	if err != nil {
		return token, err
	}

	authorization := fmt.Sprintf(
		HeaderAuthorizationFormat, getAuthorizationType(),
		app.spMchId, nonce, timestamp, app.mchSslSerialNo, sign,
	)

	return authorization, nil
}

// 报文解密
func (app *App) decryptGCM(aesKey, nonceV, ciphertextV, additionalDataV string) ([]byte, error) {
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

// 对消息的散列值进行数字签名
func (app *App) getRsa(privateKey []byte) *rsa.PrivateKey {

	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil
	}

	pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil
	}

	key, ok := pri.(*rsa.PrivateKey)
	if ok == false {
		return key
	}

	return key
}

// 通过私钥对字符串以 SHA256WithRSA 算法生成签名信息
func (app *App) signSHA256WithRSA(source string, privateKey *rsa.PrivateKey) (signature string, err error) {
	if privateKey == nil {
		return "", fmt.Errorf("private key should not be nil")
	}
	h := crypto.Hash.New(crypto.SHA256)
	_, err = h.Write([]byte(source))
	if err != nil {
		return "", nil
	}
	hashed := h.Sum(nil)
	signatureByte, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signatureByte), nil
}
