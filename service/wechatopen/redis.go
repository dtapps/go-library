package wechatopen

import (
	"context"
	"fmt"
	"time"
)

// 微信后台推送的ticke
func (app *App) getComponentVerifyTicketCacheKeyName() string {
	return fmt.Sprintf("wechat_open:component_verify_ticket:%v", app.componentAppId)
}

// SetComponentVerifyTicket 设置微信后台推送的ticke
func (app *App) SetComponentVerifyTicket(componentVerifyTicket string) string {
	if componentVerifyTicket == "" {
		return ""
	}
	app.redis.Db.Set(context.Background(), app.getComponentVerifyTicketCacheKeyName(), componentVerifyTicket, time.Hour*12)
	return app.GetComponentVerifyTicket()
}

// GetComponentVerifyTicket 获取微信后台推送的ticke
func (app *App) GetComponentVerifyTicket() string {
	if app.redis.Db == nil {
		return app.componentVerifyTicket
	}
	result, _ := app.redis.Db.Get(context.Background(), app.getComponentVerifyTicketCacheKeyName()).Result()
	return result
}

// 令牌
func (app *App) getComponentAccessTokenCacheKeyName() string {
	return fmt.Sprintf("wechat_open:component_access_token:%v", app.componentAppId)
}

// SetComponentAccessToken 设置令牌
func (app *App) SetComponentAccessToken(componentAccessToken string) string {
	if componentAccessToken == "" {
		return ""
	}
	app.redis.Db.Set(context.Background(), app.getComponentAccessTokenCacheKeyName(), componentAccessToken, time.Second*7200)
	return app.GetComponentAccessToken()
}

// GetComponentAccessToken 获取令牌
func (app *App) GetComponentAccessToken() string {
	if app.redis.Db == nil {
		return app.componentAccessToken
	}
	result, _ := app.redis.Db.Get(context.Background(), app.getComponentAccessTokenCacheKeyName()).Result()
	return result
}

// MonitorComponentAccessToken 监控令牌
func (app *App) MonitorComponentAccessToken() string {
	// 查询
	componentAccessToken := app.GetComponentAccessToken()
	// 判断
	result := app.CgiBinGetApiDomainIp(componentAccessToken)
	if len(result.Result.IpList) > 0 {
		return componentAccessToken
	}
	// 重新获取
	return app.SetComponentAccessToken(app.CgiBinComponentApiComponentToken().Result.ComponentAccessToken)
}

// 授权方令牌
func (app *App) getAuthorizerAccessTokenCacheKeyName() string {
	return fmt.Sprintf("wechat_open:authorizer_access_token:%v:%v", app.componentAppId, app.authorizerAppid)
}

// SetAuthorizerAccessToken 设置授权方令牌
func (app *App) SetAuthorizerAccessToken(authorizerAccessToken string) string {
	if authorizerAccessToken == "" {
		return ""
	}
	app.redis.Db.Set(context.Background(), app.getAuthorizerAccessTokenCacheKeyName(), authorizerAccessToken, time.Hour*2)
	return app.GetComponentAccessToken()
}

// GetAuthorizerAccessToken 获取授权方令牌
func (app *App) GetAuthorizerAccessToken() string {
	if app.redis.Db == nil {
		return app.authorizerAccessToken
	}
	result, _ := app.redis.Db.Get(context.Background(), app.getAuthorizerAccessTokenCacheKeyName()).Result()
	return result
}

// MonitorAuthorizerAccessToken 监控授权方令牌
func (app *App) MonitorAuthorizerAccessToken(authorizerRefreshToken string) string {
	// 查询
	authorizerAccessToken := app.GetAuthorizerAccessToken()
	// 判断
	if authorizerAccessToken != "" {
		return authorizerAccessToken
	}
	// 重新获取
	return app.SetAuthorizerAccessToken(app.CgiBinComponentApiAuthorizerToken(authorizerRefreshToken).Result.AuthorizerAccessToken)
}

// 预授权码
func (app *App) getPreAuthCodeCacheKeyName() string {
	return fmt.Sprintf("wechat_open:pre_auth_code:%v", app.componentAppId)
}

// SetPreAuthCode 设置预授权码
func (app *App) SetPreAuthCode(preAuthCode string) string {
	if preAuthCode == "" {
		return ""
	}
	app.redis.Db.Set(context.Background(), app.getPreAuthCodeCacheKeyName(), preAuthCode, time.Second*1700)
	return app.GetComponentAccessToken()
}

// GetPreAuthCode 获取预授权码
func (app *App) GetPreAuthCode() string {
	if app.redis.Db == nil {
		return app.authorizerAccessToken
	}
	result, _ := app.redis.Db.Get(context.Background(), app.getPreAuthCodeCacheKeyName()).Result()
	return result
}

// DelPreAuthCode 删除预授权码
func (app *App) DelPreAuthCode() error {
	return app.redis.Db.Del(context.Background(), app.getPreAuthCodeCacheKeyName()).Err()
}

// MonitorPreAuthCode 监控预授权码
func (app *App) MonitorPreAuthCode() string {
	// 查询
	preAuthCode := app.GetPreAuthCode()
	// 判断
	if preAuthCode != "" {
		return preAuthCode
	}
	// 重新获取
	return app.SetPreAuthCode(app.CgiBinComponentApiCreatePreAuthCoden().Result.PreAuthCode)
}
