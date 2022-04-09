package wechatoffice

import (
	"context"
	"errors"
	"time"
)

// GetJsapiTicketMonitor 监控api_ticket
func (app *App) GetJsapiTicketMonitor() (string, error) {
	if app.Redis.Db == nil {
		return "", errors.New("驱动没有初始化")
	}
	result := app.DebugCgiBinTicketCheck()
	if result.Result.Errcode == 0 {
		return app.JsapiTicket, nil
	}
	app.AccessToken = app.GetAccessToken()
	token := app.CgiBinTicketGetTicket("jsapi")
	app.Redis.Db.Set(context.Background(), app.getJsapiTicketCacheKeyName(), token.Result.Ticket, time.Second*7000)
	return token.Result.Ticket, nil
}
