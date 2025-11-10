package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type CgiBinComponentApiStartPushTicketResponse struct {
	APIResponse        // 错误
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值
}

// CgiBinComponentApiStartPushTicket 启动ticket推送服务
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_verify_ticket_service.html
func (c *Client) CgiBinComponentApiStartPushTicket(ctx context.Context, notMustParams ...*gorequest.Params) (response CgiBinComponentApiStartPushTicketResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("component_appid", c.GetComponentAppId())      // 第三方平台appid
	params.Set("component_secret", c.GetComponentAppSecret()) // 第三方平台app_secret

	// 请求
	err = c.request(ctx, "cgi-bin/component/api_start_push_ticket", params, http.MethodPost, &response)
	return
}
