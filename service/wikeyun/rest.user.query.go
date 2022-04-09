package wikeyun

import (
	"encoding/json"
)

type RestUserQueryResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Money string `json:"money"`
		ID    int    `json:"id"`
	} `json:"data"`
}

type RestUserQueryResult struct {
	Result RestUserQueryResponse // 结果
	Body   []byte                // 内容
	Err    error                 // 错误
}

func NewRestUserQueryResult(result RestUserQueryResponse, body []byte, err error) *RestUserQueryResult {
	return &RestUserQueryResult{Result: result, Body: body, Err: err}
}

// RestUserQuery 查询余额接口
func (app *App) RestUserQuery() *RestUserQueryResult {
	// 请求
	body, err := app.request("https://router.wikeyun.cn/rest/User/query", map[string]interface{}{})
	// 定义
	var response RestUserQueryResponse
	err = json.Unmarshal(body, &response)
	return NewRestUserQueryResult(response, body, err)
}
