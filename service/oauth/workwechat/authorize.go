package workwechat

import "fmt"

// GetRedirectUri 构造网页授权链接 https://work.weixin.qq.com/api/doc/90000/90135/91022
func (app *App) GetRedirectUri() string {
	return fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect", app.AppID, app.RedirectUri)
}
