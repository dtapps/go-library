package wechatoffice

import (
	"errors"
)

// AuthGetJsapiTicketMonitor 监控api_ticket
func (app *App) AuthGetJsapiTicketMonitor(qdType string) error {
	result := app.GetCallBackIp()
	if len(result.Result.IpList) <= 0 {
		switch qdType {
		case qdTypeDb:
			app.GetJsapiTicketDb()
			return nil
		case qdTypeRdb:
			app.GetJsapiTicketRDb()
			return nil
		default:
			return errors.New("驱动类型不在范围内")
		}
	}
	return nil
}
