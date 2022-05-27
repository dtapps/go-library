package wechatopen

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
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
	Http            gorequest.Response                        // 请求
	Err             error                                     // 错误
	authorizerAppid string                                    // 授权方 appid
}

func NewCgiBinComponentApiAuthorizerTokenResult(result CgiBinComponentApiAuthorizerTokenResponse, body []byte, http gorequest.Response, err error, authorizerAppid string) *CgiBinComponentApiAuthorizerTokenResult {
	return &CgiBinComponentApiAuthorizerTokenResult{Result: result, Body: body, Http: http, Err: err, authorizerAppid: authorizerAppid}
}

// CgiBinComponentApiAuthorizerToken 获取/刷新接口调用令牌
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_authorizer_token.html
func (app *App) CgiBinComponentApiAuthorizerToken(authorizerRefreshToken string) *CgiBinComponentApiAuthorizerTokenResult {
	// 参数
	param := NewParams()
	param["component_appid"] = app.componentAppId              // 第三方平台 appid
	param["authorizer_appid"] = app.authorizerAppid            // 授权方 appid
	param["authorizer_refresh_token"] = authorizerRefreshToken // 授权码, 会在授权成功时返回给第三方平台
	params := app.NewParamsWith(param)
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=%v", app.GetComponentAccessToken()), params, http.MethodPost)
	// 定义
	var response CgiBinComponentApiAuthorizerTokenResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewCgiBinComponentApiAuthorizerTokenResult(response, request.ResponseBody, request, err, app.authorizerAppid)
}
