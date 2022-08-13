package wechatminiprogram

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
)

type UserPhone struct {
	SessionKey    string `json:"session_key"`
	EncryptedData string `json:"encrypted_data"`
	Iv            string `json:"iv"`
}

type UserPhoneResponse struct {
	PhoneNumber     string `json:"phoneNumber"`     // 用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string `json:"purePhoneNumber"` // 没有区号的手机号
	CountryCode     string `json:"countryCode"`     // 区号
	Watermark       struct {
		AppID     string `json:"appid"`
		Timestamp int64  `json:"timestamp"`
	} `json:"watermark"`
}

type UserPhoneResult struct {
	Result UserPhoneResponse // 结果
	Err    error             // 错误
}

func newUserPhoneResult(result UserPhoneResponse, err error) *UserPhoneResult {
	return &UserPhoneResult{Result: result, Err: err}
}

// UserPhone 解密手机号信息
func (c *Client) UserPhone(ctx context.Context, param UserPhone) *UserPhoneResult {
	var response UserPhoneResponse
	aesKey, err := base64.StdEncoding.DecodeString(param.SessionKey)
	if err != nil {
		return newUserPhoneResult(response, err)
	}
	cipherText, err := base64.StdEncoding.DecodeString(param.EncryptedData)
	if err != nil {
		return newUserPhoneResult(response, err)
	}
	ivBytes, err := base64.StdEncoding.DecodeString(param.Iv)
	if err != nil {
		return newUserPhoneResult(response, err)
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return newUserPhoneResult(response, err)
	}
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(cipherText, cipherText)
	cipherText, err = c.pkcs7Unpaid(cipherText, block.BlockSize())
	if err != nil {
		return newUserPhoneResult(response, err)
	}
	err = json.Unmarshal(cipherText, &response)
	if err != nil {
		return newUserPhoneResult(response, err)
	}
	if response.Watermark.AppID != c.getAppId() {
		return newUserPhoneResult(response, errors.New("c id not match"))
	}
	return newUserPhoneResult(response, err)
}
