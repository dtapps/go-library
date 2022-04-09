package wechatminiprogram

import (
	"context"
	"errors"
	"time"
)

func (app *App) GetAccessTokenMonitor() (string, error) {
	if app.Redis.Db == nil {
		return "", errors.New("驱动没有初始化")
	}
	result := app.GetCallBackIp()
	if len(result.Result.IpList) <= 0 {
		token := app.CgiBinToken()
		app.Redis.Db.Set(context.Background(), app.getAccessTokenCacheKeyName(), token.Result.AccessToken, time.Second*7000)
		return token.Result.AccessToken, nil
	}
	return app.AccessToken, nil
}
