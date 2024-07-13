package aswzk

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PhoneBillOrderResponse struct {
	Code    int         `json:"code"`
	Info    string      `json:"info"`
	Data    interface{} `json:"data"`
	Time    int64       `json:"time"`
	TraceID string      `json:"trace_id"`
}

type PhoneBillOrderResult struct {
	Result PhoneBillOrderResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newPhoneBillOrderResult(result PhoneBillOrderResponse, body []byte, http gorequest.Response) *PhoneBillOrderResult {
	return &PhoneBillOrderResult{Result: result, Body: body, Http: http}
}

// PhoneBillOrder 话费订单下单
func (c *Client) PhoneBillOrder(ctx context.Context, notMustParams ...gorequest.Params) (*PhoneBillOrderResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "phone_bill/order")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response PhoneBillOrderResponse
	request, err := c.request(ctx, "phone_bill/order", params, http.MethodPost, &response)
	return newPhoneBillOrderResult(response, request.ResponseBody, request), err
}
