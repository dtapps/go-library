package wechatoffice

import (
	"fmt"
	"time"
)

func (app *App) GetAccessToken() string {
	if app.redis.Db == nil {
		return app.accessToken
	}
	newCache := app.redis.NewSimpleStringCache(app.redis.NewStringOperation(), time.Second*7000)
	newCache.DBGetter = func() string {
		token := app.CgiBinToken()
		return token.Result.AccessToken
	}
	return newCache.GetCache(app.getAccessTokenCacheKeyName())
}

func (app *App) getAccessTokenCacheKeyName() string {
	return fmt.Sprintf("wechat_access_token:%v", app.appId)
}
