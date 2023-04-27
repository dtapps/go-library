package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaModifyDomainResponse struct {
	Errcode                int      `json:"errcode"`                 // 错误码
	Errmsg                 string   `json:"errmsg"`                  // 错误信息
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

type WxaModifyDomainResult struct {
	Result WxaModifyDomainResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
}

func newWxaModifyDomainResult(result WxaModifyDomainResponse, body []byte, http gorequest.Response) *WxaModifyDomainResult {
	return &WxaModifyDomainResult{Result: result, Body: body, Http: http}
}

// WxaModifyDomain 配置小程序服务器域名
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/domain-management/modifyServerDomain.html
func (c *Client) WxaModifyDomain(ctx context.Context, notMustParams ...gorequest.Params) (*WxaModifyDomainResult, error) {
	// 检查
	if err := c.checkAuthorizerConfig(ctx); err != nil {
		return newWxaModifyDomainResult(WxaModifyDomainResponse{}, []byte{}, gorequest.Response{}), err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/modify_domain?access_token="+c.GetAuthorizerAccessToken(ctx), params, http.MethodPost)
	if err != nil {
		return newWxaModifyDomainResult(WxaModifyDomainResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaModifyDomainResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaModifyDomainResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaModifyDomainResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
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
	}
	return "系统繁忙"
}
