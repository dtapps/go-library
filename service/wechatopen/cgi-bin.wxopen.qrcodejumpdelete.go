package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// DeleteJumpQRCode 删除已设置的二维码规则
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/qrcode/qrcodejumpdelete.html
func (c *Client) DeleteJumpQRCode(ctx context.Context, prefix string, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("prefix", prefix)

	// 请求
	err = c.request(ctx, "cgi-bin/wxopen/qrcodejumpdelete?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}
