package wechatpayopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type FundAppMchTransferPartnerSubMchAuthorizationsResponse struct {
	SubMchid   string `json:"sub_mchid"`    // 【子商户号】 微信支付分配的商户号
	OutApplyNo string `json:"out_apply_no"` // 【商户申请单号】 商户申请单号
	State      string `json:"state"`        // 【【授权申请单状态】 授权申请单状态 INVITED:  已邀请 PASSED:  已通过 REJECTED:  已拒绝 EXPIRED:  已过期
	AcceptTime string `json:"accept_time"`  // 【授权申请受理时间】 使用rfc3339所定义的格式，格式为yyyy-MM-DDThh:mm:ss+TIMEZONE
	UpdateTime string `json:"update_time"`  // 【最后一次状态变更时间】 单据最后更新时间，按照使用rfc3339所定义的格式，格式为yyyy-MM-DDThh:mm:ss+TIMEZONE
}

// FundAppMchTransferPartnerSubMchAuthorizations 申请子商户商家转账授权
// https://pay.weixin.qq.com/doc/v3/partner/4015469102
func (c *Client) FundAppMchTransferPartnerSubMchAuthorizations(ctx context.Context, notMustParams ...*gorequest.Params) (response FundAppMchTransferPartnerSubMchAuthorizationsResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号

	// 请求
	err = c.NewRequest(ctx, "/v3/fund-app/mch-transfer/partner/sub-mch-authorizations", params, http.MethodPost, &response, &apiError)
	return
}
