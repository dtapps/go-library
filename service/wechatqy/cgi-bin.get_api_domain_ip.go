package wechatqy

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type CgiBinGetApiDomainIp struct {
	IpList  []string `json:"ip_list,omitempty"`
	Errcode int      `json:"errcode"`
	Errmsg  string   `json:"errmsg"`
}

// CgiBinGetApiDomainIp 获取企业微信接口IP段
// https://developer.work.weixin.qq.com/document/path/97073
func (c *Client) CgiBinGetApiDomainIp(ctx context.Context, accessToken string, notMustParams ...*gorequest.Params) (response CgiBinGetApiDomainIp, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, apiUrl+"cgi-bin/get_api_domain_ip?access_token="+accessToken, params, http.MethodGet, &response)
	return
}
