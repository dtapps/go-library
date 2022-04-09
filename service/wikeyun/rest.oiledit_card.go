package wikeyun

// RestOilCardEdit 编辑充值卡
func (app *App) RestOilCardEdit(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Oil/editCard", params)
	return body, err
}
