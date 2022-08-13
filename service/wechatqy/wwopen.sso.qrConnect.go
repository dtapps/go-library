package wechatqy

import (
	"context"
	"fmt"
)

// WwOpenSsoQrConnect 构造独立窗口登录二维码
// https://open.work.weixin.qq.com/api/doc/90000/90135/91019
func (c *Client) WwOpenSsoQrConnect(ctx context.Context) string {
	return fmt.Sprintf("https://open.work.weixin.qq.com/wwopen/sso/qrConnect?appid=%s&agentid=%d&redirect_uri=%s&state=STATE&lang=zh", c.GetAppId(), c.GetAgentId(), c.GetRedirectUri())
}
