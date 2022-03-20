package wechatoffice

import (
	"github.com/dtapps/go-library/utils/gotime"
	"time"
)

type WechatAccessTokenDbModel struct {
	ID          int64     `json:"id"`
	AppID       string    `json:"app_id"`
	AppSecret   string    `json:"app_secret"`
	AccessToken string    `json:"access_token"`
	ExpiresIn   int       `json:"expires_in"`
	ExpiresTime string    `json:"expires_time"`
	CreateTime  time.Time `gorm:"autoCreateTime" json:"-"`
}

func (m *WechatAccessTokenDbModel) TableName() string {
	return "wechat_access_token"
}

func (app *App) wechatAccessTokenModelTake() (result WechatAccessTokenDbModel) {
	app.Db.Where("app_id = ?", app.AppId).Where("expires_time > ?", gotime.Current().Format()).Take(&result)
	return result
}

func (app *App) GetAccessTokenDb() string {
	wat := app.wechatAccessTokenModelTake()
	if wat.AccessToken != "" {
		return wat.AccessToken
	} else {
		token := app.AuthGetAccessToken()
		if token.Result.AccessToken == "" {
			return ""
		} else {
			// 创建
			app.Db.Create(&WechatAccessTokenDbModel{
				AppID:       app.AppId,
				AppSecret:   app.AppSecret,
				AccessToken: token.Result.AccessToken,
				ExpiresIn:   token.Result.ExpiresIn,
				ExpiresTime: gotime.Current().AfterSeconds(7000).Format(),
			})
			return token.Result.AccessToken
		}
	}
}
