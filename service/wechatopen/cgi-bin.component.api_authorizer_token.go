package wechatopen

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CgiBinComponentApiAuthorizerTokenResponse struct {
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int64  `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}

type CgiBinComponentApiAuthorizerTokenResult struct {
	Result          CgiBinComponentApiAuthorizerTokenResponse // 结果
	Body            []byte                                    // 内容
	Err             error                                     // 错误
	authorizerAppid string                                    // 授权方 appid
}

func NewCgiBinComponentApiAuthorizerTokenResult(result CgiBinComponentApiAuthorizerTokenResponse, body []byte, err error, authorizerAppid string) *CgiBinComponentApiAuthorizerTokenResult {
	return &CgiBinComponentApiAuthorizerTokenResult{Result: result, Body: body, Err: err, authorizerAppid: authorizerAppid}
}

// CgiBinComponentApiAuthorizerToken 获取/刷新接口调用令牌
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_authorizer_token.html
func (app *App) CgiBinComponentApiAuthorizerToken() *CgiBinComponentApiAuthorizerTokenResult {
	app.componentAccessToken = app.GetComponentAccessToken()
	// 参数
	param := NewParams()
	param["component_appid"] = app.ComponentAppId                       // 第三方平台 appid
	param["authorizer_appid"] = app.AuthorizerAppid                     // 授权方 appid
	param["authorizer_refresh_token"] = app.GetAuthorizerRefreshToken() // 授权码, 会在授权成功时返回给第三方平台
	params := app.NewParamsWith(param)
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=%v", app.componentAccessToken), params, http.MethodPost)
	// 定义
	var response CgiBinComponentApiAuthorizerTokenResponse
	err = json.Unmarshal(body, &response)
	return NewCgiBinComponentApiAuthorizerTokenResult(response, body, err, app.AuthorizerAppid)
}
