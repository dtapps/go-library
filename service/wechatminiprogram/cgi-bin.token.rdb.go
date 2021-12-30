package wechatminiprogram

import (
	"fmt"
	"gopkg.in/dtapps/go-library.v3/utils/goredis"
	"time"
)

func (app *App) GetAccessTokenRDb() string {
	cacheName := fmt.Sprintf("wechat_access_token:%v", app.AppId)
	redis := goredis.App{
		Rdb: app.RDb,
	}
	newCache := redis.NewSimpleStringCache(redis.NewStringOperation(), time.Minute*7000)
	newCache.DBGetter = func() string {
		token := app.AuthGetAccessToken()
		return token.Result.AccessToken
	}
	return newCache.GetCache(cacheName)
}
