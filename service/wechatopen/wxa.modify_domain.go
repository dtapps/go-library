package wechatopen

import (
	"encoding/json"
	"fmt"
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
	Err    error                   // 错误
}

func NewWxaModifyDomainResult(result WxaModifyDomainResponse, body []byte, err error) *WxaModifyDomainResult {
	return &WxaModifyDomainResult{Result: result, Body: body, Err: err}
}

// WxaModifyDomain 设置服务器域名
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/Server_Address_Configuration.html
func (app *App) WxaModifyDomain(notMustParams ...Params) *WxaModifyDomainResult {
	app.authorizerAccessToken = app.GetAuthorizerAccessToken()
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/modify_domain?access_token=%s", app.authorizerAccessToken), params, http.MethodPost)
	// 定义
	var response WxaModifyDomainResponse
	err = json.Unmarshal(body, &response)
	return NewWxaModifyDomainResult(response, body, err)
}
