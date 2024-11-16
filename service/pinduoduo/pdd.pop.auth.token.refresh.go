package pinduoduo

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type PddDdkPopAuthTokenRefreshResponse struct {
	PopAuthTokenRefreshResponse PddDdkPopAuthTokenPopAuthTokenResponse `json:"pop_auth_token_refresh_response"`
}

type PddDdkPopAuthTokenRefreshResult struct {
	Result PddDdkPopAuthTokenRefreshResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
}

func newPddDdkPopAuthTokenRefreshResult(result PddDdkPopAuthTokenRefreshResponse, body []byte, http gorequest.Response) *PddDdkPopAuthTokenRefreshResult {
	return &PddDdkPopAuthTokenRefreshResult{Result: result, Body: body, Http: http}
}

// PopAuthTokenRefresh 刷新Access Token
// https://open.pinduoduo.com/application/document/api?id=pdd.pop.auth.token.refresh
func (c *Client) PopAuthTokenRefresh(ctx context.Context, refreshToken string, notMustParams ...*gorequest.Params) (*PddDdkPopAuthTokenRefreshResult, error) {

	// 参数
	params := NewParamsWithType("pdd.pop.auth.token.refresh", notMustParams...)
	if refreshToken != "" {
		params.Set("refresh_token", refreshToken) // grantType==refresh_token 时需要
	}

	// 请求
	var response PddDdkPopAuthTokenRefreshResponse
	request, err := c.request(ctx, params, &response)
	return newPddDdkPopAuthTokenRefreshResult(response, request.ResponseBody, request), err
}
