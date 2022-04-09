package wikeyun

// RestOilCardDel 油卡删除
func (app *App) RestOilCardDel(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Oil/delCard", params)
	return body, err
}
