package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type GetCallBackIpResponse struct {
	IpList []string `json:"ip_list,omitempty"`
}

type GetCallBackIpResult struct {
	Result GetCallBackIpResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func NewGetCallBackIpResult(result GetCallBackIpResponse, body []byte, http gorequest.Response) *GetCallBackIpResult {
	return &GetCallBackIpResult{Result: result, Body: body, Http: http}
}

// CgiBinGetApiDomainIp 获取微信API接口 IP地址
// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
func (c *Client) CgiBinGetApiDomainIp(ctx context.Context, componentAccessToken string, notMustParams ...gorequest.Params) (*GetCallBackIpResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "cgi-bin/get_api_domain_ip")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response GetCallBackIpResponse
	request, err := c.request(ctx, span, "cgi-bin/get_api_domain_ip?access_token="+componentAccessToken, params, http.MethodGet, &response)
	return NewGetCallBackIpResult(response, request.ResponseBody, request), err
}
