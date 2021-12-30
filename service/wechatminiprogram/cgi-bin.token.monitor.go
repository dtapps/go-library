package wechatminiprogram

import (
	"errors"
)

var (
	QdTypeDb  = "DB"
	QdTypeRdb = "redis"
)

func (app *App) AuthGetAccessTokenMonitor(qdType string) (string, error) {
	switch qdType {
	case QdTypeDb:
		if app.Db == nil {
			return "", errors.New("驱动没有初始化")
		}
		app.AccessToken = app.GetAccessTokenDb()
		result := app.GetCallBackIp()
		if len(result.Result.IpList) <= 0 {
			return app.GetAccessTokenDb(), nil
		}
		return app.AccessToken, nil
	case QdTypeRdb:
		if app.RDb == nil {
			return "", errors.New("驱动没有初始化")
		}
		app.AccessToken = app.GetAccessTokenRDb()
		result := app.GetCallBackIp()
		if len(result.Result.IpList) <= 0 {
			return app.GetAccessTokenRDb(), nil
		}
		return app.AccessToken, nil
	default:
		return "", errors.New("驱动类型不在范围内")
	}
	return "", nil
}
