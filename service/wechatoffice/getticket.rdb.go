package wechatoffice

import (
	"fmt"
	"github.com/dtapps/go-library/utils/goredis"
	"time"
)

// GetJsapiTicketRDb 获取api_ticket
func (app *App) GetJsapiTicketRDb() string {
	cacheName := fmt.Sprintf("wechat_jsapi_ticket:%v", app.AppId)
	redis := goredis.App{
		Rdb: app.RDb,
	}
	newCache := redis.NewSimpleStringCache(redis.NewStringOperation(), time.Second*7000)
	newCache.DBGetter = func() string {
		token := app.GetTicket(app.GetAccessTokenRDb(), "jsapi")
		return token.Result.Ticket
	}
	return newCache.GetCache(cacheName)
}
