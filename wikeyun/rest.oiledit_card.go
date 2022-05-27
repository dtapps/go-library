package wikeyun

// RestOilCardEdit 编辑充值卡
func (app *App) RestOilCardEdit(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request("https://router.wikeyun.cn/rest/Oil/editCard", params)
	return request.ResponseBody, err
}
