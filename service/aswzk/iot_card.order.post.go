package aswzk

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type IotCardOrderPostResponse struct {
	Code    int         `json:"code"`
	Info    string      `json:"info"`
	Data    interface{} `json:"data"`
	Time    int64       `json:"time"`
	TraceID string      `json:"trace_id"`
}

// IotCardOrderPost 物联卡订单下单
func (c *Client) IotCardOrderPost(ctx context.Context, notMustParams ...*gorequest.Params) (response IotCardOrderPostResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "iot_card/order", params, http.MethodPost, &response)
	return
}
