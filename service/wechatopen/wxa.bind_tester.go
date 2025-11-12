package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type BindTesterResponse struct {
	APIResponse        // 错误
	Userstr     string `json:"userstr"` // 人员对应的唯一字符串
}

// BindTester 绑定体验者
// https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/member-management/api_bindtester.html
func (c *Client) BindTester(ctx context.Context, wechatid string, notMustParams ...*gorequest.Params) (response BindTesterResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("wechatid", wechatid)

	// 请求
	err = c.request(ctx, "wxa/bind_tester?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetBindTesterErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 85001:
		return "微信号不存在或微信号设置为不可搜索"
	case 85002:
		return "小程序绑定的体验者数量达到上限"
	case 85003:
		return "微信号绑定的小程序体验者达到上限"
	case 85004:
		return "微信号已经绑定"
	default:
		return errmsg
	}
}
