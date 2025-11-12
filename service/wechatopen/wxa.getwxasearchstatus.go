package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type GetSearchStatusResponse struct {
	APIResponse       // 错误
	Status      int64 `json:"status"` // 1 表示不可搜索，0 表示可搜索
}

// GetSearchStatus 获取搜索状态
// https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/basic-info-management/api_getsearchstatus.html
func (c *Client) GetSearchStatus(ctx context.Context, notMustParams ...*gorequest.Params) (response GetSearchStatusResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, c.WithUrlAuthorizerAccessToken("wxa/getwxasearchstatus"), params, http.MethodGet, &response)
	return
}

// GetSearchStatusErrcodeInfo 错误描述
func GetSearchStatusErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 40001:
		return "获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的公众号调用接口"
	default:
		return errmsg
	}
}
