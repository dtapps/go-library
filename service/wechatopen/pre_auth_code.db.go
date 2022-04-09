package wechatopen

import (
	"gitee.com/dtapps/go-library/utils/gotime"
	"gorm.io/gorm"
	"time"
)

// GetPreAuthCodeMonitor 获取预授权码和监控
func (app *App) GetPreAuthCodeMonitor() string {
	// 查询
	preAuthCode := app.GetPreAuthCode()
	if preAuthCode != "" {
		return preAuthCode
	}
	// 重新获取
	return app.SetPreAuthCode(app.CgiBinComponentApiCreatePreAuthCoden())
}

// SetPreAuthCode 设置预授权码和自动获取
func (app *App) SetPreAuthCode(info *CgiBinComponentApiCreatePreAuthCodenResult) string {
	if app.Db == nil || info.Result.PreAuthCode == "" {
		return ""
	}
	app.Db.Create(&PreAuthCode{
		AppId:       app.ComponentAppId,
		PreAuthCode: info.Result.PreAuthCode,
		ExpiresIn:   info.Result.ExpiresIn,
		ExpireTime:  gotime.Current().AfterSeconds(1700).Time,
	})
	return info.Result.PreAuthCode
}

type PreAuthCode struct {
	gorm.Model
	AppId       string    `json:"app_id"`        // 第三方平台 appid
	PreAuthCode string    `json:"pre_auth_code"` // 预授权码
	ExpiresIn   int64     `json:"expires_in"`    // 有效期，单位：秒
	ExpireTime  time.Time `json:"expire_time"`   // 过期时间
}

func (m *PreAuthCode) TableName() string {
	return "pre_auth_code"
}

// PreAuthCodeDelete 删除过期或使用过的预授权码
func (app *App) PreAuthCodeDelete(id uint) int64 {
	return app.Db.Where("id = ?", id).Delete(&PreAuthCode{}).RowsAffected
}
