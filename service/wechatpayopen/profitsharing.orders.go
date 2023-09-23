package wechatpayopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type ProfitSharingOrdersResponse struct {
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

type ProfitSharingOrdersResult struct {
	Result ProfitSharingOrdersResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
}

func newProfitSharingOrdersResult(result ProfitSharingOrdersResponse, body []byte, http gorequest.Response) *ProfitSharingOrdersResult {
	return &ProfitSharingOrdersResult{Result: result, Body: body, Http: http}
}

// ProfitSharingOrders 请求分账API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_1.shtml
func (c *Client) ProfitSharingOrders(ctx context.Context, notMustParams ...*gorequest.Params) (*ProfitSharingOrdersResult, ApiError, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号
	params.Set("appid", c.GetSpAppid())      // 应用ID
	params.Set("sub_appid", c.GetSubAppid()) // 子商户应用ID
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/profitsharing/orders", params, http.MethodPost)
	if err != nil {
		return newProfitSharingOrdersResult(ProfitSharingOrdersResponse{}, request.ResponseBody, request), ApiError{}, err
	}
	// 定义
	var response ProfitSharingOrdersResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = gojson.Unmarshal(request.ResponseBody, &apiError)
	return newProfitSharingOrdersResult(response, request.ResponseBody, request), apiError, err
}
