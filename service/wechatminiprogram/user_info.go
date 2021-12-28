package wechatminiprogram

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
)

// UserInfo 请求参数
type UserInfo struct {
	SessionKey    string `json:"session_key"`
	EncryptedData string `json:"encrypted_data"`
	Iv            string `json:"iv"`
}

// UserInfoResult 返回参数
type UserInfoResult struct {
	OpenID    string `json:"openId"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
	UnionId   string `json:"unionId"`
	Watermark struct {
		AppID     string `json:"appid"`
		Timestamp int64  `json:"timestamp"`
	} `json:"watermark"`
}

// UserInfo 解密用户信息
func (app *App) UserInfo(param UserInfo) (result UserInfoResult, err error) {
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
