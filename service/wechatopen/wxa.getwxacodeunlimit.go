package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// GetUnlimitedQRCode 获取不限制的小程序码
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/getUnlimitedQRCode.html
func (c *Client) GetUnlimitedQRCode(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (response APIResponse, body []byte, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	body, err = c.requestImage(ctx, "wxa/getwxacodeunlimit?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetGetUnlimitedQRCodeErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 45009:
		return "调用分钟频率受限(目前5000次/分钟，会调整)，如需大量小程序码，建议预生成"
	case 41030:
		return "page 不合法（页面不存在或者小程序没有发布、根路径前加 /或者携带参数）"
	case 40097:
		return "env_version 不合法"
	default:
		return errmsg
	}
}
