package wechatpayopen

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type MerchantFundBalanceResponse struct {
	AvailableAmount int64 `json:"available_amount"` // 可用余额
	PendingAmount   int64 `json:"pending_amount"`   // 不可用余额
}

type MerchantFundBalanceResult struct {
	Result MerchantFundBalanceResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
}

func newMerchantFundBalanceResult(result MerchantFundBalanceResponse, body []byte, http gorequest.Response) *MerchantFundBalanceResult {
	return &MerchantFundBalanceResult{Result: result, Body: body, Http: http}
}

// MerchantFundBalance 查询电商平台账户实时余额API
// accountType 账户类型 BASIC：基本账户 OPERATION：运营账户 FEES：手续费账户
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_7_3.shtml
func (c *Client) MerchantFundBalance(ctx context.Context, accountType string, notMustParams ...gorequest.Params) (*MerchantFundBalanceResult, ApiError, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, fmt.Sprintf("v3/merchant/fund/balance/%s", accountType))
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response MerchantFundBalanceResponse
	var apiError ApiError
	request, err := c.request(ctx, fmt.Sprintf("v3/merchant/fund/balance/%s", accountType), params, http.MethodGet, &response, &apiError)
	return newMerchantFundBalanceResult(response, request.ResponseBody, request), apiError, err
}
