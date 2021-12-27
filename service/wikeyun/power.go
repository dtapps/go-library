package wikeyun

type PowerAddCardResult struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		CardNum    string `json:"card_num"`
		StoreId    string `json:"store_id"`
		CreateTime int    `json:"create_time"`
		Type       int    `json:"type"` // 缴费单位
		CmsUid     int    `json:"cms_uid"`
		Province   string `json:"province"` // 缴费省份
		City       string `json:"city"`     // 缴费城市
		Id         string `json:"id"`       // 缴费卡编号
	} `json:"data"`
}

// PowerAddCard 添加充值卡
func (app *App) PowerAddCard(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Power/addCard", params)
	return body, err
}

// PowerEditCard 编辑充值卡
func (app *App) PowerEditCard(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Power/editCard", params)
	return body, err
}

// PowerDelCard 充值卡删除
func (app *App) PowerDelCard(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Power/delCard", params)
	return body, err
}

// PowerCardInfo 充值卡详情
func (app *App) PowerCardInfo(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Power/cardInfo", params)
	return body, err
}

type PowerPushOrderResult struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		OrderNumber string `json:"order_number"`
	} `json:"data"`
}

// PowerPushOrder 充值下单
func (app *App) PowerPushOrder(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Power/pushOrder", params)
	return body, err
}

type PowerQueryResult struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		OrderNumber string `json:"order_number"` // 订单号
		OrderNo     string `json:"order_no"`     // 订单号
		CardId      string `json:"card_id"`      // 卡编号
		Amount      int    `json:"amount"`       // 充值金额
		CostPrice   string `json:"cost_price"`   // 成本价
		Fanli       string `json:"fanli"`        // 平台返利
		Status      int    `json:"status"`       // 交易结果（0 待支付 1 已付充值中 2充值成功 3充值失败需要退款 4退款成功 6 待充值 7 已匹配）
	} `json:"data"`
}

func (app *App) PowerQuery(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Power/query", params)
	return body, err
}
