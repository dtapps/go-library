package wechatopen

import (
	"github.com/dtapps/go-library/utils/gotime"
	"gorm.io/gorm"
	"time"
)

// GetAuthorizerAccessTokenMonitor 获取获取/刷新接口调用令牌和监控
func (app *App) GetAuthorizerAccessTokenMonitor() string {
	// 查询
	authorizerAccessToken := app.GetAuthorizerAccessToken()
	if authorizerAccessToken != "" {
		return authorizerAccessToken
	}
	// 重新获取
	return app.SetAuthorizerAccessToken(app.CgiBinComponentApiAuthorizerToken()).AuthorizerAccessToken
}

// SetAuthorizerAccessToken 设置获取/刷新接口调用令牌和自动获取
func (app *App) SetAuthorizerAccessToken(info *CgiBinComponentApiAuthorizerTokenResult) CgiBinComponentApiAuthorizerTokenResponse {
	if app.Db == nil || info.Result.AuthorizerAccessToken == "" || info.Result.AuthorizerRefreshToken == "" || info.authorizerAppid == "" {
		return CgiBinComponentApiAuthorizerTokenResponse{}
	}
	app.Db.Create(&AuthorizerAccessToken{
		ComponentAppId:         app.ComponentAppId,
		AuthorizerAppId:        info.authorizerAppid,
		AuthorizerAccessToken:  info.Result.AuthorizerAccessToken,
		AuthorizerRefreshToken: info.Result.AuthorizerRefreshToken,
		ExpiresIn:              info.Result.ExpiresIn,
		ExpireTime:             gotime.Current().AfterHour(2).Time,
	})
	return info.Result
}

type AuthorizerAccessToken struct {
	gorm.Model
	ComponentAppId         string    `json:"component_app_id"`         // 第三方平台 appid
	AuthorizerAppId        string    `json:"authorizer_app_id"`        // 授权方 appid
	AuthorizerAccessToken  string    `json:"authorizer_access_token"`  // 接口调用令牌（在授权的公众号/小程序具备 API 权限时，才有此返回值）
	AuthorizerRefreshToken string    `json:"authorizer_refresh_token"` // 刷新令牌（在授权的公众号具备API权限时，才有此返回值），刷新令牌主要用于第三方平台获取和刷新已授权用户的 authorizer_access_token。一旦丢失，只能让用户重新授权，才能再次拿到新的刷新令牌。用户重新授权后，之前的刷新令牌会失效
	ExpiresIn              int64     `json:"expires_in"`               // 有效期，单位：秒
	ExpireTime             time.Time `json:"expire_time"`              // 过期时间
}

func (m *AuthorizerAccessToken) TableName() string {
	return "authorizer_access_token"
}
