package aswzk

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type ElectricityBillOrderQueryResponse struct {
	Code int    `json:"code"` // 状态码
	Info string `json:"info"` // 状态信息
	Data struct {
		RechargeAccount string  `json:"recharge_account"`          // 充值账号
		RechargeMoney   float64 `json:"recharge_money"`            // 充值金额
		RechargeType    string  `json:"recharge_type"`             // 充值类型
		RechargeReason  string  `json:"recharge_reason,omitempty"` // 充值失败原因
		OrderID         string  `json:"order_id"`                  // 订单编号
		OrderNo         string  `json:"order_no"`                  // 商户订单编号
		Remark          string  `json:"remark"`                    // 订单备注
		OrderStatus     string  `json:"order_status"`              // 订单状态
		OrderCost       float64 `json:"order_cost,omitempty"`      // 订单成本价
	} `json:"data,omitempty"`
	Time    int    `json:"time"`
	TraceId string `json:"trace_id"`
}

type ElectricityBillOrderQueryResult struct {
	Result ElectricityBillOrderQueryResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
}

func newElectricityBillOrderQueryResult(result ElectricityBillOrderQueryResponse, body []byte, http gorequest.Response) *ElectricityBillOrderQueryResult {
	return &ElectricityBillOrderQueryResult{Result: result, Body: body, Http: http}
}

// ElectricityBillOrderQuery 电费订单查询
func (c *Client) ElectricityBillOrderQuery(ctx context.Context, orderID string, orderNo string, notMustParams ...gorequest.Params) (*ElectricityBillOrderQueryResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "electricity_bill/order")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("order_id", orderID) // 订单编号
	params.Set("order_no", orderNo) // 商户订单编号

	// 请求
	var response ElectricityBillOrderQueryResponse
	request, err := c.request(ctx, "electricity_bill/order", params, http.MethodGet, &response)
	return newElectricityBillOrderQueryResult(response, request.ResponseBody, request), err
}
