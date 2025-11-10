package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// WxaApiWxaembeddedDelAuthorize 取消授权小程序
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/embedded-management/deleteAuthorizedEmbedded.html
func (c *Client) WxaApiWxaembeddedDelAuthorize(ctx context.Context, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxaapi/wxaembedded/del_authorize?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetWxaApiWxaembeddedDelAuthorizeErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 89416:
		return "取消半屏小程序授权 appid 参数为空"
	case 89431:
		return "不支持此类型小程序"
	case 89432:
		return "不是小程序"
	default:
		return errmsg
	}
}
