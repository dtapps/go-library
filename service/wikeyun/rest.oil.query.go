package wikeyun

// RestOilOrderQuery 订单查询
func (app *App) RestOilOrderQuery(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request("https://router.wikeyun.cn/rest/Oil/query", params)
	return request.ResponseBody, err
}
