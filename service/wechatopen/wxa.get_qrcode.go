package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// GetTrialQRCode 获取体验版二维码
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/getTrialQRCode.html
func (c *Client) GetTrialQRCode(ctx context.Context, authorizerAccessToken, path string, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	if path != "" {
		params.Set("path", path) // 指定二维码扫码后直接进入指定页面并可同时带上参数）
	}

	// 请求
	err = c.request(ctx, "wxa/get_qrcode?access_token="+authorizerAccessToken, params, http.MethodGet, &response)

	// 判断内容是否为图片
	//if request.HeaderIsImg() == false {
	//	err = json.Unmarshal(request.ResponseBody, &response)
	//}
	return
}
