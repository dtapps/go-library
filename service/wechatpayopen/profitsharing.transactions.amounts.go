package wechatpayopen

import (
	"context"
	"fmt"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type ProfitSharingTransactionsAmountsResponse struct {
	TransactionId string `json:"transaction_id"` // 微信订单号
	UnsplitAmount int    `json:"unsplit_amount"` // 订单剩余待分金额
}

// ProfitSharingTransactionsAmounts 查询剩余待分金额API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_6.shtml
func (c *Client) ProfitSharingTransactionsAmounts(ctx context.Context, transactionId string, notMustParams ...*gorequest.Params) (response ProfitSharingTransactionsAmountsResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, fmt.Sprintf("v3/profitsharing/transactions/%s", transactionId), params, http.MethodGet, &response, &apiError)
	return
}
