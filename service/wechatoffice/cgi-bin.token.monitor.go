package wechatoffice

import (
	"errors"
)

var (
	QdTypeDb  = "DB"
	QdTypeRdb = "redis"
)

func (app *App) AuthGetAccessTokenMonitor(qdType string) error {
	result := app.GetCallBackIp()
	if len(result.Result.IpList) <= 0 {
		switch qdType {
		case QdTypeDb:
			if app.Db == nil {
				return errors.New("驱动没有初始化")
			}
			app.GetAccessTokenDb()
		case QdTypeRdb:
			if app.RDb == nil {
				return errors.New("驱动没有初始化")
			}
			app.GetAccessTokenRDb()
			return nil
		default:
			return errors.New("驱动类型不在范围内")
		}
	}
	return nil
}
