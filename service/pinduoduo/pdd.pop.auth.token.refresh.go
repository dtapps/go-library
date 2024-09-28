package pinduoduo

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type PddDdkPopAuthTokenRefreshResponse struct {
	PopAuthTokenRefreshResponse struct {
		AccessToken           string   `json:"access_token"`             // access_token
		ExpiresAt             int64    `json:"expires_at"`               // access_token过期时间点
		ExpiresIn             int      `json:"expires_in"`               // access_token过期时间段，10（表示10秒后过期）
		OwnerId               string   `json:"owner_id"`                 // 商家店铺id
		OwnerName             string   `json:"owner_name"`               // 商家账号名称
		R1ExpiresAt           int64    `json:"r1_expires_at"`            // r1级别API或字段的访问过期时间点
		R1ExpiresIn           int      `json:"r1_expires_in"`            // r1级别API或字段的访问过期时间； 10（表示10秒后过期）
		R2ExpiresAt           int64    `json:"r2_expires_at"`            // r2级别API或字段的访问过期时间点
		R2ExpiresIn           int      `json:"r2_expires_in"`            // r2级别API或字段的访问过期时间；10（表示10秒后过期）
		RefreshToken          string   `json:"refresh_token"`            // refresh token，可用来刷新access_token
		RefreshTokenExpiresAt int64    `json:"refresh_token_expires_at"` // Refresh token过期时间点
		RefreshTokenExpiresIn int      `json:"refresh_token_expires_in"` // refresh_token过期时间段，10表示10秒后过期
		Scope                 []string `json:"scope"`                    // 接口列表
		W1ExpiresAt           int64    `json:"w1_expires_at"`            // w1级别API或字段的访问过期时间点
		W1ExpiresIn           int      `json:"w1_expires_in"`            // w1级别API或字段的访问过期时间； 10（表示10秒后过期）
		W2ExpiresAt           int64    `json:"w2_expires_at"`            // w2级别API或字段的访问过期时间点
		W2ExpiresIn           int      `json:"w2_expires_in"`            // w2级别API或字段的访问过期时间；10（表示10秒后过期）
		RequestId             string   `json:"request_id"`
	} `json:"pop_auth_token_refresh_response"`
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
func (c *Client) PopAuthTokenRefresh(ctx context.Context, refresh_token string, notMustParams ...gorequest.Params) (*PddDdkPopAuthTokenRefreshResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "pdd.pop.auth.token.refresh")
	defer span.End()

	// 参数
	param := gorequest.NewParamsWith(notMustParams...)
	if refresh_token != "" {
		param.Set("refresh_token", refresh_token) // grantType==refresh_token 时需要
	}
	params := NewParamsWithType("pdd.pop.auth.token.refresh", notMustParams...)

	// 请求
	var response PddDdkPopAuthTokenRefreshResponse
	request, err := c.request(ctx, span, params, &response)
	return newPddDdkPopAuthTokenRefreshResult(response, request.ResponseBody, request), err
}
