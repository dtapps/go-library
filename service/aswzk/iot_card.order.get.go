package aswzk

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type IotCardOrderGetResponse struct {
	Code int    `json:"code"` // 状态码
	Info string `json:"info"` // 状态信息
	Data struct {
		RechargeIccid  string  `json:"recharge_iccid"`            // 充值ICCID
		PackageCode    float64 `json:"package_code"`              // 套餐编号
		RechargeReason string  `json:"recharge_reason,omitempty"` // 充值失败原因
		OrderID        string  `json:"order_id"`                  // 订单编号
		OrderNo        string  `json:"order_no"`                  // 商户订单编号
		Remark         string  `json:"remark"`                    // 订单备注
		OrderStatus    string  `json:"order_status"`              // 订单状态
		OrderCost      float64 `json:"order_cost,omitempty"`      // 订单成本价
	} `json:"data,omitempty"`
	Time    int    `json:"time"`
	TraceId string `json:"trace_id"`
}

type IotCardOrderGetResult struct {
	Result IotCardOrderGetResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
}

func newIotCardOrderGetResult(result IotCardOrderGetResponse, body []byte, http gorequest.Response) *IotCardOrderGetResult {
	return &IotCardOrderGetResult{Result: result, Body: body, Http: http}
}

// IotCardOrderGet 物联卡订单查询
func (c *Client) IotCardOrderGet(ctx context.Context, notMustParams ...gorequest.Params) (*IotCardOrderGetResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response IotCardOrderGetResponse
	request, err := c.request(ctx, "iot_card/order", params, http.MethodGet, &response)
	return newIotCardOrderGetResult(response, request.ResponseBody, request), err
}
