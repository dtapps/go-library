package wechatoffice

import (
	"errors"
)

// AuthGetJsapiTicketMonitor 监控api_ticket
func (app *App) AuthGetJsapiTicketMonitor(qdType string) error {
	result := app.GetCallBackIp()
	if len(result.Result.IpList) <= 0 {
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
	return nil
}
