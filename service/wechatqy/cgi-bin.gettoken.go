package wechatqy

import (
	"context"
	"fmt"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type CgiBinGetToken struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// CgiBinGetToken 获取access_token
// https://open.work.weixin.qq.com/api/doc/90000/90135/91039
func (c *Client) CgiBinGetToken(ctx context.Context, notMustParams ...*gorequest.Params) (response CgiBinGetToken, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, apiUrl+fmt.Sprintf("cgi-bin/gettoken?corpid=%s&corpsecret=%s", c.GetAppId(), c.GetSecret()), params, http.MethodGet, &response)
	return
}
