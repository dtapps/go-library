package wikeyun

// RestOilCardInfo 油卡详情
func (app *App) RestOilCardInfo(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Oil/cardInfo", params)
	return body, err
}
