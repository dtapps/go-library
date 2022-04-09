package wechatoffice

import (
	"fmt"
	"time"
)

// GetJsapiTicket 获取api_ticket
func (app *App) GetJsapiTicket() string {
	if app.Redis.Db == nil {
		return app.JsapiTicket
	}
	newCache := app.Redis.NewSimpleStringCache(app.Redis.NewStringOperation(), time.Second*7000)
	newCache.DBGetter = func() string {
		token := app.CgiBinTicketGetTicket("jsapi")
		return token.Result.Ticket
	}
	return newCache.GetCache(app.getJsapiTicketCacheKeyName())
}

func (app *App) getJsapiTicketCacheKeyName() string {
	return fmt.Sprintf("wechat_jsapi_ticket:%v", app.AppId)
}
