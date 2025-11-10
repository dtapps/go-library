package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// WxaApiWxaembeddedAddEmbedded 添加半屏小程序
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/embedded-management/addEmbedded.html
func (c *Client) WxaApiWxaembeddedAddEmbedded(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxaapi/wxaembedded/add_embedded?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetWxaApiWxaembeddedAddEmbeddedErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 89408:
		return "半屏小程序系统错误"
	case 89409:
		return "获取半屏小程序列表参数错误"
	case 89410:
		return "添加半屏小程序 appid 参数错误"
	case 89411:
		return "添加半屏小程序 appid 参数为空"
	case 89412:
		return "添加半屏小程序申请理由不得超过30个字"
	case 89413:
		return "该小程序被申请次数已达24h限制"
	case 89414:
		return "每天仅允许申请50次半屏小程序"
	case 89420:
		return "不支持添加个人主体小程序"
	case 89423:
		return "申请次数添加到达上限"
	case 89425:
		return "申请添加已超时"
	case 89426:
		return "申请添加状态异常"
	case 89427:
		return "申请号和授权号相同"
	case 89428:
		return "该小程序已申请，不允许重复添加"
	case 89429:
		return "已到达同一小程序每日最多申请次数"
	case 89430:
		return "该小程序已设置自动拒绝申请"
	case 89431:
		return "不支持此类型小程序"
	case 89432:
		return "不是小程序"
	case 89418:
		return "获取半屏小程序每日申请次数失败"
	case 89424:
		return "授权次数到达上限"
	case 89419:
		return "获取半屏小程序每日授权次数失败"
	default:
		return errmsg
	}
}
