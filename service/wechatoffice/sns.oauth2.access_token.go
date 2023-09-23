package wechatoffice

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type SnsOauth2AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	ExpiresIn    int    `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"` // 用户刷新access_token
	Openid       string `json:"openid"`        // 用户唯一标识
	Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
}

type SnsOauth2AccessTokenResult struct {
	Result SnsOauth2AccessTokenResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
}

func newSnsOauth2AccessTokenResult(result SnsOauth2AccessTokenResponse, body []byte, http gorequest.Response) *SnsOauth2AccessTokenResult {
	return &SnsOauth2AccessTokenResult{Result: result, Body: body, Http: http}
}

// SnsOauth2AccessToken 通过code换取网页授权access_token
// https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html#0
func (c *Client) SnsOauth2AccessToken(ctx context.Context, code string, notMustParams ...*gorequest.Params) (*SnsOauth2AccessTokenResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", c.GetAppId(), c.GetAppSecret(), code), params, http.MethodGet)
	if err != nil {
		return newSnsOauth2AccessTokenResult(SnsOauth2AccessTokenResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response SnsOauth2AccessTokenResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newSnsOauth2AccessTokenResult(response, request.ResponseBody, request), err
}
