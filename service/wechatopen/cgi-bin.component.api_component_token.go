package wechatopen

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type CgiBinComponentApiComponentTokenResponse struct {
	ComponentAccessToken string `json:"component_access_token"` // 第三方平台 access_token
	ExpiresIn            int64  `json:"expires_in"`             // 有效期，单位：秒
}

type CgiBinComponentApiComponentTokenResult struct {
	Result CgiBinComponentApiComponentTokenResponse // 结果
	Body   []byte                                   // 内容
	Http   gorequest.Response                       // 请求
	Err    error                                    // 错误
}

func newCgiBinComponentApiComponentTokenResult(result CgiBinComponentApiComponentTokenResponse, body []byte, http gorequest.Response, err error) *CgiBinComponentApiComponentTokenResult {
	return &CgiBinComponentApiComponentTokenResult{Result: result, Body: body, Http: http, Err: err}
}

// CgiBinComponentApiComponentToken 令牌
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_access_token.html
func (c *Client) CgiBinComponentApiComponentToken() *CgiBinComponentApiComponentTokenResult {
	// 参数
	param := gorequest.NewParams()
	param["component_appid"] = c.config.ComponentAppId              // 第三方平台 appid
	param["component_appsecret"] = c.config.ComponentAppSecret      // 第三方平台 appsecret
	param["component_verify_ticket"] = c.GetComponentVerifyTicket() // 微信后台推送的 ticket
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(apiUrl+"/cgi-bin/component/api_component_token", params, http.MethodPost)
	// 定义
	var response CgiBinComponentApiComponentTokenResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newCgiBinComponentApiComponentTokenResult(response, request.ResponseBody, request, err)
}
