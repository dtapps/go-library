package dingdanxia

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WaiMaiMeituanOrderIdResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Order struct {
			OrderId         string      `json:"orderId"`         // 订单号
			Sid             string      `json:"sid"`             // 合作方唯一标识
			Payprice        string      `json:"payprice"`        // 订单用户实际支付金额
			Profit          string      `json:"profit"`          // 订单预估返佣金额
			CpaProfit       string      `json:"cpaProfit"`       // 订单预估cpa总收益（优选、话费券）
			Smstitle        string      `json:"smstitle"`        // 订单标题
			Status          int         `json:"status"`          // 订单状态，外卖、话费、闪购、优选、酒店订单会返回该字段 1 已付款 8 已完成 9 已退款或风控
			TradeTypeList   []int       `json:"tradeTypeList"`   // 订单的奖励类型 话费订单类型返回该字段 3 首购奖励 5 留存奖励 优选订单类型返回该字段 2 cps 3 首购奖励
			RiskOrder       interface{} `json:"riskOrder"`       // 0表示非风控订单，1表示风控订单
			Refundprofit    interface{} `json:"refundprofit"`    // 订单需要扣除的返佣金额，外卖、话费、闪购、优选、酒店订单若发生退款会返回该字段
			CpaRefundProfit interface{} `json:"cpaRefundProfit"` // 订单需要扣除的cpa返佣金额（优选、话费券）
			RefundInfoList  struct {
				RefundPrice interface{} `json:"refundPrice"` // 订单实际退款金额，外卖、话费、闪购、优选、酒店订单若发生退款会返回该字段
				RefundTime  interface{} `json:"refundTime"`  // 订单退款时间，10位时间戳
				RefundType  interface{} `json:"refundType"`  // ALL_REFUND(1, "全部退"), PART_REFUND(2, "部分退"), RISK_REFUND(3, "风控退");
			} `json:"refundInfoList,omitempty"` // 退款列表
			RefundProfitList struct { // 退款佣金明细
				RefundProfit     interface{} `json:"refundProfit"`     // 退款佣金
				RefundFinishTime interface{} `json:"refundFinishTime"` // 佣金产生时间,10位时间戳
				Type             interface{} `json:"type"`             // 券订单: 1 流量订单: 2 首单: 3 复购: 4 留存: 5 二单: 6 唤起: 7
			} `json:"refundProfitList,omitempty"`
			BusinessLine    int         `json:"businessLine"`
			SubBusinessLine int         `json:"subBusinessLine"`
			ActId           int         `json:"actId"`
			Quantity        int         `json:"quantity"`
			Paytime         string      `json:"paytime"`
			ModTime         string      `json:"modTime"`
			Appkey          string      `json:"appkey"`
			Extra           interface{} `json:"extra"`
		} `json:"order"`
	} `json:"data"`
}

type WaiMaiMeituanOrderIdResult struct {
	Result WaiMaiMeituanOrderIdResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
}

func newWaiMaiMeituanOrderIdResult(result WaiMaiMeituanOrderIdResponse, body []byte, http gorequest.Response) *WaiMaiMeituanOrderIdResult {
	return &WaiMaiMeituanOrderIdResult{Result: result, Body: body, Http: http}
}

// WaiMaiMeituanOrderId 美团联盟外卖/闪购/优选/酒店订单查询API（订单号版）
// https://www.dingdanxia.com/doc/179/173
func (c *Client) WaiMaiMeituanOrderId(ctx context.Context, orderId string, Type int, notMustParams ...*gorequest.Params) (*WaiMaiMeituanOrderIdResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("orderid", orderId) // 订单号
	if Type <= 0 {
		Type = 4
	}
	params.Set("type", Type) // 2-酒店 4-外卖 6-闪购 8-优选 默认4
	// 请求
	request, err := c.request(ctx, apiUrl+"/waimai/meituan_orderid", params, http.MethodPost)
	if err != nil {
		return newWaiMaiMeituanOrderIdResult(WaiMaiMeituanOrderIdResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WaiMaiMeituanOrderIdResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWaiMaiMeituanOrderIdResult(response, request.ResponseBody, request), err
}
