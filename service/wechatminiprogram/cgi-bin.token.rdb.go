package wechatminiprogram

import (
	"fmt"
	"github.com/dtapps/go-library/utils/goredis"
	"time"
)

func (app *App) GetAccessTokenRDb() string {
	cacheName := fmt.Sprintf("wechat_access_token:%v", app.AppId)
	redis := goredis.App{
		Rdb: app.RDb,
	}
	newCache := redis.NewSimpleStringCache(redis.NewStringOperation(), time.Second*7000)
	newCache.DBGetter = func() string {
		token := app.AuthGetAccessToken()
		return token.Result.AccessToken
	}
	return newCache.GetCache(cacheName)
}
