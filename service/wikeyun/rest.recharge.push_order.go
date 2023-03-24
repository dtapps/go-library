package wikeyun

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type RestRechargePushOrderResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderNumber string `json:"order_number"`
	} `json:"data"`
}

type RestRechargePushOrderResult struct {
	Result RestRechargePushOrderResponse // 结果
	Body   []byte                        // 内容
	Http   gorequest.Response            // 请求
	Err    error                         // 错误
}

func newRestRechargePushOrderResult(result RestRechargePushOrderResponse, body []byte, http gorequest.Response, err error) *RestRechargePushOrderResult {
	return &RestRechargePushOrderResult{Result: result, Body: body, Http: http, Err: err}
}

// RestRechargePushOrder 话费充值推送
// https://open.wikeyun.cn/#/apiDocument/9/document/298
func (c *Client) RestRechargePushOrder(ctx context.Context, notMustParams ...gorequest.Params) *RestRechargePushOrderResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.GetStoreId()) // 店铺ID
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Recharge/pushOrder", params)
	// 定义
	var response RestRechargePushOrderResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestRechargePushOrderResult(response, request.ResponseBody, request, err)
}
