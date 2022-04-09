package wikeyun

// PowerCardInfo 充值卡详情
func (app *App) PowerCardInfo(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Power/cardInfo", params)
	return body, err
}
