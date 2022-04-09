package wechatopen

import (
	"gitee.com/dtapps/go-library/utils/gotime"
	"gorm.io/gorm"
	"time"
)

// GetComponentAccessTokenMonitor 获取令牌和监控
func (app *App) GetComponentAccessTokenMonitor() string {
	// 查询
	componentAccessToken := app.GetComponentAccessToken()
	// 判断
	result := app.CgiBinGetApiDomainIp(componentAccessToken)
	if len(result.Result.IpList) > 0 {
		return componentAccessToken
	}
	// 重新获取
	return app.SetComponentAccessToken(app.CgiBinComponentApiComponentToken())
}

// SetComponentAccessToken 设置令牌
func (app *App) SetComponentAccessToken(info *CgiBinComponentApiComponentTokenResult) string {
	if app.Db == nil || info.Result.ComponentAccessToken == "" {
		return ""
	}
	app.Db.Create(&ComponentAccessToken{
		AppId:                app.ComponentAppId,
		ComponentAccessToken: info.Result.ComponentAccessToken,
		ExpiresIn:            info.Result.ExpiresIn,
		ExpireTime:           gotime.Current().AfterSeconds(7200).Time,
	})
	return info.Result.ComponentAccessToken
}

type ComponentAccessToken struct {
	gorm.Model
	AppId                string    `json:"app_id"`                 // 第三方平台 appid
	ComponentAccessToken string    `json:"component_access_token"` // 第三方平台 access_token
	ExpiresIn            int64     `json:"expires_in"`             // 有效期，单位：秒
	ExpireTime           time.Time `json:"expire_time"`            // 过期时间
}

func (m *ComponentAccessToken) TableName() string {
	return "component_access_token"
}
