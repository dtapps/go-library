package wechatoffice

import "errors"

// AuthGetJsapiTicketMonitor 监控api_ticket
func (app *App) AuthGetJsapiTicketMonitor(qdType string) (string, error) {
	switch qdType {
	case QdTypeDb:
		if app.Db == nil {
			return "", errors.New("驱动没有初始化")
		}
		app.AccessToken = app.GetAccessTokenDb()
		return app.GetJsapiTicketDb(), nil
	case QdTypeRdb:
		if app.RDb == nil {
			return "", errors.New("驱动没有初始化")
		}
		app.AccessToken = app.GetAccessTokenRDb()
		return app.GetJsapiTicketRDb(), nil
	default:
		return "", errors.New("驱动类型不在范围内")
	}
}
