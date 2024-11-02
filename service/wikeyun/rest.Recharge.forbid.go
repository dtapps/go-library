package wikeyun

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type RestRechargeForbidResponse struct {
	Code string   `json:"code"`
	Msg  string   `json:"msg"`
	Time string   `json:"time"`
	Data struct{} `json:"data"`
}

type RestRechargeForbidResult struct {
	Result RestRechargeForbidResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
}

func newRestRechargeForbidResult(result RestRechargeForbidResponse, body []byte, http gorequest.Response) *RestRechargeForbidResult {
	return &RestRechargeForbidResult{Result: result, Body: body, Http: http}
}

// RestRechargeForbid 禁启用非API渠道下单
// status = 1 禁用 0启用
// https://open.wikeyun.cn/#/apiDocument/9/document/445
func (c *Client) RestRechargeForbid(ctx context.Context, status int64, notMustParams ...gorequest.Params) (*RestRechargeForbidResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.GetStoreId()) // 店铺ID
	params.Set("status", status)           // 1 禁用 0启用

	// 请求
	var response RestRechargeForbidResponse
	request, err := c.request(ctx, "rest/Recharge/forbid", params, &response)
	return newRestRechargeForbidResult(response, request.ResponseBody, request), err
}
