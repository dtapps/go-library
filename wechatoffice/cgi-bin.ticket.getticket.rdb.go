package wechatoffice

import (
	"fmt"
	"time"
)

// GetJsapiTicket 获取api_ticket
func (app *App) GetJsapiTicket() string {
	if app.redis.Db == nil {
		return app.jsapiTicket
	}
	newCache := app.redis.NewSimpleStringCache(app.redis.NewStringOperation(), time.Second*7000)
	newCache.DBGetter = func() string {
		token := app.CgiBinTicketGetTicket("jsapi")
		return token.Result.Ticket
	}
	return newCache.GetCache(app.getJsapiTicketCacheKeyName())
}

func (app *App) getJsapiTicketCacheKeyName() string {
	return fmt.Sprintf("wechat_jsapi_ticket:%v", app.appId)
}
