package jobs

import "go.dtapp.net/library/utils/goip"

func (app *App) RefreshIp() {
	xip := goip.GetOutsideIp()
	if app.OutsideIp == "" || app.OutsideIp == "0.0.0.0" {
		return
	}
	if app.OutsideIp == xip {
		return
	}
	app.Db.Where("ips = ?", app.OutsideIp).Delete(&TaskIp{}) // 删除
	app.OutsideIp = xip
}
