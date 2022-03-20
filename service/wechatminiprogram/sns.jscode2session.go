package wechatminiprogram

import (
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gohttp"
	"net/http"
)

type AuthCode2SessionResponse struct {
	OpenId     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	Unionid    string `json:"unionid"`     // 用户在开放平台的唯一标识符，若当前小程序已绑定到微信开放平台帐号下会返回
	Errcode    string `json:"errcode"`     // 错误码
	Errmsg     string `json:"errmsg"`      // 错误信息
}

type AuthCode2SessionResult struct {
	Result AuthCode2SessionResponse // 结果
	Body   []byte                   // 内容
	Http   gohttp.Response          // 请求
	Err    error                    // 错误
}

func NewAuthCode2SessionResult(result AuthCode2SessionResponse, body []byte, http gohttp.Response, err error) *AuthCode2SessionResult {
	return &AuthCode2SessionResult{Result: result, Body: body, Http: http, Err: err}
}

func (app *App) AuthCode2Session(jsCode string) *AuthCode2SessionResult {
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", app.AppId, app.AppSecret, jsCode), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response AuthCode2SessionResponse
	err = json.Unmarshal(request.Body, &response)
	return NewAuthCode2SessionResult(response, request.Body, request, err)
}
