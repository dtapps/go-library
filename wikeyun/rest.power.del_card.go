package wikeyun

import (
	"encoding/json"
	"go.dtapp.net/library/gorequest"
)

type RestPowerDelCardResponse struct {
	Data string `json:"data"`
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
}

type RestPowerDelCardResult struct {
	Result RestPowerDelCardResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
	Err    error                    // 错误
}

func NewRestPowerDelCardResult(result RestPowerDelCardResponse, body []byte, http gorequest.Response, err error) *RestPowerDelCardResult {
	return &RestPowerDelCardResult{Result: result, Body: body, Http: http, Err: err}
}

// RestPowerDelCard 删除电费充值卡
// https://open.wikeyun.cn/#/apiDocument/9/document/330
func (app *App) RestPowerDelCard(cardId string) *RestPowerDelCardResult {
	// 参数
	param := NewParams()
	param.Set("card_id", cardId)
	params := app.NewParamsWith(param)
	// 请求
	request, err := app.request("https://router.wikeyun.cn/rest/Power/delCard", params)
	// 定义
	var response RestPowerDelCardResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewRestPowerDelCardResult(response, request.ResponseBody, request, err)
}
