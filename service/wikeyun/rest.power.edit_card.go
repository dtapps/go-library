package wikeyun

// RestPowerEditCard 编辑充值卡
func (app *App) RestPowerEditCard(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Power/editCard", params)
	return body, err
}
