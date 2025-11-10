package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type CgiBinGetApiDomainIpResponse struct {
	IpList []string `json:"ip_list,omitempty"`
}

// CgiBinGetApiDomainIp 获取微信API接口 IP地址
// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
func (c *Client) CgiBinGetApiDomainIp(ctx context.Context, componentAccessToken string, notMustParams ...*gorequest.Params) (response CgiBinGetApiDomainIpResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "cgi-bin/get_api_domain_ip?access_token="+componentAccessToken, params, http.MethodGet, &response)
	return
}
