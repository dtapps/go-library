package wechatminiprogram

import (
	"dtapps/dta/library/utils/gohttp"
	"encoding/json"
	"fmt"
	"net/http"
)

type CgiBinTokenResponse struct {
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值
	Errcode     int    `json:"errcode"`      // 错误码
	Errmsg      string `json:"errmsg"`       // 错误信息
}

type CgiBinTokenResult struct {
	Result CgiBinTokenResponse // 结果
	Body   []byte              // 内容
	Http   gohttp.Response     // 请求
	Err    error               // 错误
}

func NewCgiBinTokenResult(result CgiBinTokenResponse, body []byte, http gohttp.Response, err error) *CgiBinTokenResult {
	return &CgiBinTokenResult{Result: result, Body: body, Http: http, Err: err}
}

// CgiBinToken
// 接口调用凭证
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func (app *App) CgiBinToken() *CgiBinTokenResult {
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", app.AppId, app.AppSecret), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response CgiBinTokenResponse
	err = json.Unmarshal(request.Body, &response)
	return NewCgiBinTokenResult(response, request.Body, request, err)
}
