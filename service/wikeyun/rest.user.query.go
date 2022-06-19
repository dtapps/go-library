package wikeyun

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
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

func newRestUserQueryResult(result RestUserQueryResponse, body []byte, http gorequest.Response, err error) *RestUserQueryResult {
	return &RestUserQueryResult{Result: result, Body: body, Http: http, Err: err}
}

// RestUserQuery 用户信息
// https://open.wikeyun.cn/#/apiDocument/10/document/336
func (c *Client) RestUserQuery() *RestUserQueryResult {
	// 请求
	request, err := c.request(apiUrl+"/rest/User/query", map[string]interface{}{})
	// 定义
	var response RestUserQueryResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newRestUserQueryResult(response, request.ResponseBody, request, err)
}
