package wechatpayopen

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type ProfitSharingTransactionsAmountsResponse struct {
	TransactionId string `json:"transaction_id"` // 微信订单号
	UnsplitAmount int    `json:"unsplit_amount"` // 订单剩余待分金额
}

type ProfitSharingTransactionsAmountsResult struct {
	Result ProfitSharingTransactionsAmountsResponse // 结果
	Body   []byte                                   // 内容
	Http   gorequest.Response                       // 请求
}

func newProfitSharingTransactionsAmountsResult(result ProfitSharingTransactionsAmountsResponse, body []byte, http gorequest.Response) *ProfitSharingTransactionsAmountsResult {
	return &ProfitSharingTransactionsAmountsResult{Result: result, Body: body, Http: http}
}

// ProfitSharingTransactionsAmounts 查询剩余待分金额API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_6.shtml
func (c *Client) ProfitSharingTransactionsAmounts(ctx context.Context, transactionId string, notMustParams ...gorequest.Params) (*ProfitSharingTransactionsAmountsResult, ApiError, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response ProfitSharingTransactionsAmountsResponse
	var apiError ApiError
	request, err := c.request(ctx, fmt.Sprintf("v3/profitsharing/transactions/%s", transactionId), params, http.MethodGet, &response, &apiError)
	return newProfitSharingTransactionsAmountsResult(response, request.ResponseBody, request), apiError, err
}
