package wechatpayopen

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type ProfitSharingOrdersUnfreezeResponse struct {
	SubMchid      string `json:"sub_mchid"`      // 子商户号
	TransactionId string `json:"transaction_id"` // 微信订单号
	OutOrderNo    string `json:"out_order_no"`   // 商户分账单号
	OrderId       string `json:"order_id"`       // 微信分账单号
	State         string `json:"state"`          // 分账单状态
	Receivers     []struct {
		Amount      int    `json:"amount"`      // 分账金额
		Description string `json:"description"` // 分账描述
		Type        string `json:"type"`        // 分账接收方类型
		Account     string `json:"account"`     // 分账接收方账号
		Result      string `json:"result"`      // 分账结果
		FailReason  string `json:"fail_reason"` // 分账失败原因
		DetailId    string `json:"detail_id"`   // 分账明细单号
		CreateTime  string `json:"create_time"` // 分账创建时间
		FinishTime  string `json:"finish_time"` // 分账完成时间
	} `json:"receivers,omitempty"` // 分账接收方列表
}

type ProfitSharingOrdersUnfreezeResult struct {
	Result   ProfitSharingOrdersUnfreezeResponse // 结果
	Body     []byte                              // 内容
	Http     gorequest.Response                  // 请求
	Err      error                               // 错误
	ApiError ApiError                            // 接口错误
}

func newProfitSharingOrdersUnfreezeResult(result ProfitSharingOrdersUnfreezeResponse, body []byte, http gorequest.Response, err error, apiError ApiError) *ProfitSharingOrdersUnfreezeResult {
	return &ProfitSharingOrdersUnfreezeResult{Result: result, Body: body, Http: http, Err: err, ApiError: apiError}
}

// ProfitSharingOrdersUnfreeze 解冻剩余资金API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_5.shtml
func (c *Client) ProfitSharingOrdersUnfreeze(ctx context.Context, transactionId, outOrderNo, description string) *ProfitSharingOrdersUnfreezeResult {
	// 参数
	params := gorequest.NewParams()
	params.Set("sub_mchid", c.GetSubMchId())    // 子商户号
	params.Set("transaction_id", transactionId) // 微信订单号
	params.Set("out_order_no", outOrderNo)      // 商户分账单号
	params.Set("description", description)      // 分账描述
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/profitsharing/orders/unfreeze", params, http.MethodPost)
	if err != nil {
		return newProfitSharingOrdersUnfreezeResult(ProfitSharingOrdersUnfreezeResponse{}, request.ResponseBody, request, err, ApiError{})
	}
	// 定义
	var response ProfitSharingOrdersUnfreezeResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = json.Unmarshal(request.ResponseBody, &apiError)
	return newProfitSharingOrdersUnfreezeResult(response, request.ResponseBody, request, err, apiError)
}
