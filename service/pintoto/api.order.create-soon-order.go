package pintoto

import (
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
)

type ApiOrderCreateSoonOrder struct {
	ShowId           string `json:"showId"`                  // 排期的showId,由影院接口得来
	Seat             string `json:"seat"`                    // 用户所选的座位，例：1排1座,1排2座 以英文的逗号 “ , “隔开。 如果座位是情侣座，请传入 ： 1排1座(情侣座),1排2座(情侣座)
	ReservedPhone    string `json:"reservedPhone,omitempty"` // 下单时预留的手机号，方便问题沟通
	ThirdOrderId     string `json:"thirdOrderId"`            // 接入方的订单号， 接入方须保证此订单号唯一性
	NotifyUrl        string `json:"notifyUrl"`               // 回调地址，各个场景发生时，将通过此地址通知接入方，详情请看【回调api】
	AcceptChangeSeat int    `json:"acceptChangeSeat"`        // 是否允许调座，1-允许，0-不允许
	SeatId           string `json:"seatId,omitempty"`        // 座位接口的seatId字段， 如果有多个，则以竖线分割
	SeatNo           string `json:"seatNo,omitempty"`        // 座位接口的seatNo字段，如果有多个，则以竖线分割
	NetPrice         int    `json:"netPrice"`                // 所下单所有座位的市场总价单位：分，不可随意乱传，必须是真实价格，如座位有分区定价，也许一一计算后得到总价，否则自动出票失败。由于场次价格延迟问题，有可能造成场次价格和最终价格不一致，此时会出票失败。
	TestType         int    `json:"testType"`                // 仅当为调用测试环境时候，此字段有用， 可模拟秒出票结果。 200 模拟出票成功结果 201 模拟正在出票中结果 500模拟出票失败结果
}

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
	Err    error                           // 错误
}

func newApiOrderCreateSoonOrderResult(result ApiOrderCreateSoonOrderResponse, body []byte, http gorequest.Response, err error) *ApiOrderCreateSoonOrderResult {
	return &ApiOrderCreateSoonOrderResult{Result: result, Body: body, Http: http, Err: err}
}

// ApiOrderCreateSoonOrder 秒出单下单 https://www.showdoc.com.cn/1154868044931571/6437295495912025
func (c *Client) ApiOrderCreateSoonOrder(param ApiOrderCreateSoonOrder) *ApiOrderCreateSoonOrderResult {
	// api params
	params := map[string]interface{}{}
	b, _ := json.Marshal(&param)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		params[k] = v
	}
	request, err := c.request(apiUrl+"/api/order/create-soon-order", params)
	// 定义
	var response ApiOrderCreateSoonOrderResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newApiOrderCreateSoonOrderResult(response, request.ResponseBody, request, err)
}
