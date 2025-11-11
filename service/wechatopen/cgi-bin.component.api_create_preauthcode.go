package wechatopen

import (
	"context"
	"fmt"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type CgiBinComponentApiCreatePreAuthCodenResponse struct {
	PreAuthCode string `json:"pre_auth_code"` // 预授权码
	ExpiresIn   int64  `json:"expires_in"`    // 有效期，单位：秒
}

// CgiBinComponentApiCreatePreAuthCoden 预授权码
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/pre_auth_code.html
func (c *Client) CgiBinComponentApiCreatePreAuthCoden(ctx context.Context, notMustParams ...*gorequest.Params) (response CgiBinComponentApiCreatePreAuthCodenResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("component_appid", c.GetComponentAppId()) // 第三方平台appid

	// 请求
	err = c.request(ctx, fmt.Sprintf("cgi-bin/component/api_create_preauthcode?component_access_token=%s", c.GetComponentAccessToken()), params, http.MethodPost, &response)
	return
}
