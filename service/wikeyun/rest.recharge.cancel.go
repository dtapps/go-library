package wikeyun

import (
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
)

type RestRechargeCancelResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type RestRechargeCancelResult struct {
	Result RestRechargeCancelResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
	Err    error                      // 错误
}

func newRestRechargeCancelResult(result RestRechargeCancelResponse, body []byte, http gorequest.Response, err error) *RestRechargeCancelResult {
	return &RestRechargeCancelResult{Result: result, Body: body, Http: http, Err: err}
}

// RestRechargeCancel 话费订单取消
// https://open.wikeyun.cn/#/apiDocument/9/document/300
func (c *Client) RestRechargeCancel(orderNumber string) *RestRechargeCancelResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("order_number", orderNumber) // 取消的单号，多个用英文逗号隔开
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(apiUrl+"/rest/Recharge/cancel", params)
	// 定义
	var response RestRechargeCancelResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newRestRechargeCancelResult(response, request.ResponseBody, request, err)
}
