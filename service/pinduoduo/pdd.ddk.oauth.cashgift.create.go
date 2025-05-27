package pinduoduo

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type OauthCashGiftCreate struct {
	CreateCashgiftResponse struct {
		CashGiftId float64 `json:"cash_gift_id"` // 礼金ID
		Success    bool    `json:"success"`      // 创建结果
	} `json:"create_cashgift_response"`
}

// OauthCashGiftCreate 创建多多礼金
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.oauth.cashgift.create
func (c *Client) OauthCashGiftCreate(ctx context.Context, notMustParams ...*gorequest.Params) (response OauthCashGiftCreate, err error) {

	// 参数
	params := NewParamsWithType("pdd.ddk.oauth.cashgift.create", notMustParams...)

	// 请求
	err = c.request(ctx, params, &response)
	return
}
