package wechatoffice

import (
	"github.com/dtapps/go-library/utils/gotime"
	"time"
)

type WechatJsapiTicketDbModel struct {
	ID          int64     `json:"id"`
	AppID       string    `json:"app_id"`
	AppSecret   string    `json:"app_secret"`
	Ticket      string    `json:"ticket"`
	ExpiresIn   int       `json:"expires_in"`
	ExpiresTime string    `json:"expires_time"`
	CreateTime  time.Time `gorm:"autoCreateTime" json:"-"`
}

func (m *WechatJsapiTicketDbModel) TableName() string {
	return "wechat_jsapi_ticket"
}

func (app *App) wechatJsapiTicketModelTake() (result WechatJsapiTicketDbModel) {
	app.Db.Where("app_id = ?", app.AppId).Where("expires_time > ?", gotime.Current().Format()).Take(&result)
	return result
}

// GetJsapiTicketDb 获取api_ticket
func (app *App) GetJsapiTicketDb() string {
	wat := app.wechatJsapiTicketModelTake()
	if wat.Ticket != "" {
		return wat.Ticket
	} else {
		token := app.GetTicket(app.GetAccessTokenDb(), "jsapi")
		if token.Result.Ticket == "" {
			return ""
		} else {
			// 创建
			app.Db.Create(&WechatJsapiTicketDbModel{
				AppID:       app.AppId,
				AppSecret:   app.AppSecret,
				Ticket:      token.Result.Ticket,
				ExpiresIn:   token.Result.ExpiresIn,
				ExpiresTime: gotime.Current().AfterSeconds(7000).Format(),
			})
			return token.Result.Ticket
		}
	}
}
