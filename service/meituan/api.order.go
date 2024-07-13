package meituan

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type ApiOrderResponse struct {
	Status int    `json:"status"`
	Des    string `json:"des"`
	Data   struct {
		ActId           int    `json:"actId,omitempty"`           // 活动id，可以在联盟活动列表中查看获取
		BusinessLine    int    `json:"businessLine,omitempty"`    // 业务线
		SubBusinessLine int    `json:"subBusinessLine,omitempty"` // 子业务线
		Quantity        int    `json:"quantity,omitempty"`        // 商品数量
		OrderId         string `json:"orderId,omitempty"`         // 订单id
		Paytime         string `json:"paytime,omitempty"`         // 订单支付时间，10位时间戳
		ModTime         string `json:"modTime,omitempty"`         // 订单信息修改时间，10位时间戳
		Payprice        string `json:"payprice,omitempty"`        // 订单用户实际支付金额
		Profit          string `json:"profit,omitempty"`          // 订单预估返佣金额
		CpaProfit       string `json:"cpaProfit,omitempty"`       // 订单预估cpa总收益（优选、话费券）
		Sid             string `json:"sid,omitempty"`             // 订单对应的推广位sid
		Appkey          string `json:"appkey,omitempty"`          // 订单对应的appkey，外卖、话费、闪购、优选、酒店订单会返回该字段
		Smstitle        string `json:"smstitle,omitempty"`        // 订单标题
		Status          int    `json:"status,omitempty"`          // 订单状态，外卖、话费、闪购、优选、酒店订单会返回该字段 1 已付款 8 已完成 9 已退款或风控
		TradeTypeList   []int  `json:"tradeTypeList,omitempty"`   // 订单的奖励类型 3 首购奖励 5 留存奖励 2 cps 3 首购奖励
		RiskApiOrder    int    `json:"riskApiOrder,omitempty"`    // 0表示非风控订单，1表示风控订单
		Refundprofit    string `json:"refundprofit,omitempty"`    // 订单需要扣除的返佣金额，外卖、话费、闪购、优选、酒店订单若发生退款会返回该字段
		CpaRefundProfit string `json:"cpaRefundProfit,omitempty"` // 订单需要扣除的cpa返佣金额（优选、话费券）
		RefundInfoList  struct {
			Id          string `json:"id,omitempty"`
			RefundPrice string `json:"refundPrice,omitempty"`
			RefundTime  string `json:"refundTime,omitempty"`
			RefundType  int    `json:"refundType,omitempty"`
		} `json:"refundInfoList,omitempty"` // 退款列表
		RefundProfitList struct {
			Id               string `json:"id,omitempty"`
			RefundProfit     string `json:"refundProfit,omitempty"`
			RefundFinishTime string `json:"refundFinishTime,omitempty"`
			Type             int    `json:"type,omitempty"`
		} `json:"refundProfitList,omitempty"`
		ConsumeProfitList struct {
			Id                string `json:"id,omitempty"`
			ConsumeProfit     string `json:"consumeProfit,omitempty"`
			ConsumeFinishTime string `json:"consumeFinishTime,omitempty"`
			Type              string `json:"type,omitempty"`
		} `json:"consumeProfitList,omitempty"`
		CouponCode  string `json:"coupon_code,omitempty"` // 券码
		ProductId   string `json:"productId,omitempty"`   // 商品ID
		ProductName string `json:"productName,omitempty"` // 商品名称
	} `json:"data"`
}

type ApiOrderResult struct {
	Result ApiOrderResponse   // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newApiOrderResult(result ApiOrderResponse, body []byte, http gorequest.Response) *ApiOrderResult {
	return &ApiOrderResult{Result: result, Body: body, Http: http}
}

// ApiOrder 单订单查询接口（新版）
// https://union.meituan.com/v2/apiDetail?id=24
func (c *Client) ApiOrder(ctx context.Context, notMustParams ...gorequest.Params) (*ApiOrderResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "api/order")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appkey", c.GetAppKey())
	params.Set("sign", c.getSign(c.GetSecret(), params))

	// 请求
	var response ApiOrderResponse
	request, err := c.request(ctx, "api/order", params, http.MethodGet, &response)
	return newApiOrderResult(response, request.ResponseBody, request), err
}
