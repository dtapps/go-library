package wikeyun

// RestPowerEditCard 编辑电费充值卡
// https://open.wikeyun.cn/#/apiDocument/9/document/329
func (app *App) RestPowerEditCard(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request("https://router.wikeyun.cn/rest/Power/editCard", params)
	return request.ResponseBody, err
}
