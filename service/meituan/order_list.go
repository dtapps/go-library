package meituan

import (
	"encoding/json"
	"gopkg.in/dtapps/go-library.v2/utils/gotime"
)

// OrderList 请求参数
type OrderList struct {
	Type          string `json:"type"`                    // 查询订单类型 0 团购订单 2 酒店订单 4 外卖订单 5 话费&团好货订单 6 闪购订单 8 优选订单
	StartTime     string `json:"startTime"`               // 查询起始时间10位时间戳，以下单时间为准
	EndTime       string `json:"endTime"`                 // 查询截止时间10位时间戳，以下单时间为准
	Page          string `json:"page"`                    // 分页参数，起始值从1开始
	Limit         string `json:"limit"`                   // 每页显示数据条数，最大值为100
	QueryTimeType string `json:"queryTimeType,omitempty"` // 查询时间类型，枚举值 1 按订单支付时间查询 2 按订单发生修改时间查询
}

// OrderListResult 返回参数
type OrderListResult struct {
	DataList []struct {
		Orderid                     string        `json:"orderid"`                     // 订单id
		Paytime                     string        `json:"paytime"`                     // 订单支付时间，10位时间戳
		Payprice                    string        `json:"payprice"`                    // 订单用户实际支付金额
		Profit                      string        `json:"profit"`                      // 订单预估返佣金额
		CpaProfit                   string        `json:"cpaProfit"`                   // 订单预估cpa总收益（优选、话费券）
		Sid                         string        `json:"sid"`                         // 订单对应的推广位sid
		Appkey                      string        `json:"appkey,omitempty"`            // 订单对应的appkey，外卖、话费、闪购、优选订单会返回该字段
		Smstitle                    string        `json:"smstitle"`                    // 订单标题
		Refundprice                 string        `json:"refundprice,omitempty"`       // 订单实际退款金额，外卖、话费、闪购、优选、酒店订单若发生退款会返回该字段
		Refundtime                  string        `json:"refundtime,omitempty"`        // 订单退款时间，10位时间戳，外卖、话费、闪购、优选、酒店订单若发生退款会返回该字段
		Refundprofit                string        `json:"refundprofit,omitempty"`      // 订单需要扣除的返佣金额，外卖、话费、闪购、优选、酒店订单若发生退款会返回该字段
		CpaRefundProfit             string        `json:"cpaRefundProfit"`             // 订单需要扣除的cpa返佣金额（优选、话费券）
		Status                      string        `json:"status,omitempty"`            // 订单状态，外卖、话费、闪购、优选、酒店订单会返回该字段 1 已付款 8 已完成 9 已退款或风控
		TradeTypeList               []interface{} `json:"tradeTypeList,omitempty"`     // 订单的奖励类型 话费订单类型返回该字段 3 首购奖励 5 留存奖励 优选订单类型返回该字段 2 cps 3 首购奖励
		TradeTypeBusinessTypeMapStr string        `json:"tradeTypeBusinessTypeMapStr"` // 奖励类型对应平台类型的映射 格式：{3:[3,5]} value的枚举值：1 外卖 2 分销酒店 3 平台 4 券类型酒店 5 团好货 6 优选
		RiskOrder                   string        `json:"riskOrder,omitempty"`         // 0表示正常退款，1表示风控退款，订单状态为退款时有效
	} `json:"dataList"` // 订单列表
	Total int `json:"total"` // 查询条件命中的总数据条数，用于计算分页参数
}

// OrderList 订单列表查询(新) https://union.meituan.com/v2/apiDetail?id=1
func (app *App) OrderList(param OrderList) (result OrderListResult, err error) {
	// 处理默认数据
	if param.Page == "" {
		param.Page = "1"
	}
	if param.Limit == "" {
		param.Limit = "100"
	}
	// 接口参数
	params := map[string]interface{}{}
	b, _ := json.Marshal(&param)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		params[k] = v
	}
	// 请求时刻10位时间戳(秒级)，有效期60s
	params["ts"] = gotime.Current().Timestamp()
	params["key"] = app.AppKey
	params["sign"] = app.getSign(app.Secret, params)
	body, err := app.request("https://runion.meituan.com/api/orderList", params, "GET")

	if err != nil {
		return
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
