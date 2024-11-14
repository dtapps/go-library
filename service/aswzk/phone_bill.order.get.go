package aswzk

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PhoneBillOrderQueryResponse struct {
	Code int    `json:"code"`
	Info string `json:"info"`
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

type PhoneBillOrderQueryResult struct {
	Result PhoneBillOrderQueryResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
}

func newPhoneBillOrderQueryResult(result PhoneBillOrderQueryResponse, body []byte, http gorequest.Response) *PhoneBillOrderQueryResult {
	return &PhoneBillOrderQueryResult{Result: result, Body: body, Http: http}
}

// PhoneBillOrderQuery 话费订单查询
func (c *Client) PhoneBillOrderQuery(ctx context.Context, orderID, orderNo string, notMustParams ...*gorequest.Params) (*PhoneBillOrderQueryResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("order_id", orderID) // 订单编号
	params.Set("order_no", orderNo) // 商户订单编号

	// 请求
	var response PhoneBillOrderQueryResponse
	request, err := c.request(ctx, "phone_bill/order", params, http.MethodGet, &response)
	return newPhoneBillOrderQueryResult(response, request.ResponseBody, request), err
}
