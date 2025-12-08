package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type ApplyPrivacyInterfaceResponse struct {
	APIResponse       // 错误
	AuditId     int64 `json:"audit_id"` // 审核单id
}

// ApplyPrivacyInterface 申请地理位置接口
// https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/privacy-api-management/api_applyprivacyinterface.html
func (c *Client) ApplyPrivacyInterface(ctx context.Context, api_name string, notMustParams ...*gorequest.Params) (response ApplyPrivacyInterfaceResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("api_name", api_name) // 申请的 api 英文名，例如wx.choosePoi，严格区分大小写

	// 请求
	err = c.request(ctx, "wxa/security/apply_privacy_interface?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetApplyPrivacyInterfaceErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 61031:
		return "审核中，请不要重复申请"
	case 61032:
		return "视频格式不对，要传mp4格式的"
	case 61033:
		return "视频下载失败"
	case 61034:
		return "必填的参数没填，检查后重新提交"
	case 61035:
		return "输入的api（api_name严格区分大小写）无需申请，可以直接使用"
	case 61036:
		return "该帐号不可申请，请检查类目是否符合"
	case 61037:
		return "需要以ntf-8的编码格式提交"
	default:
		return errmsg
	}
}
