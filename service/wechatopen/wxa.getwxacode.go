package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// GetQRCode 获取小程序码
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/getQRCode.html
func (c *Client) GetQRCode(ctx context.Context, notMustParams ...*gorequest.Params) (response APIResponse, body []byte, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	body, err = c.requestImage(ctx, "wxa/getwxacode?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetGetQRCodeErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 40001:
		return "获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的公众号调用接口"
	case 40159:
		return "path 不能为空，且长度不能大于1024"
	case 45029:
		return "生成码个数总和到达最大个数限制"
	case 85096:
		return "scancode_time为系统保留参数，不允许配置"
	case 40097:
		return "参数错误"
	default:
		return errmsg
	}
}
