package wikeyun

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
)

type RestPowerQueryResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		Id            uint   `json:"id"`
		Fanli         string `json:"fanli"`
		Amount        int64  `json:"amount"`
		Status        int    `json:"status"`
		CardId        string `json:"card_id"`
		OrderNo       string `json:"order_no"`
		CostPrice     string `json:"cost_price"`
		OrderNumber   string `json:"order_number"`
		ArrivedAmount int64  `json:"arrived_amount"`
		Reason        string `json:"reason"`
	} `json:"data"`
}

type RestPowerQueryResult struct {
	Result RestPowerQueryResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
	Err    error                  // 错误
}

func newRestPowerQueryResult(result RestPowerQueryResponse, body []byte, http gorequest.Response, err error) *RestPowerQueryResult {
	return &RestPowerQueryResult{Result: result, Body: body, Http: http, Err: err}
}

// RestPowerQuery 电费订单查询
// https://open.wikeyun.cn/#/apiDocument/9/document/313
func (c *Client) RestPowerQuery(ctx context.Context, orderNumber string) *RestPowerQueryResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("order_number", orderNumber) // 平台单号
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Power/query", params)
	// 定义
	var response RestPowerQueryResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newRestPowerQueryResult(response, request.ResponseBody, request, err)
}
