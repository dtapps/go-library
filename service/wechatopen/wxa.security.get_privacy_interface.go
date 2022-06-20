package wechatopen

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WxaSecurityGetPrivacyInterfaceResponse struct {
	Errcode       int    `json:"errcode"` // 返回码
	Errmsg        string `json:"errmsg"`  // 返回码信息
	InterfaceList []struct {
		ApiName    string `json:"api_name"`              // api 英文名
		ApiChName  string `json:"api_ch_name"`           // api 中文名
		ApiDesc    string `json:"api_desc"`              // api描述
		ApplyTime  int64  `json:"apply_time,omitempty"`  // 申请时间 ，该字段发起申请后才会有
		Status     int    `json:"status"`                // 接口状态，该字段发起申请后才会有
		AuditId    int    `json:"audit_id,omitempty"`    // 申请单号，该字段发起申请后才会有
		FailReason string `json:"fail_reason,omitempty"` // 申请被驳回原因或者无权限，该字段申请驳回时才会有
		ApiLink    string `json:"api_link"`              // api文档链接
		GroupName  string `json:"group_name"`            // 分组名
	} `json:"interface_list"` // 隐私接口
}

type WxaSecurityGetPrivacyInterfaceResult struct {
	Result WxaSecurityGetPrivacyInterfaceResponse // 结果
	Body   []byte                                 // 内容
	Http   gorequest.Response                     // 请求
	Err    error                                  // 错误
}

func newWxaSecurityGetPrivacyInterfaceResult(result WxaSecurityGetPrivacyInterfaceResponse, body []byte, http gorequest.Response, err error) *WxaSecurityGetPrivacyInterfaceResult {
	return &WxaSecurityGetPrivacyInterfaceResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaSecurityGetPrivacyInterface 获取接口列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/apply_api/get_privacy_interface.html
func (c *Client) WxaSecurityGetPrivacyInterface() *WxaSecurityGetPrivacyInterfaceResult {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(fmt.Sprintf(apiUrl+"/wxa/security/get_privacy_interface?access_token=%s", c.GetAuthorizerAccessToken()), params, http.MethodGet)
	// 定义
	var response WxaSecurityGetPrivacyInterfaceResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newWxaSecurityGetPrivacyInterfaceResult(response, request.ResponseBody, request, err)
}

// ErrcodeInfo 错误描述
func (resp *WxaSecurityGetPrivacyInterfaceResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 61031:
		return "审核中，请不要重复申请"
	}
	return "系统繁忙"
}
