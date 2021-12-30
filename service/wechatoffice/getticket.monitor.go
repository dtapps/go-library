package wechatoffice

import "errors"

// AuthGetJsapiTicketMonitor 监控api_ticket
func (app *App) AuthGetJsapiTicketMonitor(qdType string) error {
	switch qdType {
	case QdTypeDb:
		app.GetJsapiTicketDb()
		return nil
	case QdTypeRdb:
		app.GetJsapiTicketRDb()
		return nil
	default:
		return errors.New("驱动类型不在范围内")
	}
}
