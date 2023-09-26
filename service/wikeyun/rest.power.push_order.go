package wikeyun

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type RestPowerPushOrderResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		OrderNumber string `json:"order_number"`
	} `json:"data"`
}

type RestPowerPushOrderResult struct {
	Result RestPowerPushOrderResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
}

func newRestPowerPushOrderResult(result RestPowerPushOrderResponse, body []byte, http gorequest.Response) *RestPowerPushOrderResult {
	return &RestPowerPushOrderResult{Result: result, Body: body, Http: http}
}

// RestPowerPushOrder 电费充值API
// https://open.wikeyun.cn/#/apiDocument/9/document/311
func (c *Client) RestPowerPushOrder(ctx context.Context, notMustParams ...gorequest.Params) (*RestPowerPushOrderResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.GetStoreId()) // 店铺ID
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Power/pushOrder", params)
	if err != nil {
		return newRestPowerPushOrderResult(RestPowerPushOrderResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RestPowerPushOrderResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestPowerPushOrderResult(response, request.ResponseBody, request), err
}
