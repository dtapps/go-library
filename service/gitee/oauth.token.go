package gitee

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type OauthTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type OauthTokenResult struct {
	Result OauthTokenResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
	Err    error              // 错误
}

func newOauthTokenResult(result OauthTokenResponse, body []byte, http gorequest.Response, err error) *OauthTokenResult {
	return &OauthTokenResult{Result: result, Body: body, Http: http, Err: err}
}

// OauthToken OAuth2 获取 AccessToken 认证步骤
// https://gitee.com/api/v5/oauth_doc#/list-item-2
func (c *Client) OauthToken(code string) *OauthTokenResult {
	// 参数
	params := gorequest.NewParamsWith()
	// 请求
	request, err := c.request(apiUrl+fmt.Sprintf("/oauth/token?grant_type=authorization_code&code=%s&client_id=%s&redirect_uri=%s&client_secret=%s", code, c.config.ClientID, c.config.RedirectUri, c.config.ClientSecret), params, http.MethodPost)
	// 定义
	var response OauthTokenResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newOauthTokenResult(response, request.ResponseBody, request, err)
}
