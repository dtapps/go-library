package wikeyun

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type RestPowerForbidResponse struct {
	Code string   `json:"code"`
	Msg  string   `json:"msg"`
	Time string   `json:"time"`
	Data struct{} `json:"data"`
}

type RestPowerForbidResult struct {
	Result RestPowerForbidResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
}

func newRestPowerForbidResult(result RestPowerForbidResponse, body []byte, http gorequest.Response) *RestPowerForbidResult {
	return &RestPowerForbidResult{Result: result, Body: body, Http: http}
}

// RestPowerForbid 禁启用非API渠道电费充值
// status = 1 禁用 0启用
// https://open.wikeyun.cn/#/apiDocument/9/document/446
func (c *Client) RestPowerForbid(ctx context.Context, status int64, notMustParams ...gorequest.Params) (*RestPowerForbidResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "rest/Power/forbid")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.GetStoreId()) // 店铺ID
	params.Set("status", status)           // 1 禁用 0启用

	// 请求
	var response RestPowerForbidResponse
	request, err := c.request(ctx, "rest/Power/forbid", params, &response)
	return newRestPowerForbidResult(response, request.ResponseBody, request), err
}
