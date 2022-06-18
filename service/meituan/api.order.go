package meituan

import (
	"encoding/json"
	gorequest2 "go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type ApiOrderResponse struct {
	Status int    `json:"status"`
	Des    string `json:"des"`
	Data   struct {
		BusinessLine     int         `json:"businessLine"`           // 业务线
		SubBusinessLine  int         `json:"subBusinessLine"`        // 子业务线
		ActId            int         `json:"actId"`                  // 活动id，可以在联盟活动列表中查看获取
		Quantity         int         `json:"quantity"`               // 商品数量
		ApiOrderId       string      `json:"ApiOrderId"`             // 订单id
		Paytime          string      `json:"paytime"`                // 订单支付时间，10位时间戳
		ModTime          string      `json:"modTime"`                // 订单信息修改时间，10位时间戳
		Payprice         string      `json:"payprice"`               // 订单用户实际支付金额
		Profit           string      `json:"profit,omitempty"`       // 订单预估返佣金额
		CpaProfit        string      `json:"cpaProfit,omitempty"`    // 订单预估cpa总收益（优选、话费券）
		Sid              string      `json:"sid"`                    // 订单对应的推广位sid
		Appkey           string      `json:"appkey"`                 // 订单对应的appkey，外卖、话费、闪购、优选、酒店订单会返回该字段
		Smstitle         string      `json:"smstitle"`               // 订单标题
		Status           int         `json:"status"`                 // 订单状态，外卖、话费、闪购、优选、酒店订单会返回该字段 1 已付款 8 已完成 9 已退款或风控
		TradeTypeList    []int       `json:"tradeTypeList"`          // 订单的奖励类型 3 首购奖励 5 留存奖励 2 cps 3 首购奖励
		RiskApiOrder     int         `json:"riskApiOrder"`           // 0表示非风控订单，1表示风控订单
		Refundprofit     string      `json:"refundprofit,omitempty"` // 订单需要扣除的返佣金额，外卖、话费、闪购、优选、酒店订单若发生退款会返回该字段
		CpaRefundProfit  interface{} `json:"cpaRefundProfit"`        // 订单需要扣除的cpa返佣金额（优选、话费券）
		RefundInfoList   interface{} `json:"refundInfoList"`         // 退款列表
		RefundProfitList interface{} `json:"refundProfitList"`
		Extra            interface{} `json:"extra"`
	} `json:"data"`
}

type ApiOrderResult struct {
	Result ApiOrderResponse    // 结果
	Body   []byte              // 内容
	Http   gorequest2.Response // 请求
	Err    error               // 错误
}

func NewApiOrderResult(result ApiOrderResponse, body []byte, http gorequest2.Response, err error) *ApiOrderResult {
	return &ApiOrderResult{Result: result, Body: body, Http: http, Err: err}
}

// ApiOrder 单订单查询接口（新版）
// https://union.meituan.com/v2/apiDetail?id=24
func (app *App) ApiOrder(notMustParams ...gorequest2.Params) *ApiOrderResult {
	// 参数
	params := gorequest2.NewParamsWith(notMustParams...)
	// 请求时刻10位时间戳(秒级)，有效期60s
	params["appkey"] = app.appKey
	params["sign"] = app.getSign(app.secret, params)
	// 请求
	request, err := app.request("https://openapi.meituan.com/api/order", params, http.MethodGet)
	// 定义
	var response ApiOrderResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewApiOrderResult(response, request.ResponseBody, request, err)
}
