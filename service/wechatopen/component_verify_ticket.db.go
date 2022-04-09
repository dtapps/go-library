package wechatopen

import (
	"github.com/dtapps/go-library/utils/gotime"
	"gorm.io/gorm"
	"time"
)

// SetComponentVerifyTicket 设置微信后台推送的ticket
func (app *App) SetComponentVerifyTicket(info *ResponseServeHttpVerifyTicket) string {
	if info.ComponentVerifyTicket == "" {
		return ""
	}
	app.Db.Create(&ComponentVerifyTicket{
		AppId:                 info.AppId,
		CreateTime:            info.CreateTime,
		InfoType:              info.InfoType,
		ComponentVerifyTicket: info.ComponentVerifyTicket,
		ExpireTime:            gotime.Current().AfterHour(12).Time,
	})
	return info.ComponentVerifyTicket
}

type ComponentVerifyTicket struct {
	gorm.Model
	AppId                 string    `json:"app_id"`                  // 第三方平台 appid
	CreateTime            int64     `json:"create_time"`             // 时间戳，单位：s
	InfoType              string    `json:"info_type"`               // 固定为："component_verify_ticket"
	ComponentVerifyTicket string    `json:"component_verify_ticket"` // Ticket 内容
	ExpireTime            time.Time `json:"expire_time"`             // 过期时间
}

func (m *ComponentVerifyTicket) TableName() string {
	return "component_verify_ticket"
}
