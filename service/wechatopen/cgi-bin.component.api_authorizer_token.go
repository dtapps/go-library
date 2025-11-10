package wechatopen

import (
	"context"
	"fmt"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type CgiBinComponentApiAuthorizerTokenResponse struct {
	AuthorizerAccessToken  string `json:"authorizer_access_token"`  // 授权方令牌
	ExpiresIn              int64  `json:"expires_in"`               // 有效期，单位：秒
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"` // 刷新令牌
}

// CgiBinComponentApiAuthorizerToken 获取/刷新接口调用令牌
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_authorizer_token.html
func (c *Client) CgiBinComponentApiAuthorizerToken(ctx context.Context, componentAccessToken, authorizerAppid, authorizerRefreshToken string, notMustParams ...*gorequest.Params) (response CgiBinComponentApiAuthorizerTokenResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("component_appid", c.GetComponentAppId())           // 第三方平台appid
	params.Set("authorizer_appid", authorizerAppid)                // 授权方appid
	params.Set("authorizer_refresh_token", authorizerRefreshToken) // 授权码会在授权成功时返回给第三方平台

	// 请求
	err = c.request(ctx, fmt.Sprintf("cgi-bin/component/api_authorizer_token?component_access_token=%s", componentAccessToken), params, http.MethodPost, &response)
	return
}
