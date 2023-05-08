package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaModifyDomainDirectlyResponse struct {
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

type WxaModifyDomainDirectlyResult struct {
	Result WxaModifyDomainDirectlyResponse // 结果
	Body   []byte                          // 内容
	Http   gorequest.Response              // 请求
}

func newWxaModifyDomainDirectlyResult(result WxaModifyDomainDirectlyResponse, body []byte, http gorequest.Response) *WxaModifyDomainDirectlyResult {
	return &WxaModifyDomainDirectlyResult{Result: result, Body: body, Http: http}
}

// WxaModifyDomainDirectly 快速配置小程序服务器域名
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/domain-management/modifyServerDomainDirectly.html
func (c *Client) WxaModifyDomainDirectly(ctx context.Context, notMustParams ...gorequest.Params) (*WxaModifyDomainDirectlyResult, error) {
	// 检查
	if err := c.checkAuthorizerConfig(ctx); err != nil {
		return newWxaModifyDomainDirectlyResult(WxaModifyDomainDirectlyResponse{}, []byte{}, gorequest.Response{}), err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/modify_domain_directly?access_token="+GetAuthorizerAccessToken(ctx, c), params, http.MethodPost)
	if err != nil {
		return newWxaModifyDomainDirectlyResult(WxaModifyDomainDirectlyResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaModifyDomainDirectlyResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaModifyDomainDirectlyResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaModifyDomainDirectlyResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 85015:
		return "该账号不是小程序账号"
	case 86100:
		return "该 URL 的协议头有误"
	case 45082:
		return "域名需要 icp 备案，否则无法添加"
	case 86101:
		return "不支持配置api.weixin.qq.com"
	case 85016:
		return "域名数量超限制"
	case 86102:
		return "每个月只能修改50次，超过域名修改次数限制"
	}
	return "系统繁忙"
}
