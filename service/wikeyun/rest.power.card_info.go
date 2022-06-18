package wikeyun

// PowerCardInfo 电费充值卡详情
// https://open.wikeyun.cn/#/apiDocument/9/document/333
func (app *App) PowerCardInfo(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request("https://router.wikeyun.cn/rest/Power/cardInfo", params)
	return request.ResponseBody, err
}
