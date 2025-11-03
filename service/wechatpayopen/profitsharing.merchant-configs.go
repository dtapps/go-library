package wechatpayopen

import (
	"context"
	"fmt"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type ProfitSharingMerchantConfigsResponse struct {
	SubMchid string `json:"sub_mchid"` // 子商户号
	MaxRatio int    `json:"max_ratio"` // 最大分账比例
}

// ProfitSharingMerchantConfigs 查询最大分账比例API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_7.shtml
func (c *Client) ProfitSharingMerchantConfigs(ctx context.Context, notMustParams ...*gorequest.Params) (response ProfitSharingMerchantConfigsResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, fmt.Sprintf("v3/profitsharing/merchant-configs/%s", c.GetSubMchId()), params, http.MethodGet, &response, &apiError)
	return
}
