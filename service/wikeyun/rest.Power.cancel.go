package wikeyun

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
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
}

func newRestPowerCancelResult(result RestPowerCancelResponse, body []byte, http gorequest.Response) *RestPowerCancelResult {
	return &RestPowerCancelResult{Result: result, Body: body, Http: http}
}

// RestPowerCancel 电费订单取消
// order_number = 取消的单号，多个用英文逗号隔开
// https://open.wikeyun.cn/#/apiDocument/9/document/323
func (c *Client) RestPowerCancel(ctx context.Context, orderNumber string, notMustParams ...gorequest.Params) (*RestPowerCancelResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "rest/Power/cancel")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("order_number", orderNumber) // 取消的单号，多个用英文逗号隔开

	// 请求
	var response RestPowerCancelResponse
	request, err := c.request(ctx, "rest/Power/cancel", params, &response)
	return newRestPowerCancelResult(response, request.ResponseBody, request), err
}
