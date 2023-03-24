package wechatpayopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type ProfitSharingTransactionsAmountsResponse struct {
	TransactionId string `json:"transaction_id"` // 微信订单号
	UnsplitAmount int    `json:"unsplit_amount"` // 订单剩余待分金额
}

type ProfitSharingTransactionsAmountsResult struct {
	Result   ProfitSharingTransactionsAmountsResponse // 结果
	Body     []byte                                   // 内容
	Http     gorequest.Response                       // 请求
	Err      error                                    // 错误
	ApiError ApiError                                 // 接口错误
}

func newProfitSharingTransactionsAmountsResult(result ProfitSharingTransactionsAmountsResponse, body []byte, http gorequest.Response, err error, apiError ApiError) *ProfitSharingTransactionsAmountsResult {
	return &ProfitSharingTransactionsAmountsResult{Result: result, Body: body, Http: http, Err: err, ApiError: apiError}
}

// ProfitSharingTransactionsAmounts 查询剩余待分金额API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_6.shtml
func (c *Client) ProfitSharingTransactionsAmounts(ctx context.Context, transactionId string) *ProfitSharingTransactionsAmountsResult {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/profitsharing/transactions/"+transactionId, params, http.MethodGet)
	if err != nil {
		return newProfitSharingTransactionsAmountsResult(ProfitSharingTransactionsAmountsResponse{}, request.ResponseBody, request, err, ApiError{})
	}
	// 定义
	var response ProfitSharingTransactionsAmountsResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = gojson.Unmarshal(request.ResponseBody, &apiError)
	return newProfitSharingTransactionsAmountsResult(response, request.ResponseBody, request, err, apiError)
}
