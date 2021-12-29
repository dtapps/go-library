package dingdanxia

import (
	"encoding/json"
	"net/http"
)

// JdJyOrderDetailsResponse 返回参数
type JdJyOrderDetailsResponse struct {
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

// NewJdJyOrderDetailsResult 构造函数
func NewJdJyOrderDetailsResult(jdJyOrderDetailsResponse JdJyOrderDetailsResponse, byte []byte, err error) *Result {
	return &Result{JdJyOrderDetailsResponse: jdJyOrderDetailsResponse, Byte: byte, Err: err}
}

// JdJyOrderDetails 【官方不维护】 京佣订单
func (app *App) JdJyOrderDetails(notMustParams ...Params) *Result {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request("https://api.tbk.dingdanxia.com/jd/jy_order_details", params, http.MethodPost)
	// 定义
	var response JdJyOrderDetailsResponse
	err = json.Unmarshal(body, &response)
	return NewJdJyOrderDetailsResult(response, body, err)
}
