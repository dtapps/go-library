package wechatpayopen

import (
	"context"
	"fmt"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type EcommerceFundBalanceResponse struct {
	SubMchid        string `json:"sub_mchid"`        // 二级商户号
	AccountType     string `json:"account_type"`     // 账户类型
	AvailableAmount int64  `json:"available_amount"` // 可用余额
	PendingAmount   int64  `json:"pending_amount"`   // 不可用余额
}

// EcommerceFundBalance 查询二级商户账户实时余额API
// accountType 账户类型 BASIC：基本账户 OPERATION：运营账户 FEES：手续费账户
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_7_1.shtml
func (c *Client) EcommerceFundBalance(ctx context.Context, accountType string, notMustParams ...*gorequest.Params) (response EcommerceFundBalanceResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, fmt.Sprintf("v3/ecommerce/fund/balance/%s?account_type=%s", c.GetSubMchId(), accountType), params, http.MethodGet, &response, nil)
	return
}
