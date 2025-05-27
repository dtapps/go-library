package pinduoduo

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type PopAuthTokenCreatePopAuthTokenCreateResponse struct {
	AccessToken           string   `json:"access_token"`             // access_token
	ExpiresAt             int64    `json:"expires_at"`               // access_token过期时间点
	ExpiresIn             int64    `json:"expires_in"`               // access_token过期时间段，10（表示10秒后过期）
	OwnerId               string   `json:"owner_id"`                 // 商家店铺id
	OwnerName             string   `json:"owner_name"`               // 商家账号名称
	R1ExpiresAt           int64    `json:"r1_expires_at"`            // r1级别API或字段的访问过期时间点
	R1ExpiresIn           int64    `json:"r1_expires_in"`            // r1级别API或字段的访问过期时间； 10（表示10秒后过期）
	R2ExpiresAt           int64    `json:"r2_expires_at"`            // r2级别API或字段的访问过期时间点
	R2ExpiresIn           int64    `json:"r2_expires_in"`            // r2级别API或字段的访问过期时间；10（表示10秒后过期）
	RefreshToken          string   `json:"refresh_token"`            // refresh token，可用来刷新access_token
	RefreshTokenExpiresAt int64    `json:"refresh_token_expires_at"` // Refresh token过期时间点
	RefreshTokenExpiresIn int64    `json:"refresh_token_expires_in"` // refresh_token过期时间段，10表示10秒后过期
	Scope                 []string `json:"scope"`                    // 接口列表
	W1ExpiresAt           int64    `json:"w1_expires_at"`            // w1级别API或字段的访问过期时间点
	W1ExpiresIn           int64    `json:"w1_expires_in"`            // w1级别API或字段的访问过期时间； 10（表示10秒后过期）
	W2ExpiresAt           int64    `json:"w2_expires_at"`            // w2级别API或字段的访问过期时间点
	W2ExpiresIn           int64    `json:"w2_expires_in"`            // w2级别API或字段的访问过期时间；10（表示10秒后过期）
	RequestId             string   `json:"request_id"`
}

type PopAuthTokenCreate struct {
	PopAuthTokenCreateResponse PopAuthTokenCreatePopAuthTokenCreateResponse `json:"pop_auth_token_create_response"`
}

// PopAuthTokenCreate 获取Access Token
// https://open.pinduoduo.com/application/document/api?id=pdd.pop.auth.token.create
func (c *Client) PopAuthTokenCreate(ctx context.Context, code string, notMustParams ...*gorequest.Params) (response PopAuthTokenCreate, err error) {

	// 参数
	params := NewParamsWithType("pdd.pop.auth.token.create", notMustParams...)
	if code != "" {
		params.Set("code", code) // 授权code，grantType==authorization_code 时需要
	}

	// 请求
	err = c.request(ctx, params, &response)
	return
}
