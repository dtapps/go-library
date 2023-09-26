package nldyp

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type PartnerData4GetOrderStatusResponse struct {
	Status int `json:"status"`
	Data   struct {
		OrderNo     string      `json:"orderNo"` // 系统商锁座订单号
		PrintNo     interface{} `json:"printNo"`
		OrderStatus int         `json:"orderStatus"` //
		PlanOrderNo interface{} `json:"planOrderNo"`
		TicketInfo  []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"ticketInfo"`
		Tickets struct {
			Field1 struct {
				SeatNo     string      `json:"seatNo"`
				TicketCode interface{} `json:"ticketCode"`
				InfoCode   interface{} `json:"infoCode"`
				PrintFlag  interface{} `json:"printFlag"`
			} `json:"0000000000000001-1-05"`
			Field2 struct {
				SeatNo     string      `json:"seatNo"`
				TicketCode interface{} `json:"ticketCode"`
				InfoCode   interface{} `json:"infoCode"`
				PrintFlag  interface{} `json:"printFlag"`
			} `json:"0000000000000001-1-06"`
		} `json:"tickets"`
	} `json:"data"`
	Content string `json:"content"`
}

type PartnerData4GetOrderStatusResult struct {
	Result PartnerData4GetOrderStatusResponse // 结果
	Body   []byte                             // 内容
	Http   gorequest.Response                 // 请求
}

func newPartnerData4GetOrderStatusResult(result PartnerData4GetOrderStatusResponse, body []byte, http gorequest.Response) *PartnerData4GetOrderStatusResult {
	return &PartnerData4GetOrderStatusResult{Result: result, Body: body, Http: http}
}

// PartnerData4GetOrderStatus 查询秒出票订单状态
// https://docs.apipost.cn/preview/fa101f4865dc783f/66e7c2e894fda4a6?target_id=fce3fbc1-28e1-4757-8665-ffa316a60bfb
func (c *Client) PartnerData4GetOrderStatus(ctx context.Context, notMustParams ...gorequest.Params) (*PartnerData4GetOrderStatusResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/partner/data4/getOrderStatus", params)
	if err != nil {
		return newPartnerData4GetOrderStatusResult(PartnerData4GetOrderStatusResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PartnerData4GetOrderStatusResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPartnerData4GetOrderStatusResult(response, request.ResponseBody, request), err
}
