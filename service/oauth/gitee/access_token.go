package gitee

import (
	"encoding/json"
	"fmt"
)

// GetAccessToken 请求参数
type GetAccessToken struct {
	Code string
}

// GetAccessTokenResult 返回参数
type GetAccessTokenResult struct {
	AccessToken string `json:"access_token"`
}

// GetAccessToken OAuth2 获取 AccessToken 认证步骤 https://gitee.com/api/v5/oauth_doc#/list-item-2
func (app *App) GetAccessToken(param GetAccessToken) (result GetAccessTokenResult, err error) {
	url := fmt.Sprintf("https://gitee.com/oauth/token?grant_type=authorization_code&code=%s&client_id=%s&redirect_uri=%s&client_secret=%s", param.Code, app.ClientID, app.RedirectUri, app.ClientSecret)

	// api params
	params := map[string]interface{}{}

	// common params

	// request
	body, err := app.request(url, params, "POST")
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
