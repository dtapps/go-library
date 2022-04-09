package wechatminiprogram

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

type UserInfo struct {
	SessionKey    string `json:"session_key"`
	EncryptedData string `json:"encrypted_data"`
	Iv            string `json:"iv"`
}

type UserInfoResponse struct {
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

type UserInfoResult struct {
	Result UserInfoResponse // 结果
	Err    error            // 错误
}

func NewUserInfoResult(result UserInfoResponse, err error) *UserInfoResult {
	return &UserInfoResult{Result: result, Err: err}
}

// UserInfo 解密用户信息
func (app *App) UserInfo(param UserInfo) *UserInfoResult {
	var response UserInfoResponse
	aesKey, err := base64.StdEncoding.DecodeString(param.SessionKey)
	if err != nil {
		return NewUserInfoResult(response, err)
	}
	cipherText, err := base64.StdEncoding.DecodeString(param.EncryptedData)
	if err != nil {
		return NewUserInfoResult(response, err)
	}
	ivBytes, err := base64.StdEncoding.DecodeString(param.Iv)
	if err != nil {
		return NewUserInfoResult(response, err)
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return NewUserInfoResult(response, err)
	}
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(cipherText, cipherText)
	cipherText, err = app.pkcs7Unpaid(cipherText, block.BlockSize())
	if err != nil {
		return NewUserInfoResult(response, err)
	}
	err = json.Unmarshal(cipherText, &response)
	if err != nil {
		return NewUserInfoResult(response, err)
	}
	if response.Watermark.AppID != app.AppId {
		return NewUserInfoResult(response, errors.New("app id not match"))
	}
	return NewUserInfoResult(response, err)
}

func (u *UserInfoResponse) UserInfoAvatarUrlReal() string {
	return strings.Replace(u.AvatarUrl, "/132", "/0", -1)
}
