package pinduoduo

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type PddDdkOauthCashGiftStatusUpdateResponse struct {
	UpdateCashgiftResponse struct {
		CashGiftId float64 `json:"cash_gift_id"` // 多多礼金ID
	} `json:"update_cashgift_response"`
}

type PddDdkOauthCashGiftStatusUpdateResult struct {
	Result PddDdkOauthCashGiftStatusUpdateResponse // 结果
	Body   []byte                                  // 内容
	Http   gorequest.Response                      // 请求
}

func newPddDdkOauthCashGiftStatusUpdateResult(result PddDdkOauthCashGiftStatusUpdateResponse, body []byte, http gorequest.Response) *PddDdkOauthCashGiftStatusUpdateResult {
	return &PddDdkOauthCashGiftStatusUpdateResult{Result: result, Body: body, Http: http}
}

// OauthCashGiftStatusUpdate 多多礼金状态更新接口
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.oauth.cashgift.status.update
func (c *Client) OauthCashGiftStatusUpdate(ctx context.Context, notMustParams ...gorequest.Params) (*PddDdkOauthCashGiftStatusUpdateResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "pdd.ddk.oauth.cashgift.status.update")
	defer c.TraceEndSpan()

	// 参数
	params := NewParamsWithType("pdd.ddk.oauth.cashgift.status.update", notMustParams...)

	// 请求
	var response PddDdkOauthCashGiftStatusUpdateResponse
	request, err := c.request(ctx, params, &response)
	return newPddDdkOauthCashGiftStatusUpdateResult(response, request.ResponseBody, request), err
}
