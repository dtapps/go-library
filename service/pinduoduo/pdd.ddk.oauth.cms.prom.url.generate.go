package pinduoduo

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type OauthCmsUrlGenerateUrlList struct {
	MobileShortUrl           string `json:"mobile_short_url"`             // 唤醒拼多多app短链
	MobileUrl                string `json:"mobile_url"`                   // 唤醒拼多多app链接
	MultiGroupMobileShortUrl string `json:"multi_group_mobile_short_url"` // 多人团唤醒拼多多app链接
	MultiGroupMobileUrl      string `json:"multi_group_mobile_url"`       // 多人团唤醒拼多多app长链接
	MultiGroupShortUrl       string `json:"multi_group_short_url"`        // 多人团短链
	MultiGroupUrl            string `json:"multi_group_url"`              // 多人团长链
	MultiUrlList             struct {
		MobileShortUrl string `json:"mobile_short_url"` // 双人团唤醒拼多多app短链接
		MobileUrl      string `json:"mobile_url"`       // 双人团唤醒拼多多app长链接
		SchemaUrl      string `json:"schema_url"`       // schema的链接
		ShortUrl       string `json:"short_url"`        // 双人团短链接
		Url            string `json:"url"`              // 双人团长链接
	} `json:"multi_url_list"` // 双人团链接列表
	ShortUrl      string `json:"short_url"`
	Sign          string `json:"sign"`
	SingleUrlList struct {
		MobileShortUrl string `json:"mobile_short_url"` // 唤醒拼多多app短链接
		MobileUrl      string `json:"mobile_url"`       // 唤醒拼多多app长链接
		SchemaUrl      string `json:"schema_url"`       // schema的链接
		ShortUrl       string `json:"short_url"`        // 短链接
		Url            string `json:"url"`              // 长链接
	} `json:"single_url_list"` // 单人团链接列表
	Url       string `json:"url"` // h5长链接
	WeAppInfo struct {
		AppId             string `json:"app_id"`              // 小程序id
		BannerUrl         string `json:"banner_url"`          // Banner图
		Desc              string `json:"desc"`                // 描述
		PagePath          string `json:"page_path"`           // 小程序path值
		SourceDisplayName string `json:"source_display_name"` // 来源名
		Title             string `json:"title"`               // 小程序标题
		UserName          string `json:"user_name"`           // 用户名
		WeAppIconUrl      string `json:"we_app_icon_url"`     // 小程序图片
	} `json:"we_app_info"` // 拼多多福利券微信小程序信息
}

type OauthCmsUrlGenerate struct {
	CmsPromotionUrlGenerateResponse struct {
		Total   int64                        `json:"total"`
		UrlList []OauthCmsUrlGenerateUrlList `json:"url_list"` // 链接列表
	} `json:"cms_promotion_url_generate_response"`
}

// OauthCmsUrlGenerate 生成商城推广链接接口
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.oauth.cms.prom.url.generate
func (c *Client) OauthCmsUrlGenerate(ctx context.Context, notMustParams ...*gorequest.Params) (response OauthCmsUrlGenerate, err error) {

	// 参数
	params := NewParamsWithType("pdd.ddk.oauth.cms.prom.url.generate", notMustParams...)

	// 请求
	err = c.request(ctx, params, &response)
	return
}
