package wechatopen

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type CgiBinComponentApiAuthorizerTokenResponse struct {
	AuthorizerAccessToken  string `json:"authorizer_access_token"`  // 授权方令牌
	ExpiresIn              int64  `json:"expires_in"`               // 有效期，单位：秒
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"` // 刷新令牌
}

type CgiBinComponentApiAuthorizerTokenResult struct {
	Result          CgiBinComponentApiAuthorizerTokenResponse // 结果
	Body            []byte                                    // 内容
	Http            gorequest.Response                        // 请求
	authorizerAppid string                                    // 授权方 appid
}

func newCgiBinComponentApiAuthorizerTokenResult(result CgiBinComponentApiAuthorizerTokenResponse, body []byte, http gorequest.Response, authorizerAppid string) *CgiBinComponentApiAuthorizerTokenResult {
	return &CgiBinComponentApiAuthorizerTokenResult{Result: result, Body: body, Http: http, authorizerAppid: authorizerAppid}
}

// CgiBinComponentApiAuthorizerToken 获取/刷新接口调用令牌
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_authorizer_token.html
func (c *Client) CgiBinComponentApiAuthorizerToken(ctx context.Context, componentAccessToken, authorizerAppid, authorizerRefreshToken string, notMustParams ...*gorequest.Params) (*CgiBinComponentApiAuthorizerTokenResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("component_appid", c.GetComponentAppId())           // 第三方平台appid
	params.Set("authorizer_appid", authorizerAppid)                // 授权方appid
	params.Set("authorizer_refresh_token", authorizerRefreshToken) // 授权码会在授权成功时返回给第三方平台

	// 请求
	var response CgiBinComponentApiAuthorizerTokenResponse
	request, err := c.request(ctx, fmt.Sprintf("cgi-bin/component/api_authorizer_token?component_access_token=%s", componentAccessToken), params, http.MethodPost, &response)
	return newCgiBinComponentApiAuthorizerTokenResult(response, request.ResponseBody, request, params.Get("authorizer_appid").(string)), err
}
