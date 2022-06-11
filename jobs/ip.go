package jobs

import (
	"go.dtapp.net/library/goip"
	"gorm.io/gorm"
)

// RefreshIp 刷新Ip
func (app *App) RefreshIp(tx *gorm.DB) {
	xip := goip.GetOutsideIp()
	if app.OutsideIp == "" || app.OutsideIp == "0.0.0.0" {
		return
	}
	if app.OutsideIp == xip {
		return
	}
	tx.Where("ips = ?", app.OutsideIp).Delete(&TaskIp{}) // 删除
	app.OutsideIp = xip
}
