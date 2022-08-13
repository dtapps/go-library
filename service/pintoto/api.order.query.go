package pintoto

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
)

type ApiOrderQueryResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		AppKey          string      `json:"appKey"`          // 下单appKey
		ThirdOrderId    string      `json:"thirdOrderId"`    // 接入方的订单号
		OrderStatus     int         `json:"orderStatus"`     // 订单状态：2-受理中，3-待出票，4-已出票待结算，5-已结算，10-订单关闭
		OrderStatusStr  string      `json:"orderStatusStr"`  // 订单状态说明
		InitPrice       int         `json:"initPrice"`       // 订单市场价：分
		OrderPrice      int         `json:"orderPrice"`      // 订单成本价：分，接入方可拿次字段作为下单成本价
		Seat            string      `json:"seat"`            // 座位：英文逗号隔开
		OrderNum        int         `json:"orderNum"`        // 座位数
		ReservedPhone   string      `json:"reservedPhone"`   // 下单预留手机号码
		CreateTime      string      `json:"createTime"`      // 下单时间
		ReadyTicketTime string      `json:"readyTicketTime"` // 待出票时间
		TicketTime      string      `json:"ticketTime"`      // 出票时间
		NotifyUrl       string      `json:"notifyUrl"`       // 回调通知地址
		CloseTime       interface{} `json:"closeTime"`       // 关闭时间
		CloseCause      interface{} `json:"closeCause"`      // 关闭原因
		TicketCode      []struct {
			Code string `json:"code"`
			Type int    `json:"type"`
			Url  string `json:"url"`
		} `json:"ticketCode"` // 取票码，type为1时，为字符串，type为2时，为取票码原始截图。 理论上一个取票码包含各字符串和原始截图， 原始截图可能不和字符串同步返回，有滞后性。
	} `json:"data"`
	Code int `json:"code"`
}

type ApiOrderQueryResult struct {
	Result ApiOrderQueryResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
	Err    error                 // 错误
}

func newApiOrderQueryResult(result ApiOrderQueryResponse, body []byte, http gorequest.Response, err error) *ApiOrderQueryResult {
	return &ApiOrderQueryResult{Result: result, Body: body, Http: http, Err: err}
}

// ApiOrderQuery 订单查询 https://www.showdoc.com.cn/1154868044931571/5965244588489845
func (c *Client) ApiOrderQuery(ctx context.Context, thirdOrderId string) *ApiOrderQueryResult {
	// 测试
	param := gorequest.NewParams()
	param.Set("thirdOrderId", thirdOrderId)
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(ctx, apiUrl+"/api/order/query", params)
	// 定义
	var response ApiOrderQueryResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newApiOrderQueryResult(response, request.ResponseBody, request, err)
}
