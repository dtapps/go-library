package wechatpayopen

import (
	"context"
	"fmt"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type MerchantFundDayEndBalanceResponse struct {
	AvailableAmount int64 `json:"available_amount"` // 可用余额
	PendingAmount   int64 `json:"pending_amount"`   // 不可用余额
}

// MerchantFundDayEndBalance 查询电商平台账户日终余额API
// accountType 账户类型 BASIC：基本账户 OPERATION：运营账户 FEES：手续费账户
// date 日期 示例值：2019-08-17
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_7_4.shtml
func (c *Client) MerchantFundDayEndBalance(ctx context.Context, accountType, date string, notMustParams ...*gorequest.Params) (response MerchantFundDayEndBalanceResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, fmt.Sprintf("/v3/merchant/fund/dayendbalance/%s?date=%s", accountType, date), params, http.MethodGet, &response, &apiError)
	return
}
