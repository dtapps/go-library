package workwechat

import "fmt"

// QrConnect 构造独立窗口登录二维码 https://open.work.weixin.qq.com/api/doc/90000/90135/91019
func (app *App) QrConnect() string {
	return fmt.Sprintf("https://open.work.weixin.qq.com/wwopen/sso/qrConnect?appid=%s&agentid=%d&redirect_uri=%s&state=STATE&lang=zh", app.AppID, app.AgentID, app.RedirectUri)
}
