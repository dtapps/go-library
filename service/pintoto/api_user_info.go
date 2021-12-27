package pintoto

import (
	"encoding/json"
)

type ApiUserInfoResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Nickname     string  `json:"nickname"`      // 用户昵称
		Mobile       int64   `json:"mobile"`        // 注册号码
		Balance      float64 `json:"balance"`       // 账户余额
		FreezeAmount float64 `json:"freeze_amount"` // 冻结金额
	} `json:"data"`
	Code int `json:"code"`
}

// ApiUserInfo 账号信息查询 https://www.showdoc.com.cn/1154868044931571/6269224958928211
func (app *App) ApiUserInfo() (result ApiUserInfoResult, err error) {
	body, err := app.request("https://movieapi2.pintoto.cn/api/user/info", map[string]interface{}{})
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
