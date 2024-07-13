package wechatpayopen

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type EcommerceFundEndDayBalanceResponse struct {
	SubMchid        string `json:"sub_mchid"`        // 二级商户号
	AvailableAmount int64  `json:"available_amount"` // 可用余额
	PendingAmount   int64  `json:"pending_amount"`   // 不可用余额
}

type EcommerceFundEndDayBalanceResult struct {
	Result EcommerceFundEndDayBalanceResponse // 结果
	Body   []byte                             // 内容
	Http   gorequest.Response                 // 请求
}

func newEcommerceFundEndDayBalanceResult(result EcommerceFundEndDayBalanceResponse, body []byte, http gorequest.Response) *EcommerceFundEndDayBalanceResult {
	return &EcommerceFundEndDayBalanceResult{Result: result, Body: body, Http: http}
}

// EcommerceFundEndDayBalance 查询二级商户账户日终余额API
// date 日期 示例值：2019-08-17
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_7_2.shtml
func (c *Client) EcommerceFundEndDayBalance(ctx context.Context, date string, notMustParams ...gorequest.Params) (*EcommerceFundEndDayBalanceResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, fmt.Sprintf("v3/ecommerce/fund/enddaybalance/%s?date=%s", c.GetSubMchId(), date))
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response EcommerceFundEndDayBalanceResponse
	request, err := c.request(ctx, fmt.Sprintf("v3/ecommerce/fund/enddaybalance/%s?date=%s", c.GetSubMchId(), date), params, http.MethodGet, &response, nil)
	return newEcommerceFundEndDayBalanceResult(response, request.ResponseBody, request), err
}
