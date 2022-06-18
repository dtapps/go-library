package dingdanxia

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WaiMaiMeituanOrdersResponse struct {
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

type WaiMaiMeituanOrdersResult struct {
	Result WaiMaiMeituanOrdersResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
	Err    error                       // 错误
}

func NewWaiMaiMeituanOrdersResult(result WaiMaiMeituanOrdersResponse, body []byte, http gorequest.Response, err error) *WaiMaiMeituanOrdersResult {
	return &WaiMaiMeituanOrdersResult{Result: result, Body: body, Http: http, Err: err}
}

// WaiMaiMeituanOrders 美团联盟外卖/闪购/优选/酒店订单查询API
// https://www.dingdanxia.com/doc/176/173
func (c *Client) WaiMaiMeituanOrders(notMustParams ...gorequest.Params) *WaiMaiMeituanOrdersResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(apiUrl+"/waimai/meituan_orders", params, http.MethodPost)
	// 定义
	var response WaiMaiMeituanOrdersResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewWaiMaiMeituanOrdersResult(response, request.ResponseBody, request, err)
}
