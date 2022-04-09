package wikeyun

import "encoding/json"

type RestPowerAddCardResponse struct {
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

type RestPowerAddCardResult struct {
	Result RestPowerAddCardResponse // 结果
	Body   []byte                   // 内容
	Err    error                    // 错误
}

func NewRestPowerAddCardResult(result RestPowerAddCardResponse, body []byte, err error) *RestPowerAddCardResult {
	return &RestPowerAddCardResult{Result: result, Body: body, Err: err}
}

// RestPowerAddCard 添加充值卡
func (app *App) RestPowerAddCard(notMustParams ...Params) *RestPowerAddCardResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request("https://router.wikeyun.cn/rest/Power/addCard", params)
	// 定义
	var response RestPowerAddCardResponse
	err = json.Unmarshal(body, &response)
	return NewRestPowerAddCardResult(response, body, err)
}
