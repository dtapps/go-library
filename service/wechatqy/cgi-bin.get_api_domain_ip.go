package wechatqy

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type CgiBinGetApiDomainIpResponse struct {
	IpList  []string `json:"ip_list,omitempty"`
	Errcode int      `json:"errcode"`
	Errmsg  string   `json:"errmsg"`
}

type CgiBinGetApiDomainIpResult struct {
	Result CgiBinGetApiDomainIpResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
}

func NewCgiBinGetApiDomainIpResult(result CgiBinGetApiDomainIpResponse, body []byte, http gorequest.Response) *CgiBinGetApiDomainIpResult {
	return &CgiBinGetApiDomainIpResult{Result: result, Body: body, Http: http}
}

// CgiBinGetApiDomainIp 获取企业微信接口IP段
// https://developer.work.weixin.qq.com/document/path/97073
func (c *Client) CgiBinGetApiDomainIp(ctx context.Context, accessToken string, notMustParams ...gorequest.Params) (*CgiBinGetApiDomainIpResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "cgi-bin/get_api_domain_ip")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response CgiBinGetApiDomainIpResponse
	request, err := c.request(ctx, span, apiUrl+"cgi-bin/get_api_domain_ip?access_token="+accessToken, params, http.MethodGet, &response)
	return NewCgiBinGetApiDomainIpResult(response, request.ResponseBody, request), err
}
