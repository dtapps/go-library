package wechatopen

// SetAuthorizerAppid 设置代理商小程序
func (app *App) SetAuthorizerAppid(authorizerAppid string) {
	app.authorizerAppid = authorizerAppid
	return
}
