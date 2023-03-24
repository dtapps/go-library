package wikeyun

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type RestPowerCancelResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type RestPowerCancelResult struct {
	Result RestPowerCancelResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
	Err    error                   // 错误
}

func newRestPowerCancelResult(result RestPowerCancelResponse, body []byte, http gorequest.Response, err error) *RestPowerCancelResult {
	return &RestPowerCancelResult{Result: result, Body: body, Http: http, Err: err}
}

// RestPowerCancel 电费订单取消
// https://open.wikeyun.cn/#/apiDocument/9/document/323
func (c *Client) RestPowerCancel(ctx context.Context, orderNumber string) *RestPowerCancelResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("order_number", orderNumber) // 取消的单号，多个用英文逗号隔开
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Power/cancel", params)
	// 定义
	var response RestPowerCancelResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestPowerCancelResult(response, request.ResponseBody, request, err)
}
