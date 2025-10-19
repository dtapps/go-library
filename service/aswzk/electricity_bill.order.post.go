package aswzk

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type ElectricityBillOrderResponse struct {
	Code    int         `json:"code"`
	Info    string      `json:"info"`
	Data    interface{} `json:"data"`
	Time    int64       `json:"time"`
	TraceID string      `json:"trace_id"`
}

// ElectricityBillOrder 电费订单下单
func (c *Client) ElectricityBillOrder(ctx context.Context, notMustParams ...*gorequest.Params) (response ElectricityBillOrderResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "electricity_bill/order", params, http.MethodPost, &response)
	return
}
