package pintoto

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type ApiOrderCreateSoonOrderResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ThirdOrderId string `json:"third_order_id"` // 接入方的订单号
		Ticket       string `json:"ticket"`
		TicketStatus int    `json:"ticketStatus"`
		OrderId      string `json:"order_id"`
	} `json:"data"`
	Success bool `json:"success"`
}

type ApiOrderCreateSoonOrderResult struct {
	Result ApiOrderCreateSoonOrderResponse // 结果
	Body   []byte                          // 内容
	Http   gorequest.Response              // 请求
}

func newApiOrderCreateSoonOrderResult(result ApiOrderCreateSoonOrderResponse, body []byte, http gorequest.Response) *ApiOrderCreateSoonOrderResult {
	return &ApiOrderCreateSoonOrderResult{Result: result, Body: body, Http: http}
}

// ApiOrderCreateSoonOrder 秒出单下单 https://www.showdoc.com.cn/1154868044931571/6437295495912025
func (c *Client) ApiOrderCreateSoonOrder(ctx context.Context, notMustParams ...gorequest.Params) (*ApiOrderCreateSoonOrderResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/api/order/create-soon-order", params)
	if err != nil {
		return newApiOrderCreateSoonOrderResult(ApiOrderCreateSoonOrderResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ApiOrderCreateSoonOrderResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newApiOrderCreateSoonOrderResult(response, request.ResponseBody, request), err
}
