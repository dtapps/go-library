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
	"gitee.com/dtapps/go-library/utils/random"
	"io/ioutil"
	"net/url"
	"os"
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
func (app *App) authorization(method string, paramMap map[string]interface{}, rawUrl string) (token string, err error) {
	var body string
	if len(paramMap) != 0 {
		paramJsonBytes, err := json.Marshal(paramMap)
		if err != nil {
			return token, err
		}
		body = string(paramJsonBytes)
	}
	urlPart, err := url.Parse(rawUrl)
	if err != nil {
		return token, err
	}
	canonicalUrl := urlPart.RequestURI()
	timestamp := time.Now().Unix()
	nonce := random.Alphanumeric(32)
	message := fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n", method, canonicalUrl, timestamp, nonce, body)
	open, err := os.Open(app.MchPrivateKey) // 商户私有证书路径或者从数据库读取
	if err != nil {
		return token, err
	}
	defer open.Close()
	privateKey, err := ioutil.ReadAll(open)

	if err != nil {
		return token, err
	}

	signBytes, err := app.signPKCS1v15(message, privateKey)

	if err != nil {
		return token, err
	}

	sign := app.base64EncodeStr(signBytes)

	token = fmt.Sprintf("mchid=\"%s\",nonce_str=\"%s\",timestamp=\"%d\",serial_no=\"%s\",signature=\"%s\"",
		app.MchId, nonce, timestamp, app.PrivateSerialNo, sign)
	return token, nil
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
