package aswzk

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type IotCardOrderPostResponse struct {
	Code    int         `json:"code"`
	Info    string      `json:"info"`
	Data    interface{} `json:"data"`
	Time    int64       `json:"time"`
	TraceID string      `json:"trace_id"`
}

type IotCardOrderPostResult struct {
	Result IotCardOrderPostResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
}

func newIotCardOrderPostResult(result IotCardOrderPostResponse, body []byte, http gorequest.Response) *IotCardOrderPostResult {
	return &IotCardOrderPostResult{Result: result, Body: body, Http: http}
}

// IotCardOrderPost 物联卡订单下单
func (c *Client) IotCardOrderPost(ctx context.Context, notMustParams ...*gorequest.Params) (*IotCardOrderPostResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response IotCardOrderPostResponse
	request, err := c.request(ctx, "iot_card/order", params, http.MethodPost, &response)
	return newIotCardOrderPostResult(response, request.ResponseBody, request), err
}
