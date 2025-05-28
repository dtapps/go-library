package wechatqy

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type CgiBinGetCallBackIp struct {
	IpList  []string `json:"ip_list,omitempty"`
	Errcode int      `json:"errcode"`
	Errmsg  string   `json:"errmsg"`
}

// CgiBinGetCallBackIp 获取企业微信回调IP段
// https://developer.work.weixin.qq.com/document/path/98988
func (c *Client) CgiBinGetCallBackIp(ctx context.Context, accessToken string, notMustParams ...*gorequest.Params) (response CgiBinGetCallBackIp, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, apiUrl+"cgi-bin/getcallbackip?access_token="+accessToken, params, http.MethodGet, &response)
	return
}
