package pinduoduo

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type OauthCashGiftStatusUpdate struct {
	UpdateCashgiftResponse struct {
		CashGiftId float64 `json:"cash_gift_id"` // 多多礼金ID
	} `json:"update_cashgift_response"`
}

// OauthCashGiftStatusUpdate 多多礼金状态更新接口
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.oauth.cashgift.status.update
func (c *Client) OauthCashGiftStatusUpdate(ctx context.Context, notMustParams ...*gorequest.Params) (response OauthCashGiftStatusUpdate, err error) {

	// 参数
	params := NewParamsWithType("pdd.ddk.oauth.cashgift.status.update", notMustParams...)

	// 请求
	err = c.request(ctx, params, &response)
	return
}
