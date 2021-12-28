package wechatoffice

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
)

// UserPhone 请求参数
type UserPhone struct {
	SessionKey    string `json:"session_key"`
	EncryptedData string `json:"encrypted_data"`
	Iv            string `json:"iv"`
}

// UserPhoneResult 返回参数
type UserPhoneResult struct {
	PhoneNumber     string `json:"phoneNumber"`     // 用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string `json:"purePhoneNumber"` // 没有区号的手机号
	CountryCode     string `json:"countryCode"`     // 区号
	Watermark       struct {
		AppID     string `json:"appid"`
		Timestamp int64  `json:"timestamp"`
	} `json:"watermark"`
}

// UserPhone 解密手机号信息
func (app *App) UserPhone(param UserPhone) (result UserPhoneResult, err error) {
	aesKey, err := base64.StdEncoding.DecodeString(param.SessionKey)
	if err != nil {
		return result, err
	}
	cipherText, err := base64.StdEncoding.DecodeString(param.EncryptedData)
	if err != nil {
		return result, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(param.Iv)
	if err != nil {
		return result, err
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return result, err
	}
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(cipherText, cipherText)
	cipherText, err = app.pkcs7Unpaid(cipherText, block.BlockSize())
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(cipherText, &result)
	if err != nil {
		return result, err
	}
	if result.Watermark.AppID != app.AppId {
		return result, errors.New("app id not match")
	}
	return result, nil
}
