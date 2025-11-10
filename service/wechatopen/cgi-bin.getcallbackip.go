package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type CgiBinGetCallBackIpResponse struct {
	IpList []string `json:"ip_list,omitempty"`
}

// CgiBinGetCallBackIp 获取微信callback IP地址
// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
func (c *Client) CgiBinGetCallBackIp(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (response CgiBinGetCallBackIpResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "cgi-bin/getcallbackip?access_token="+authorizerAccessToken, params, http.MethodGet, &response)
	return
}
