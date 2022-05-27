package wikeyun

import (
	"encoding/json"
	"go.dtapp.net/library/gorequest"
)

type RestUserQueryResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Id     string `json:"id"`
		Avatar string `json:"avatar"`
		Money  string `json:"money"`
		Mobile string `json:"mobile"`
	} `json:"data"`
}

type RestUserQueryResult struct {
	Result RestUserQueryResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
	Err    error                 // 错误
}

func NewRestUserQueryResult(result RestUserQueryResponse, body []byte, http gorequest.Response, err error) *RestUserQueryResult {
	return &RestUserQueryResult{Result: result, Body: body, Http: http, Err: err}
}

// RestUserQuery 用户信息
// https://open.wikeyun.cn/#/apiDocument/10/document/336
func (app *App) RestUserQuery() *RestUserQueryResult {
	// 请求
	request, err := app.request("https://router.wikeyun.cn/rest/User/query", map[string]interface{}{})
	// 定义
	var response RestUserQueryResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewRestUserQueryResult(response, request.ResponseBody, request, err)
}
