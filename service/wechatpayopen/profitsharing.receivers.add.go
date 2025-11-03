package wechatpayopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type ProfitSharingReceiversAddResponse struct {
	SubMchid       string `json:"sub_mchid"`                 // 子商户号
	Type           string `json:"type"`                      // 分账接收方类型
	Account        string `json:"account"`                   // 分账接收方账号
	Name           string `json:"name,omitempty"`            // 分账接收方全称
	RelationType   string `json:"relation_type"`             // 与分账方的关系类型
	CustomRelation string `json:"custom_relation,omitempty"` // 自定义的分账关系
}

// ProfitSharingReceiversAdd 添加分账接收方API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_8.shtml
func (c *Client) ProfitSharingReceiversAdd(ctx context.Context, Type, account, name, relationType, customRelation string, notMustParams ...*gorequest.Params) (response ProfitSharingReceiversAddResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号
	params.Set("appid", c.GetSpAppid())      // 应用ID
	params.Set("sub_appid", c.GetSubAppid()) // 子商户应用ID
	params.Set("type", Type)                 // 分账接收方类型
	if Type == MERCHANT_ID {
		params.Set("account", account) // 商户号
		params.Set("name", name)       // 商户全称
	}
	if Type == PERSONAL_OPENID && name != "" {
		params.Set("account", account) // 个人openid
		params.Set("name", name)       // 个人姓名
	}
	if Type == PERSONAL_SUB_OPENID && name != "" {
		params.Set("account", account) // 个人sub_openid
		params.Set("name", name)       // 个人姓名
	}
	params.Set("relation_type", relationType) // 与分账方的关系类型
	if relationType == "CUSTOM" {
		params.Set("custom_relation", customRelation) // 自定义的分账关系
	}

	// 请求
	err = c.request(ctx, "v3/profitsharing/receivers/delete", params, http.MethodPost, &response, &apiError)
	return
}
