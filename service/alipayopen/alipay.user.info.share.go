package alipayopen

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
)

type AlipayUserInfoShareResponse struct {
	AlipayUserInfoShareResponse struct {
		AlipayUserId string `json:"alipay_user_id,omitempty"`
		UserId       string `json:"user_id"`       // 支付宝用户的唯一标识。以2088开头的16位数字。
		AccessToken  string `json:"access_token"`  // 访问令牌。通过该令牌调用需要授权类接口
		ExpiresIn    string `json:"expires_in"`    // 	访问令牌的有效时间，单位是秒。
		RefreshToken string `json:"refresh_token"` // 刷新令牌。通过该令牌可以刷新access_token
		ReExpiresIn  string `json:"re_expires_in"` // 刷新令牌的有效时间，单位是秒。
		AuthStart    string `json:"auth_start"`    // 授权token开始时间，作为有效期计算的起点
	} `json:"alipay_system_oauth_token_response"`
}

type AlipayUserInfoShareResult struct {
	Result   AlipayUserInfoShareResponse // 结果
	Body     []byte                      // 内容
	Http     gorequest.Response          // 请求
	ApiError ApiError                    // 接口错误
}

func newAlipayUserInfoShareResult(result AlipayUserInfoShareResponse, body []byte, http gorequest.Response, apiError ApiError) *AlipayUserInfoShareResult {
	return &AlipayUserInfoShareResult{Result: result, Body: body, Http: http, ApiError: apiError}
}

// AlipayUserInfoShare 换取授权访问令牌
// https://opendocs.alipay.com/open/02xtlb
func (c *Client) AlipayUserInfoShare(ctx context.Context, authToken string) (*AlipayUserInfoShareResult, error) {
	// 参数
	params := gorequest.NewParams()
	params.Set("auth_token", authToken)
	// 请求
	request, err := c.request(ctx, c.newParamsWithType("alipay.user.info.share", params))
	if err != nil {
		return nil, err
	}
	// 定义
	var response AlipayUserInfoShareResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = json.Unmarshal(request.ResponseBody, &apiError)
	return newAlipayUserInfoShareResult(response, request.ResponseBody, request, apiError), err
}
