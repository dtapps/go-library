package pinduoduo

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
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
	Err    error                                   // 错误
}

func newPddDdkOauthCashGiftStatusUpdateResult(result PddDdkOauthCashGiftStatusUpdateResponse, body []byte, http gorequest.Response, err error) *PddDdkOauthCashGiftStatusUpdateResult {
	return &PddDdkOauthCashGiftStatusUpdateResult{Result: result, Body: body, Http: http, Err: err}
}

// StatusUpdate 多多礼金状态更新接口
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.oauth.cashgift.status.update
func (c *PddDdkOauthCashGiftApi) StatusUpdate(ctx context.Context, notMustParams ...Params) *PddDdkOauthCashGiftStatusUpdateResult {
	// 参数
	params := NewParamsWithType("pdd.ddk.oauth.cashgift.status.update", notMustParams...)
	// 请求
	request, err := c.client.request(ctx, params)
	// 定义
	var response PddDdkOauthCashGiftStatusUpdateResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPddDdkOauthCashGiftStatusUpdateResult(response, request.ResponseBody, request, err)
}
