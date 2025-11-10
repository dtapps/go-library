package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// Release 发布已通过审核的小程序
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/release.html
func (c *Client) Release(ctx context.Context, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/release?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetReleaseErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 85019:
		return "没有审核版本"
	case 85020:
		return "审核状态未满足发布"
	default:
		return errmsg
	}
}
