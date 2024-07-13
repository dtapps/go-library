package wikeyun

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type RestRechargeQueryResponse struct {
	Code string `json:"code"`
	Data struct {
		Id             int    `json:"id,omitempty"`
		Fanli          string `json:"fanli"`            // 平台返利金额
		Amount         int64  `json:"amount"`           // 充值金额
		CostPrice      string `json:"cost_price"`       // 充值成本价格
		Status         int    `json:"status"`           // 订单状态 0 待支付 1 已付 充值中 2充值成功 3充值失败 需要退款 4退款成功 5已超时 6待充值 7 已匹配 8已存单 9 已取消 10返销 11部分到账 12取消中
		OrderNumber    string `json:"order_number"`     // 平台订单号
		OrderNo        string `json:"order_no"`         // 第三方单号
		OrgOrderNumber string `json:"org_order_number"` // 组织订单号
		StoreId        int    `json:"store_id"`         // 店铺ID
		Mobile         string `json:"mobile"`           // 充值手机号
		ArrivedAmount  string `json:"arrived_amount"`   // 到账金额
		Reason         string `json:"reason,omitempty"` // 失败原因
	} `json:"data"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
}

type RestRechargeQueryResult struct {
	Result RestRechargeQueryResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newRestRechargeQueryResult(result RestRechargeQueryResponse, body []byte, http gorequest.Response) *RestRechargeQueryResult {
	return &RestRechargeQueryResult{Result: result, Body: body, Http: http}
}

// RestRechargeQuery 话费订单查询
// order_number = 平台单号，平台商户单号二选一
// order_no = 商户单号
// https://open.wikeyun.cn/#/apiDocument/9/document/299
func (c *Client) RestRechargeQuery(ctx context.Context, notMustParams ...gorequest.Params) (*RestRechargeQueryResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "rest/Recharge/query")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.GetStoreId()) // 店铺ID

	// 请求
	var response RestRechargeQueryResponse
	request, err := c.request(ctx, "rest/Recharge/query", params, &response)
	return newRestRechargeQueryResult(response, request.ResponseBody, request), err
}

func (resp RestRechargeQueryResponse) GetStatusDesc(status int) string {
	switch status {
	case 1:
		return "充值中"
	case 2:
		return "充值成功"
	case 3:
		return "充值失败"
	case 4:
		return "退款成功"
	case 5:
		return "已超时"
	case 6:
		return "待充值"
	case 7:
		return "已匹配"
	case 8:
		return "已存单"
	case 9:
		return "已取消"
	case 10:
		return "返销"
	case 11:
		return "部分到账"
	case 12:
		return "取消中"
	}
	return "待支付"
}
