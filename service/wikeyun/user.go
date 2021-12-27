package wikeyun

import (
	"encoding/json"
	"errors"
)

// UserQueryResponse 返回参数
type UserQueryResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Money string `json:"money"`
		ID    int    `json:"id"`
	} `json:"data"`
}

// UserQuery 查询余额接口
func (app *App) UserQuery() (result UserQueryResponse, err error) {
	// 请求
	body, err := app.request("https://router.wikeyun.cn/rest/User/query", map[string]interface{}{})
	if err != nil {
		return result, errors.New(err.Error())
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, errors.New(err.Error())
	}
	return result, nil
}
