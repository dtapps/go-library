package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type GetVersionInfoResponse struct {
	APIResponse // 错误
	ExpInfo     struct {
		ExpTime    int64  `json:"exp_time"`    // 提交体验版的时间
		ExpVersion string `json:"exp_version"` // 体验版版本信息
		ExpDesc    string `json:"exp_desc"`    // 体验版版本描述
	} `json:"exp_info"` // 体验版信息
	ReleaseInfo struct {
		ReleaseTime    int64  `json:"release_time"`    // 发布线上版的时间
		ReleaseVersion string `json:"release_version"` // 线上版版本信息
		ReleaseDesc    string `json:"release_desc"`    // 线上版本描述
	} `json:"release_info"` // 线上版信息
}

// GetVersionInfo 查询小程序版本信息
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/getVersionInfo.html
func (c *Client) GetVersionInfo(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (response GetVersionInfoResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/getversioninfo?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}
