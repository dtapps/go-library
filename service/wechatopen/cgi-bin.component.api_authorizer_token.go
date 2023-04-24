package wechatopen

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
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
func (c *Client) CgiBinComponentApiAuthorizerToken(ctx context.Context, authorizerRefreshToken string) (*CgiBinComponentApiAuthorizerTokenResult, error) {
	// 检查
	err := c.checkComponentIsConfig()
	if err != nil {
		return nil, err
	}
	err = c.checkAuthorizerIsConfig()
	if err != nil {
		return nil, err
	}
	// 参数
	param := gorequest.NewParams()
	param["component_appid"] = c.GetComponentAppId()           // 第三方平台 appid
	param["authorizer_appid"] = c.GetAuthorizerAppid()         // 授权方 appid
	param["authorizer_refresh_token"] = authorizerRefreshToken // 授权码, 会在授权成功时返回给第三方平台
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/cgi-bin/component/api_authorizer_token?component_access_token=%v", c.GetComponentAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response CgiBinComponentApiAuthorizerTokenResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newCgiBinComponentApiAuthorizerTokenResult(response, request.ResponseBody, request, param["authorizer_appid"].(string)), nil
}
