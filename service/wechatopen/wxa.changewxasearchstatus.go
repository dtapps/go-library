package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// SetSearchStatus 设置搜索状态
// https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/basic-info-management/api_setsearchstatus.html
func (c *Client) SetSearchStatus(ctx context.Context, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, c.WithUrlAuthorizerAccessToken("wxa/changewxasearchstatus"), params, http.MethodPost, &response)
	return
}

// SetSearchStatusErrcodeInfo 错误描述
func SetSearchStatusErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 40001:
		return "获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的公众号调用接口"
	case 85083:
		return "搜索标记位被封禁，无法修改"
	default:
		return errmsg
	}
}
