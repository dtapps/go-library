package wechatpayopen

import (
	"context"
	"fmt"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type FundAppMchTransferPartnerSubMchAuthorizationsSubMchidResponse struct {
	SubMchid   string `json:"sub_mchid"` // 【子商户号】 微信支付分配的商户号
	State      string `json:"state"`     // 【授权状态】 授权状态 AUTHORIZED:  已授权 UNAUTHORIZED:  未授权 CANCELED:  已解除
	ApplyOrder struct {
		SubMchid   string `json:"sub_mchid"`    // 【子商户号】 微信支付分配的商户号
		OutApplyNo string `json:"out_apply_no"` // 【商户申请单号】 商户申请单号
		State      string `json:"state"`        // 【授权状态】 授权状态 AUTHORIZED:  已授权 UNAUTHORIZED:  未授权 CANCELED:  已解除
	} `json:"apply_order"` // 【授权申请单】 当存在授权申请单时，返回最近一次授权申请单信息
	AuthorizationTime       string `json:"authorization_time"`        // 【授权时间】 当授权状态为AUTHORIZED，返回授权时间
	CancelAuthorizationTime string `json:"cancel_authorization_time"` // 【解除授权时间】 当授权状态为CANCELED，返回解除授权时间
}

// FundAppMchTransferPartnerSubMchAuthorizations 查询子商户商家转账授权状态
// https://pay.weixin.qq.com/doc/v3/partner/4015469112
func (c *Client) FundAppMchTransferPartnerSubMchAuthorizationsSubMchid(ctx context.Context, notMustParams ...*gorequest.Params) (response FundAppMchTransferPartnerSubMchAuthorizationsSubMchidResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.NewRequest(ctx, fmt.Sprintf("/v3/fund-app/mch-transfer/partner/sub-mch-authorizations/%s", c.GetSubMchId()), params, http.MethodGet, &response, &apiError)
	return
}
