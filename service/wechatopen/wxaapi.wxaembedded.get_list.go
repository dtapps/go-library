package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type WxaApiWxaembeddedGetListResponse struct {
	APIResponse         // 错误
	EmbeddedFlag    int `json:"embedded_flag"` // 授权方式。0表示需要管理员确认，1表示自动通过，2表示自动拒绝
	WxaEmbeddedList []struct {
		Appid       string `json:"appid"`       // 半屏小程序appid
		Create_time int64  `json:"create_time"` // 添加时间
		Headimg     string `json:"headimg"`     // 头像url
		Nickname    string `json:"nickname"`    // 半屏小程序昵称
		Reason      string `json:"reason"`      // 申请理由
		Status      string `json:"status"`      // 申请状态
	} `json:"wxa_embedded_list"` // 半屏小程序列表
}

// WxaApiWxaembeddedGetList 获取半屏小程序调用列表
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/embedded-management/getEmbeddedList.html
func (c *Client) WxaApiWxaembeddedGetList(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (response WxaApiWxaembeddedGetListResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxaapi/wxaembedded/get_list?access_token="+authorizerAccessToken, params, http.MethodGet, &response)
	return
}

// ErrcodeInfo 错误描述
func GetWxaApiWxaembeddedGetListErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 89408:
		return "半屏小程序系统错误"
	case 89409:
		return "获取半屏小程序列表参数错误"
	default:
		return errmsg
	}
}
