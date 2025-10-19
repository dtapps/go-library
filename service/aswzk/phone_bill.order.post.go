package aswzk

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type PhoneBillOrderResponse struct {
	Code    int         `json:"code"`
	Info    string      `json:"info"`
	Data    interface{} `json:"data"`
	Time    int64       `json:"time"`
	TraceID string      `json:"trace_id"`
}

// PhoneBillOrder 话费订单下单
func (c *Client) PhoneBillOrder(ctx context.Context, notMustParams ...*gorequest.Params) (response PhoneBillOrderResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "phone_bill/order", params, http.MethodPost, &response)
	return
}
