package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// WxaApiWxAembeddedSetAuthorize 设置授权方式
// checkComponentIsConfig && checkAuthorizerConfig
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/embedded-management/setAuthorizedEmbedded.html
func (c *Client) WxaApiWxAembeddedSetAuthorize(ctx context.Context, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxaapi/wxaembedded/set_authorize?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetWxaApiWxAembeddedSetAuthorizeErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 89417:
		return "修改半屏小程序方式 flag 参数错误"
	default:
		return errmsg
	}
}
