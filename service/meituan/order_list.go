package meituan

import (
	"encoding/json"
	"gopkg.in/dtapps/go-library.v3/utils/gotime"
	"net/http"
)

type OrderListResponse struct {
	DataList []struct {
		Orderid                     string `json:"orderid"`         // 订单id
		Paytime                     string `json:"paytime"`         // 订单支付时间，10位时间戳
		Payprice                    string `json:"payprice"`        // 订单用户实际支付金额
		Sid                         string `json:"sid"`             // 订单对应的推广位sid
		Smstitle                    string `json:"smstitle"`        // 订单标题
		Appkey                      string `json:"appkey"`          // 订单对应的appkey，外卖、话费、闪购、优选订单会返回该字段
		Status                      int    `json:"status"`          // 订单状态，外卖、话费、闪购、优选、酒店订单会返回该字段 1 已付款 8 已完成 9 已退款或风控
		Profit                      string `json:"profit"`          // 订单预估返佣金额
		CpaProfit                   string `json:"cpaProfit"`       // 订单预估cpa总收益（优选、话费券）
		Refundtime                  string `json:"refundtime"`      // 订单退款时间，10位时间戳，外卖、话费、闪购、优选、酒店订单若发生退款会返回该字段(退款时间为最近的一次退款)
		Refundprice                 string `json:"refundprice"`     // 订单实际退款金额，外卖、话费、闪购、优选、酒店订单若发生退款会返回该字段
		Refundprofit                string `json:"refundprofit"`    // 订单需要扣除的返佣金额，外卖、话费、闪购、优选、酒店订单若发生退款会返回该字段
		CpaRefundProfit             string `json:"cpaRefundProfit"` // 订单需要扣除的cpa返佣金额（优选、话费券）
		Extra                       string `json:"extra"`
		TradeTypeList               []int  `json:"tradeTypeList"` // 订单的奖励类型 3 首购奖励 5 留存奖励 2 cps 3 首购奖励
		TradeTypeBusinessTypeMapStr string `json:"tradeTypeBusinessTypeMapStr"`
		RiskOrder                   int    `json:"riskOrder"`       // 0表示非风控订单，1表示风控订单
		BusinessLine                int    `json:"businessLine"`    // 业务线
		SubBusinessLine             int    `json:"subBusinessLine"` // 子业务线
		ActId                       int    `json:"actId"`           // 活动id，可以在联盟活动列表中查看获取
	} `json:"dataList"`
	Total int `json:"total"` // 查询条件命中的总数据条数，用于计算分页参数
}

type OrderListResult struct {
	Result OrderListResponse // 结果
	Body   []byte            // 内容
	Err    error             // 错误
}

func NewOrderListResult(result OrderListResponse, body []byte, err error) *OrderListResult {
	return &OrderListResult{Result: result, Body: body, Err: err}
}

// OrderList 订单列表查询接口（新版）
// https://union.meituan.com/v2/apiDetail?id=23
func (app *App) OrderList(notMustParams ...Params) *OrderListResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求时刻10位时间戳(秒级)，有效期60s
	params["ts"] = gotime.Current().Timestamp()
	params["appkey"] = app.AppKey
	params["sign"] = app.getSign(app.Secret, params)
	// 请求
	body, err := app.request("https://openapi.meituan.com/api/orderList", params, http.MethodGet)
	// 定义
	var response OrderListResponse
	err = json.Unmarshal(body, &response)
	return NewOrderListResult(response, body, err)
}
