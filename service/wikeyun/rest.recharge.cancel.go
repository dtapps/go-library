package wikeyun

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
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
}

func newRestRechargeCancelResult(result RestRechargeCancelResponse, body []byte, http gorequest.Response) *RestRechargeCancelResult {
	return &RestRechargeCancelResult{Result: result, Body: body, Http: http}
}

// RestRechargeCancel 话费订单取消
// https://open.wikeyun.cn/#/apiDocument/9/document/300
func (c *Client) RestRechargeCancel(ctx context.Context, orderNumber string, notMustParams ...*gorequest.Params) (*RestRechargeCancelResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("order_number", orderNumber) // 取消的单号，多个用英文逗号隔开
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Recharge/cancel", params)
	if err != nil {
		return newRestRechargeCancelResult(RestRechargeCancelResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RestRechargeCancelResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestRechargeCancelResult(response, request.ResponseBody, request), err
}
