package pintoto

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
)

type ApiUserInfoResponse struct {
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

type ApiUserInfoResult struct {
	Result ApiUserInfoResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
	Err    error               // 错误
}

func NewApiUserInfoResult(result ApiUserInfoResponse, body []byte, http gorequest.Response, err error) *ApiUserInfoResult {
	return &ApiUserInfoResult{Result: result, Body: body, Http: http, Err: err}
}

// ApiUserInfo 账号信息查询 https://www.showdoc.com.cn/1154868044931571/6269224958928211
func (app *App) ApiUserInfo() *ApiUserInfoResult {
	request, err := app.request("https://movieapi2.pintoto.cn/api/user/info", map[string]interface{}{})
	// 定义
	var response ApiUserInfoResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewApiUserInfoResult(response, request.ResponseBody, request, err)
}
