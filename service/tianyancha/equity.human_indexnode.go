package tianyancha

import (
	"net/http"
)

type EquityHumanIndexNodeResult struct {
	IsLogin    int    `json:"isLogin"`
	Message    string `json:"message"`
	Special    string `json:"special"`
	State      string `json:"state"`
	VipMessage string `json:"vipMessage"`
}

func (app *App) EquityHumanIndexNode(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("https://capi.tianyancha.com/cloud-equity-provider/v4/equity/humanIndexnode.json", params, http.MethodGet)
	return body, err
}
