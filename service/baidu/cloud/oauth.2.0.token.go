package cloud

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type Oauth2TokenResponse struct {
	RefreshToken  string `json:"refresh_token"`  // 刷新令牌
	ExpiresIn     int64  `json:"expires_in"`     // 到期时间
	SessionKey    string `json:"session_key"`    // 会话密钥
	AccessToken   string `json:"access_token"`   // 访问令牌
	Scope         string `json:"scope"`          // 范围
	SessionSecret string `json:"session_secret"` // 会话机密
}

type Oauth2TokenResult struct {
	Result Oauth2TokenResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newOauth2TokenResult(result Oauth2TokenResponse, body []byte, http gorequest.Response) *Oauth2TokenResult {
	return &Oauth2TokenResult{Result: result, Body: body, Http: http}
}

// Oauth2Token Oauth2Token
func (c *Client) Oauth2Token(ctx context.Context, notMustParams ...gorequest.Params) (*Oauth2TokenResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("grant_type", "client_credentials")
	params.Set("client_id", c.apiKey)
	params.Set("client_secret", c.secretKey)
	// 请求
	request, err := c.request(ctx, "oauth/2.0/token", params, http.MethodPost, "FORM")
	if err != nil {
		return newOauth2TokenResult(Oauth2TokenResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response Oauth2TokenResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newOauth2TokenResult(response, request.ResponseBody, request), err
}
