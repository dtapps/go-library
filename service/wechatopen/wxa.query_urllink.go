package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type QueryUrlLinkResponse struct {
	APIResponse // 错误
	UrlLinkInfo struct {
		Appid      string `json:"appid"`       // 小程序 appid
		Path       string `json:"path"`        // 小程序页面路径
		Query      string `json:"query"`       // 小程序页面query
		CreateTime int    `json:"create_time"` // 创建时间，为 Unix 时间戳
		ExpireTime int    `json:"expire_time"` // 到期失效时间，为 Unix 时间戳，0 表示永久生效
		EnvVersion string `json:"env_version"` // 要打开的小程序版本。正式版为"release"，体验版为"trial"，开发版为"develop"
	} `json:"url_link_info"` // url_link 配置
	QuotaInfo struct {
		RemainVisitQuota string `json:"remain_visit_quota"` // 	URL Scheme（加密+明文）/加密 URL Link 单天剩余访问次数
	} `json:"quota_info"`
}

// QueryUrlLink 查询加密URLLink
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-link/queryUrlLink.html
func (c *Client) QueryUrlLink(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (response QueryUrlLinkResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/query_urllink?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetQueryUrlLinkErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 40097:
		return "参数错误"
	case 85403:
		return "scheme/url link不存在"
	default:
		return errmsg
	}
}
