package aswzk

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type ElectricityBillOrderQueryResponse struct {
	Code int    `json:"code"` // 状态码
	Info string `json:"info"` // 状态信息
	Data struct {
		RechargeAccount string  `json:"recharge_account"` // 充值账号
		RechargeMoney   float64 `json:"recharge_money"`   // 充值金额
		RechargeType    string  `json:"recharge_type"`    // 充值类型
		OrderID         string  `json:"order_id"`         // 订单编号
		OrderNo         string  `json:"order_no"`         // 商户订单编号
		Remark          string  `json:"remark"`           // 订单备注
		OrderStatus     string  `json:"order_status"`     // 订单状态
		OrderCost       float64 `json:"order_cost"`       // 订单成本价
		Reason          string  `json:"reason,omitempty"` // 失败原因
	} `json:"data"`
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
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("order_id", orderID) // 订单编号
	params.Set("order_no", orderNo) // 商户订单编号
	// 请求
	request, err := c.request(ctx, apiUrl+"/electricity_bill/order", params, http.MethodGet)
	if err != nil {
		return newElectricityBillOrderQueryResult(ElectricityBillOrderQueryResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ElectricityBillOrderQueryResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newElectricityBillOrderQueryResult(response, request.ResponseBody, request), err
}
