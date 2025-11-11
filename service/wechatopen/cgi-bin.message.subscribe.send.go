package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// 发送订阅消息
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-message-management/subscribe-message/sendMessage.html
func (c *Client) SendMessage(ctx context.Context, notMustParams ...*gorequest.Params) (response APIResponse, err error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	err = c.request(ctx, c.WithUrlAuthorizerAccessToken("cgi-bin/message/subscribe/send"), params, http.MethodPost, &response)
	return
}
