package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// GetTrialQRCode 获取体验版二维码
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/getTrialQRCode.html
func (c *Client) GetTrialQRCode(ctx context.Context, path string, notMustParams ...*gorequest.Params) (response APIResponse, body []byte, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	if path != "" {
		params.Set("path", path) // 指定二维码扫码后直接进入指定页面并可同时带上参数）
	}

	// 请求
	body, err = c.requestImage(ctx, "wxa/get_qrcode?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodGet, &response)
	return
}
