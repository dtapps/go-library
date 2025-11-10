package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// WxaApiWxaembeddedDelEmbedded 删除半屏小程序
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/embedded-management/deleteEmbedded.html
func (c *Client) WxaApiWxaembeddedDelEmbedded(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxaapi/wxaembedded/del_embedded?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetWxaApiWxaembeddedDelEmbeddedErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 89408:
		return "半屏小程序系统错误"
	case 89415:
		return "删除半屏小程序 appid 参数为空"
	case 89421:
		return "删除数据未找到"
	case 89422:
		return "删除状态异常"
	case 89431:
		return "不支持此类型小程序"
	case 89432:
		return "不是小程序"
	default:
		return errmsg
	}
}
