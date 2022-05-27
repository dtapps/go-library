package wikeyun

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
)

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
	Http   gorequest.Response       // 请求
	Err    error                    // 错误
}

func NewRestPowerAddCardResult(result RestPowerAddCardResponse, body []byte, http gorequest.Response, err error) *RestPowerAddCardResult {
	return &RestPowerAddCardResult{Result: result, Body: body, Http: http, Err: err}
}

// RestPowerAddCard 添加电费充值卡
// https://open.wikeyun.cn/#/apiDocument/9/document/326
func (app *App) RestPowerAddCard(notMustParams ...Params) *RestPowerAddCardResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	params.Set("store_id", app.storeId) // 店铺ID
	// 请求
	request, err := app.request("https://router.wikeyun.cn/rest/Power/addCard", params)
	// 定义
	var response RestPowerAddCardResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewRestPowerAddCardResult(response, request.ResponseBody, request, err)
}
