package aswzk

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type ElectricityBillOrderResponse struct {
	Code    int         `json:"code"`
	Info    string      `json:"info"`
	Data    interface{} `json:"data"`
	Time    int64       `json:"time"`
	TraceID string      `json:"trace_id"`
}

type ElectricityBillOrderResult struct {
	Result ElectricityBillOrderResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
}

func newElectricityBillOrderResult(result ElectricityBillOrderResponse, body []byte, http gorequest.Response) *ElectricityBillOrderResult {
	return &ElectricityBillOrderResult{Result: result, Body: body, Http: http}
}

// ElectricityBillOrder 电费订单下单
func (c *Client) ElectricityBillOrder(ctx context.Context, notMustParams ...*gorequest.Params) (*ElectricityBillOrderResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response ElectricityBillOrderResponse
	request, err := c.request(ctx, "electricity_bill/order", params, http.MethodPost, &response)
	return newElectricityBillOrderResult(response, request.ResponseBody, request), err
}
