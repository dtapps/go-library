package meituan

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
	"net/http"
)

type ApiOrderListResponse struct {
	DataList []struct {
		ActId                       int    `json:"actId,omitempty"`           // 活动id，可以在联盟活动列表中查看获取
		BusinessLine                int    `json:"businessLine,omitempty"`    // 业务线
		SubBusinessLine             int    `json:"subBusinessLine,omitempty"` // 子业务线
		Orderid                     string `json:"orderid,omitempty"`         // 订单id
		Paytime                     string `json:"paytime,omitempty"`         // 订单支付时间，10位时间戳
		Payprice                    string `json:"payprice,omitempty"`        // 订单用户实际支付金额
		Profit                      string `json:"profit,omitempty"`          // 订单预估返佣金额
		CpaProfit                   string `json:"cpaProfit,omitempty"`       // 订单预估cpa总收益（优选、话费券）
		Sid                         string `json:"sid,omitempty"`             // 订单对应的推广位sid
		Appkey                      string `json:"appkey,omitempty"`          // 订单对应的appkey，外卖、话费、闪购、优选订单会返回该字段
		Smstitle                    string `json:"smstitle,omitempty"`        // 订单标题
		ProductId                   string `json:"productId,omitempty"`       // 商品ID
		ProductName                 string `json:"productName,omitempty"`     // 商品名称
		Refundprice                 string `json:"refundprice,omitempty"`     // 订单实际退款金额，外卖、话费、闪购、优选、酒店订单若发生退款会返回该字段
		Refundtime                  string `json:"refundtime,omitempty"`      // 订单退款时间，10位时间戳，外卖、话费、闪购、优选、酒店订单若发生退款会返回该字段(退款时间为最近的一次退款)
		Refundprofit                string `json:"refundprofit,omitempty"`    // 订单需要扣除的返佣金额，外卖、话费、闪购、优选、酒店订单若发生退款会返回该字段
		CpaRefundProfit             string `json:"cpaRefundProfit,omitempty"` // 订单需要扣除的cpa返佣金额（优选、话费券）
		Status                      int    `json:"status,omitempty"`          // 订单状态，外卖、话费、闪购、优选、酒店订单会返回该字段 1 已付款 8 已完成 9 已退款或风控
		TradeTypeList               []int  `json:"tradeTypeList,omitempty"`   // 订单的奖励类型 3 首购奖励 5 留存奖励 2 cps 3 首购奖励
		RiskOrder                   int    `json:"riskOrder,omitempty"`       // 0表示非风控订单，1表示风控订单
		Extra                       string `json:"extra,omitempty"`
		TradeTypeBusinessTypeMapStr string `json:"tradeTypeBusinessTypeMapStr,omitempty"`
	} `json:"dataList"`
	Total int `json:"total"` // 查询条件命中的总数据条数，用于计算分页参数
}

type ApiOrderListResult struct {
	Result ApiOrderListResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
}

func newApiOrderListResult(result ApiOrderListResponse, body []byte, http gorequest.Response) *ApiOrderListResult {
	return &ApiOrderListResult{Result: result, Body: body, Http: http}
}

// ApiOrderList 订单列表查询接口（新版）
// https://union.meituan.com/v2/apiDetail?id=23
func (c *Client) ApiOrderList(ctx context.Context, notMustParams ...*gorequest.Params) (*ApiOrderListResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求时刻10位时间戳(秒级)，有效期60s
	params.Set("ts", gotime.Current().Timestamp())
	params.Set("appkey", c.GetAppKey())
	params.Set("sign", c.getSign(c.GetSecret(), params))

	// 请求
	var response ApiOrderListResponse
	request, err := c.request(ctx, "api/orderList", params, http.MethodGet, &response)
	return newApiOrderListResult(response, request.ResponseBody, request), err
}
