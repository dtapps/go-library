package wechatpayopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type ProfitSharingReceiversDeleteResponse struct {
	SubMchid string `json:"sub_mchid"` // 子商户号
	Type     string `json:"type"`      // 分账接收方类型
	Account  string `json:"account"`   // 分账接收方账号
}

// ProfitSharingReceiversDelete 删除分账接收方API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_9.shtml
func (c *Client) ProfitSharingReceiversDelete(ctx context.Context, Type, account string, notMustParams ...*gorequest.Params) (response ProfitSharingReceiversDeleteResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号
	params.Set("appid", c.GetSpAppid())      // 应用ID
	params.Set("sub_appid", c.GetSubAppid()) // 子商户应用ID
	params.Set("type", Type)                 // 分账接收方类型
	if Type == MERCHANT_ID {
		params.Set("account", account) // 商户号
	}
	if Type == PERSONAL_OPENID {
		params.Set("account", account) // 个人openid
	}
	if Type == PERSONAL_SUB_OPENID {
		params.Set("account", account) // 个人sub_openid
	}

	// 请求
	err = c.request(ctx, "v3/profitsharing/receivers/delete", params, http.MethodPost, &response, &apiError)
	return
}
