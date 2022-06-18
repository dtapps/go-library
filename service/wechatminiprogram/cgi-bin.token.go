package wechatminiprogram

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
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
	Http   gorequest.Response  // 请求
	Err    error               // 错误
}

func NewCgiBinTokenResult(result CgiBinTokenResponse, body []byte, http gorequest.Response, err error) *CgiBinTokenResult {
	return &CgiBinTokenResult{Result: result, Body: body, Http: http, Err: err}
}

// CgiBinToken
// 接口调用凭证
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func (c *Client) CgiBinToken() *CgiBinTokenResult {
	// 请求
	request, err := c.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", c.getAppId(), c.getAppSecret()), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response CgiBinTokenResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewCgiBinTokenResult(response, request.ResponseBody, request, err)
}
