package wechatpayopen

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type ProfitSharingOrdersOutOrderNoResponse struct {
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

type ProfitSharingOrdersOutOrderNoResult struct {
	Result   ProfitSharingOrdersOutOrderNoResponse // 结果
	Body     []byte                                // 内容
	Http     gorequest.Response                    // 请求
	Err      error                                 // 错误
	ApiError ApiError                              // 接口错误
}

func newProfitSharingOrdersOutOrderNoResult(result ProfitSharingOrdersOutOrderNoResponse, body []byte, http gorequest.Response, err error, apiError ApiError) *ProfitSharingOrdersOutOrderNoResult {
	return &ProfitSharingOrdersOutOrderNoResult{Result: result, Body: body, Http: http, Err: err, ApiError: apiError}
}

// ProfitSharingOrdersOutOrderNo 查询分账结果API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_2.shtml
func (c *Client) ProfitSharingOrdersOutOrderNo(ctx context.Context, transactionId, outOrderNo string) *ProfitSharingOrdersOutOrderNoResult {
	// 参数
	params := gorequest.NewParams()
	params.Set("sub_mchid", c.GetSubMchId())    // 子商户号
	params.Set("transaction_id", transactionId) // 微信订单号
	params.Set("out_order_no", outOrderNo)      // 商户分账单号
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/profitsharing/orders/"+outOrderNo, params, http.MethodGet)
	if err != nil {
		return newProfitSharingOrdersOutOrderNoResult(ProfitSharingOrdersOutOrderNoResponse{}, request.ResponseBody, request, err, ApiError{})
	}
	// 定义
	var response ProfitSharingOrdersOutOrderNoResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = json.Unmarshal(request.ResponseBody, &apiError)
	return newProfitSharingOrdersOutOrderNoResult(response, request.ResponseBody, request, err, apiError)
}
