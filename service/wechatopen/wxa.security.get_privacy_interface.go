package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type GetPrivacyInterfaceResponse struct {
	APIResponse   // 错误
	InterfaceList []struct {
		ApiName    string `json:"api_name"`              // api 英文名
		ApiChName  string `json:"api_ch_name"`           // api 中文名
		ApiDesc    string `json:"api_desc"`              // api描述
		ApplyTime  int64  `json:"apply_time,omitempty"`  // 申请时间 ，该字段发起申请后才会有
		Status     int    `json:"status,omitempty"`      // 接口状态，该字段发起申请后才会有 1待申请开通 2无权限 3申请中 4申请失败 5已开通
		AuditId    int    `json:"audit_id,omitempty"`    // 申请单号，该字段发起申请后才会有
		FailReason string `json:"fail_reason,omitempty"` // 申请被驳回原因或者无权限，该字段申请驳回时才会有
		ApiLink    string `json:"api_link"`              // api文档链接
		GroupName  string `json:"group_name"`            // 分组名
	} `json:"interface_list"` // 隐私接口
}

// GetPrivacyInterface 获取接口列表
// https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/privacy-api-management/api_getprivacyinterface.html
func (c *Client) GetPrivacyInterface(ctx context.Context, notMustParams ...*gorequest.Params) (response GetPrivacyInterfaceResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/security/get_privacy_interface?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodGet, &response)
	return
}

// ErrcodeInfo 错误描述
func GetGetPrivacyInterfaceErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 61031:
		return "审核中，请不要重复申请"
	default:
		return errmsg
	}
}
