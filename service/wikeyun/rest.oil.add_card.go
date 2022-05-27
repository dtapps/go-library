package wikeyun

// RestOilCardAdd 添加充值卡
func (app *App) RestOilCardAdd(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request("https://router.wikeyun.cn/rest/Oil/addCard", params)
	return request.ResponseBody, err
}
