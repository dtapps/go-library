package wechatopen

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
	"strings"
)

type SnsComponentJsCode2sessionResponse struct {
	Openid     string `json:"openid"`      // 用户唯一标识的 openid
	SessionKey string `json:"session_key"` // 会话密钥
	Unionid    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见 UnionID 机制说明。
}

type SnsComponentJsCode2sessionResult struct {
	Result SnsComponentJsCode2sessionResponse // 结果
	Body   []byte                             // 内容
	Http   gorequest.Response                 // 请求
	Err    error                              // 错误
}

func newSnsComponentJsCode2sessionResult(result SnsComponentJsCode2sessionResponse, body []byte, http gorequest.Response, err error) *SnsComponentJsCode2sessionResult {
	return &SnsComponentJsCode2sessionResult{Result: result, Body: body, Http: http, Err: err}
}

// SnsComponentJsCode2session 小程序登录
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/others/WeChat_login.html
func (c *Client) SnsComponentJsCode2session(jsCode string) *SnsComponentJsCode2sessionResult {
	// 参数
	params := gorequest.NewParams()
	params["appid"] = c.config.AuthorizerAppid                     // 小程序的 appId
	params["js_code"] = jsCode                                     // wx.login 获取的 code
	params["grant_type"] = "authorization_code"                    // 填 authorization_code
	params["component_appid"] = c.config.ComponentAppId            // 第三方平台 appid
	params["component_access_token"] = c.GetComponentAccessToken() // 第三方平台的component_access_token
	// 请求
	request, err := c.request(apiUrl+"/sns/component/jscode2session", params, http.MethodGet)
	// 定义
	var response SnsComponentJsCode2sessionResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newSnsComponentJsCode2sessionResult(response, request.ResponseBody, request, err)
}

type UserInfo struct {
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
	return &UserInfoResult{Result: result, Err: err}
}

// UserInfo 解密用户信息
func (r *SnsComponentJsCode2sessionResult) UserInfo(param UserInfo) *UserInfoResult {
	var response UserInfoResponse
	aesKey, err := base64.StdEncoding.DecodeString(r.Result.SessionKey)
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

func (u *UserInfoResponse) UserInfoAvatarUrlReal() string {
	return UserInfoAvatarUrlReal(u.AvatarUrl)
}

func UserInfoAvatarUrlReal(avatarUrl string) string {
	return strings.Replace(avatarUrl, "/132", "/0", -1)
}

func (r *SnsComponentJsCode2sessionResult) pkcs7Unpaid(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, errors.New("invalid block size")
	}
	if len(data)%blockSize != 0 || len(data) == 0 {
		return nil, errors.New("invalid PKCS7 data")
	}
	c := data[len(data)-1]
	n := int(c)
	if n == 0 || n > len(data) {
		return nil, errors.New("invalid padding on input")
	}
	for i := 0; i < n; i++ {
		if data[len(data)-n+i] != c {
			return nil, errors.New("invalid padding on input")
		}
	}
	return data[:len(data)-n], nil
}
