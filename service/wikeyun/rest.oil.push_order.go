package wikeyun

// RestOilOrderPush 充值下单
func (app *App) RestOilOrderPush(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Oil/pushOrder", params)
	return body, err
}
