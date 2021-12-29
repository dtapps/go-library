package dingdanxia

import (
	"encoding/json"
	"net/http"
)

// WaimaiMeituanOrdersResult 返回参数
type WaimaiMeituanOrdersResult struct {
	Code         int    `json:"code"`
	Msg          string `json:"msg"`           // 描述
	TotalResults int    `json:"total_results"` // 总条数
	Data         []struct {
		Orderid     string `json:"orderid"`      // 订单ID
		Paytime     string `json:"paytime"`      // 订单支付时间
		Payprice    string `json:"payprice"`     // 订单支付金额
		Profit      string `json:"profit"`       // 订单返佣金额
		Smstitle    string `json:"smstitle"`     // 订单标题
		Sid         string `json:"sid"`          // 渠道方用户唯一标识
		Quantity    string `json:"quantity"`     // 退款笔数
		Refundtime  string `json:"refundtime"`   // 退款时间
		Money       string `json:"money"`        // 退款金额
		RefundMoney string `json:"refund_money"` // 退佣金额
		CreateTime  string `json:"create_time"`  // 数据入库更新时间（订单状态改变，该时间会变）
		Status      int    `json:"status"`       // 订单状态(1-已提交（已付款）、8-已完成（确认收货）、9-已退款)
		Type        int    `json:"type"`         // 订单类型（活动名称）4-外卖 6-闪购 8-优选 2-酒店
	} `json:"data"`
}

// NewWaimaiMeituanOrdersResult 构造函数
func NewWaimaiMeituanOrdersResult(result WaimaiMeituanOrdersResult, byte []byte, err error) *Result {
	return &Result{WaimaiMeituanOrdersResult: result, Byte: byte, Err: err}
}

// WaimaiMeituanOrders 美团联盟外卖/闪购/优选/酒店订单查询API
// https://www.dingdanxia.com/doc/176/173
func (app *App) WaimaiMeituanOrders(notMustParams ...Params) *Result {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request("https://api.tbk.dingdanxia.com/waimai/meituan_orders", params, http.MethodPost)
	// 定义
	var response WaimaiMeituanOrdersResult
	err = json.Unmarshal(body, &response)
	return NewWaimaiMeituanOrdersResult(response, body, err)
}
