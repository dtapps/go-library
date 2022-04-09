package wikeyun

// RestPowerDelCard 充值卡删除
func (app *App) RestPowerDelCard(cardId string) (body []byte, err error) {
	// 参数
	param := NewParams()
	param.Set("card_id", cardId)
	params := app.NewParamsWith(param)
	// 请求
	body, err = app.request("https://router.wikeyun.cn/rest/Power/delCard", params)
	return body, err
}
