package tianyancha

import (
	"encoding/json"
	"net/http"
)

type EquityHumanIndexNodeResponse struct {
	IsLogin    int    `json:"isLogin"`
	Message    string `json:"message"`
	Special    string `json:"special"`
	State      string `json:"state"`
	VipMessage string `json:"vipMessage"`
}

type EquityHumanIndexNodeResult struct {
	Result EquityHumanIndexNodeResponse // 结果
	Body   []byte                       // 内容
	Err    error                        // 错误
}

func NewEquityHumanIndexNodeResult(result EquityHumanIndexNodeResponse, body []byte, err error) *EquityHumanIndexNodeResult {
	return &EquityHumanIndexNodeResult{Result: result, Body: body, Err: err}
}

func (app *App) EquityHumanIndexNode(notMustParams ...Params) *EquityHumanIndexNodeResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request("https://capi.tianyancha.com/cloud-equity-provider/v4/equity/humanIndexnode.json", params, http.MethodGet)
	// 定义
	var response EquityHumanIndexNodeResponse
	err = json.Unmarshal(body, &response)
	return NewEquityHumanIndexNodeResult(response, body, err)
}
