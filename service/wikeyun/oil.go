package wikeyun

// OilCardAdd 添加充值卡
func (app *App) OilCardAdd(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Oil/addCard", params)
	return body, err
}

// OilCardEdit 编辑充值卡
func (app *App) OilCardEdit(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Oil/editCard", params)
	return body, err
}

// OilCardDel 油卡删除
func (app *App) OilCardDel(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Oil/delCard", params)
	return body, err
}

// OilCardInfo 油卡详情
func (app *App) OilCardInfo(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Oil/cardInfo", params)
	return body, err
}

// OilOrderPush 充值下单
func (app *App) OilOrderPush(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Oil/pushOrder", params)
	return body, err
}

// OilOrderQuery 订单查询
func (app *App) OilOrderQuery(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Oil/query", params)
	return body, err
}
