package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type RevertCodeReleaseResponse struct {
	APIResponse // 错误
	VersionList []struct {
		CommitTime  int    `json:"commit_time"`  // 更新时间，时间戳
		UserVersion string `json:"user_version"` // 模板版本号，开发者自定义字段
		UserDesc    string `json:"user_desc"`    // 模板描述，开发者自定义字段
		AppVersion  int    `json:"app_version"`  // 小程序版本
	} `json:"version_list"` // 模板信息列表
}

// RevertCodeRelease 小程序版本回退
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/revertCodeRelease.html
func (c *Client) RevertCodeRelease(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (response RevertCodeReleaseResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/revertcoderelease?access_token="+authorizerAccessToken, params, http.MethodGet, &response)
	return
}

// ErrcodeInfo 错误描述
func GetRevertCodeReleaseErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 40001:
		return "获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的公众号调用接口"
	default:
		return errmsg
	}
}
