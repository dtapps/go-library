package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type TckWxPayListResponse struct {
	APIResponse // 错误
	List        []struct {
		MerchantCode     string `json:"merchant_code"`
		MerchantName     string `json:"merchant_name"`
		CompanyName      string `json:"company_name"`
		MchRelationState string `json:"mch_relation_state"`
		JsapiAuthState   string `json:"jsapi_auth_state"`
		RefundAuthState  string `json:"refund_auth_state"`
	} `json:"list"`
}

// TckWxPayList 获取授权绑定的商户号列表
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/cloudbase-common/wechatpay/getWechatPayList.html
func (c *Client) TckWxPayList(ctx context.Context, componentAccessToken string, notMustParams ...*gorequest.Params) (response TckWxPayListResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "tcb/wxpaylist?access_token="+componentAccessToken, params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetTckWxPayListErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 85064:
		return "找不到草稿"
	case 85065:
		return "模板库已满"
	default:
		return errmsg
	}
}
