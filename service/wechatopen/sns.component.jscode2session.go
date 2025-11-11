package wechatopen

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go.dtapp.net/library/utils/gorequest"
)

type ThirdpartyCode2SessionResponse struct {
	APIResponse        // 错误
	Openid      string `json:"openid"`      // 用户唯一标识的 openid
	SessionKey  string `json:"session_key"` // 会话密钥
	Unionid     string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见 UnionID 机制说明。
}

// ThirdpartyCode2Session 小程序登录
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/login/thirdpartyCode2Session.html
func (c *Client) ThirdpartyCode2Session(ctx context.Context, jsCode string, notMustParams ...*gorequest.Params) (response ThirdpartyCode2SessionResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", c.GetAuthorizerAppid())          // 小程序的 AppID
	params.Set("grant_type", "authorization_code")       // 填 authorization_code
	params.Set("component_appid", c.GetComponentAppId()) // 第三方平台 appid
	params.Set("js_code", jsCode)                        // wx.login 获取的 code

	// 请求
	err = c.request(ctx, "sns/component/jscode2session?component_access_token="+c.GetComponentAccessToken(), params, http.MethodGet, &response)
	return
}

type UserInfoParams struct {
	SessionKey    string `json:"session_key"`
	EncryptedData string `json:"encrypted_data"`
	Iv            string `json:"iv"`
}

type UserInfoResponse struct {
	OpenId    string `json:"openId"`
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

func newUserInfoResult(result UserInfoResponse, err error) *UserInfoResult {
	return &UserInfoResult{Result: result}
}

// UserInfo 解密用户信息
func (r *ThirdpartyCode2SessionResponse) UserInfo(param UserInfoParams) *UserInfoResult {
	var response UserInfoResponse
	aesKey, err := base64.StdEncoding.DecodeString(r.SessionKey)
	if err != nil {
		return newUserInfoResult(response, err)
	}
	cipherText, err := base64.StdEncoding.DecodeString(param.EncryptedData)
	if err != nil {
		return newUserInfoResult(response, err)
	}
	ivBytes, err := base64.StdEncoding.DecodeString(param.Iv)
	if err != nil {
		return newUserInfoResult(response, err)
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return newUserInfoResult(response, err)
	}
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(cipherText, cipherText)
	cipherText, err = r.pkcs7Unpaid(cipherText, block.BlockSize())
	if err != nil {
		return newUserInfoResult(response, err)
	}
	err = json.Unmarshal(cipherText, &response)
	if err != nil {
		return newUserInfoResult(response, err)
	}
	return newUserInfoResult(response, err)
}

// DecryptionUserInfo 解密用户信息
func DecryptionUserInfo(param UserInfoParams) *UserInfoResult {
	var response UserInfoResponse
	aesKey, err := base64.StdEncoding.DecodeString(param.SessionKey)
	if err != nil {
		return newUserInfoResult(response, err)
	}
	cipherText, err := base64.StdEncoding.DecodeString(param.EncryptedData)
	if err != nil {
		return newUserInfoResult(response, err)
	}
	ivBytes, err := base64.StdEncoding.DecodeString(param.Iv)
	if err != nil {
		return newUserInfoResult(response, err)
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return newUserInfoResult(response, err)
	}
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(cipherText, cipherText)
	cipherText, err = pkcs7Unpaid(cipherText, block.BlockSize())
	if err != nil {
		return newUserInfoResult(response, err)
	}
	err = json.Unmarshal(cipherText, &response)
	if err != nil {
		return newUserInfoResult(response, err)
	}
	return newUserInfoResult(response, err)
}

func (u *UserInfoResponse) UserInfoAvatarUrlReal() string {
	return UserInfoAvatarUrlReal(u.AvatarUrl)
}

func UserInfoAvatarUrlReal(avatarUrl string) string {
	return strings.Replace(avatarUrl, "/132", "/0", -1)
}

func (r *ThirdpartyCode2SessionResponse) pkcs7Unpaid(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, fmt.Errorf("invalid block size")
	}
	if len(data)%blockSize != 0 || len(data) == 0 {
		return nil, fmt.Errorf("invalid PKCS7 data")
	}
	c := data[len(data)-1]
	n := int(c)
	if n == 0 || n > len(data) {
		return nil, fmt.Errorf("invalid padding on input")
	}
	for i := 0; i < n; i++ {
		if data[len(data)-n+i] != c {
			return nil, fmt.Errorf("invalid padding on input")
		}
	}
	return data[:len(data)-n], nil
}

func pkcs7Unpaid(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, fmt.Errorf("invalid block size")
	}
	if len(data)%blockSize != 0 || len(data) == 0 {
		return nil, fmt.Errorf("invalid PKCS7 data")
	}
	c := data[len(data)-1]
	n := int(c)
	if n == 0 || n > len(data) {
		return nil, fmt.Errorf("invalid padding on input")
	}
	for i := 0; i < n; i++ {
		if data[len(data)-n+i] != c {
			return nil, fmt.Errorf("invalid padding on input")
		}
	}
	return data[:len(data)-n], nil
}
