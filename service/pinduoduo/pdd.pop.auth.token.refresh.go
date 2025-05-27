package pinduoduo

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type PopAuthTokenRefresh struct {
	PopAuthTokenRefreshResponse PopAuthTokenCreatePopAuthTokenCreateResponse `json:"pop_auth_token_refresh_response"`
}

// PopAuthTokenRefresh 刷新Access Token
// https://open.pinduoduo.com/application/document/api?id=pdd.pop.auth.token.refresh
func (c *Client) PopAuthTokenRefresh(ctx context.Context, refreshToken string, notMustParams ...*gorequest.Params) (response PopAuthTokenRefresh, err error) {

	// 参数
	params := NewParamsWithType("pdd.pop.auth.token.refresh", notMustParams...)
	if refreshToken != "" {
		params.Set("refresh_token", refreshToken) // grantType==refresh_token 时需要
	}

	// 请求
	err = c.request(ctx, params, &response)
	return
}
