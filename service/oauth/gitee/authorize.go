package gitee

import "fmt"

// GetRedirectUri 获取登录地址
func (app *App) GetRedirectUri() string {
	return fmt.Sprintf("https://gitee.com/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code", app.ClientID, app.RedirectUri)
}
