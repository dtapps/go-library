package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type WxaModifyDomainResponse struct {
	APIResponse                     // 错误
	Requestdomain          []string `json:"requestdomain"`           // request 合法域名
	Wsrequestdomain        []string `json:"wsrequestdomain"`         // socket 合法域名
	Uploaddomain           []string `json:"uploaddomain"`            // uploadFile 合法域名
	Downloaddomain         []string `json:"downloaddomain"`          // downloadFile 合法域名
	Udpdomain              []string `json:"udpdomain"`               // udp 合法域名
	Tcpdomain              []string `json:"tcpdomain"`               // tcp 合法域名
	InvalidRequestdomain   []string `json:"invalid_requestdomain"`   // request 不合法域名
	InvalidWsrequestdomain []string `json:"invalid_wsrequestdomain"` // socket 不合法域名
	InvalidUploaddomain    []string `json:"invalid_uploaddomain"`    // uploadFile 不合法域名
	InvalidDownloaddomain  []string `json:"invalid_downloaddomain"`  // downloadFile 不合法域名
	InvalidUdpdomain       []string `json:"invalid_udpdomain"`       // udp 不合法域名
	InvalidTcpdomain       []string `json:"invalid_tcpdomain"`       // tcp 不合法域名
	NoIcpDomain            []string `json:"no_icp_domain"`           // 没有经过icp备案的域名
}

// WxaModifyDomain 配置小程序服务器域名
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/domain-management/modifyServerDomain.html
func (c *Client) WxaModifyDomain(ctx context.Context, notMustParams ...*gorequest.Params) (response WxaModifyDomainResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/modify_domain?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetWxaModifyDomainErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 85015:
		return "该账号不是小程序账号"
	case 85016:
		return "域名数量超限制"
	case 85017:
		return "域名输入为空，或者没有新增域名，请确认小程序已经添加了域名或该域名是否没有在第三方平台添加"
	case 85018:
		return "域名没有在第三方平台设置"
	case 85301:
		return "存在 “不符合域名规则的域名”导致无修改"
	case 85302:
		return "存在 “ 缺少ICP备案的域名”导致无修改"
	case 85303:
		return "同时存在“不符合域名规则的域名”以及“ 缺少ICP备案的域名”导致无修改"
	default:
		return errmsg
	}
}
