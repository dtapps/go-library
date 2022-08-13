package wikeyun

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
)

type RestRechargeQueryResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		Id             uint   `json:"id"`
		Fanli          string `json:"fanli"`
		Amount         int64  `json:"amount"`
		Mobile         string `json:"mobile"`
		Status         int    `json:"status"`
		OrderNo        string `json:"order_no"`
		CostPrice      string `json:"cost_price"`
		OrderNumber    string `json:"order_number"`
		OrgOrderNumber string `json:"org_order_number"`
	} `json:"data"`
}

type RestRechargeQueryResult struct {
	Result RestRechargeQueryResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
	Err    error                     // 错误
}

func newRestRechargeQueryResult(result RestRechargeQueryResponse, body []byte, http gorequest.Response, err error) *RestRechargeQueryResult {
	return &RestRechargeQueryResult{Result: result, Body: body, Http: http, Err: err}
}

// RestRechargeQuery 话费订单查询
// https://open.wikeyun.cn/#/apiDocument/9/document/299
func (c *Client) RestRechargeQuery(ctx context.Context, orderNumber string) *RestRechargeQueryResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("order_number", orderNumber) // 平台订单号
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Recharge/query", params)
	// 定义
	var response RestRechargeQueryResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newRestRechargeQueryResult(response, request.ResponseBody, request, err)
}
