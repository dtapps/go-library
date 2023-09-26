package alipayopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type AlipaySystemOauthTokenResponse struct {
	AlipaySystemOauthTokenResponse struct {
		AlipayUserId string `json:"alipay_user_id,omitempty"`
		UserId       string `json:"user_id"`       // 支付宝用户的唯一标识。以2088开头的16位数字。
		AccessToken  string `json:"access_token"`  // 访问令牌。通过该令牌调用需要授权类接口
		ExpiresIn    string `json:"expires_in"`    // 	访问令牌的有效时间，单位是秒。
		RefreshToken string `json:"refresh_token"` // 刷新令牌。通过该令牌可以刷新access_token
		ReExpiresIn  string `json:"re_expires_in"` // 刷新令牌的有效时间，单位是秒。
		AuthStart    string `json:"auth_start"`    // 授权token开始时间，作为有效期计算的起点
	} `json:"alipay_system_oauth_token_response"`
}

type AlipaySystemOauthTokenResult struct {
	Result AlipaySystemOauthTokenResponse // 结果
	Body   []byte                         // 内容
	Http   gorequest.Response             // 请求
}

func newAlipaySystemOauthTokenResult(result AlipaySystemOauthTokenResponse, body []byte, http gorequest.Response) *AlipaySystemOauthTokenResult {
	return &AlipaySystemOauthTokenResult{Result: result, Body: body, Http: http}
}

// AlipaySystemOauthToken 换取授权访问令牌
// https://opendocs.alipay.com/open/02xtla
func (c *Client) AlipaySystemOauthToken(ctx context.Context, grantType, code, refreshToken string, notMustParams ...gorequest.Params) (*AlipaySystemOauthTokenResult, ApiError, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("grant_type", grantType)
	if code != "" {
		params.Set("code", code)
	}
	if refreshToken != "" {
		params.Set("refresh_token", refreshToken)
	}
	// 请求
	request, err := c.request(ctx, c.newParamsWithType("alipay.system.oauth.token", params))
	if err != nil {
		return newAlipaySystemOauthTokenResult(AlipaySystemOauthTokenResponse{}, request.ResponseBody, request), ApiError{}, err
	}
	// 定义
	var response AlipaySystemOauthTokenResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = gojson.Unmarshal(request.ResponseBody, &apiError)
	return newAlipaySystemOauthTokenResult(response, request.ResponseBody, request), apiError, err
}
