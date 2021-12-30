package wechatoffice

import (
	"errors"
	"fmt"
	"gopkg.in/dtapps/go-library.v3/utils/goredis"
	"gopkg.in/dtapps/go-library.v3/utils/gotime"
	"time"
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
			token := app.AuthGetAccessToken()
			if token.Result.AccessToken == "" {
				return errors.New("获取AccessToken失败")
			} else {
				app.Db.Create(&WechatAccessTokenDbModel{
					AppID:       app.AppId,
					AppSecret:   app.AppSecret,
					AccessToken: token.Result.AccessToken,
					ExpiresIn:   token.Result.ExpiresIn,
					ExpiresTime: gotime.Current().AfterSeconds(7000).Format(),
				})
				return nil
			}
		case QdTypeRdb:
			if app.RDb == nil {
				return errors.New("驱动没有初始化")
			}
			cacheName := fmt.Sprintf("wechat_access_token:%v", app.AppId)
			redis := goredis.App{
				Rdb: app.RDb,
			}
			token := app.AuthGetAccessToken()
			if token.Result.AccessToken == "" {
				return errors.New("获取AccessToken失败")
			}
			redis.NewStringOperation().Set(cacheName, token.Result.AccessToken, goredis.WithExpire(time.Minute*7000))
			return nil
		default:
			return errors.New("驱动类型不在范围内")
		}
	}
	return nil
}
