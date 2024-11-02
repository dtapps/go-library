package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type CgiBinGetApiDomainIpResponse struct {
	IpList []string `json:"ip_list,omitempty"`
}

type CgiBinGetApiDomainIpResult struct {
	Result CgiBinGetApiDomainIpResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
}

func NewCgiBinGetApiDomainIpResult(result CgiBinGetApiDomainIpResponse, body []byte, http gorequest.Response) *CgiBinGetApiDomainIpResult {
	return &CgiBinGetApiDomainIpResult{Result: result, Body: body, Http: http}
}

// CgiBinGetApiDomainIp 获取微信API接口 IP地址
// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
func (c *Client) CgiBinGetApiDomainIp(ctx context.Context, componentAccessToken string, notMustParams ...gorequest.Params) (*CgiBinGetApiDomainIpResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response CgiBinGetApiDomainIpResponse
	request, err := c.request(ctx, "cgi-bin/get_api_domain_ip?access_token="+componentAccessToken, params, http.MethodGet, &response)
	return NewCgiBinGetApiDomainIpResult(response, request.ResponseBody, request), err
}
