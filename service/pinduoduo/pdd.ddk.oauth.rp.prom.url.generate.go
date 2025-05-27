package pinduoduo

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type OauthRpPromUrlGenerateResourceList struct {
	Desc string `json:"desc"`
	Url  string `json:"url"`
}

type OauthRpPromUrlGenerateUrlList struct {
	MobileShortUrl           string `json:"mobile_short_url"`
	MobileUrl                string `json:"mobile_url"`
	MultiGroupMobileShortUrl string `json:"multi_group_mobile_short_url"`
	MultiGroupMobileUrl      string `json:"multi_group_mobile_url"`
	MultiGroupShortUrl       string `json:"multi_group_short_url"`
	MultiGroupUrl            string `json:"multi_group_url"`
	QqAppInfo                struct {
		AppId             string `json:"app_id"`
		BannerUrl         string `json:"banner_url"`
		Desc              string `json:"desc"`
		PagePath          string `json:"page_path"`
		QqAppIconUrl      string `json:"qq_app_icon_url"`
		SourceDisplayName string `json:"source_display_name"`
		Title             string `json:"title"`
		UserName          string `json:"user_name"`
	} `json:"qq_app_info"`
	SchemaUrl   string `json:"schema_url"`
	ShortUrl    string `json:"short_url"`
	TzSchemaUrl string `json:"tz_schema_url"`
	Url         string `json:"url"`
	WeAppInfo   struct {
		AppId             string `json:"app_id"`
		BannerUrl         string `json:"banner_url"`
		Desc              string `json:"desc"`
		PagePath          string `json:"page_path"`
		SourceDisplayName string `json:"source_display_name"`
		Title             string `json:"title"`
		UserName          string `json:"user_name"`
		WeAppIconUrl      string `json:"we_app_icon_url"`
	} `json:"we_app_info"`
}

type OauthRpPromUrlGenerate struct {
	RpPromotionUrlGenerateResponse struct {
		ResourceList []OauthRpPromUrlGenerateResourceList `json:"resource_list"`
		UrlList      []OauthRpPromUrlGenerateUrlList      `json:"url_list"`
	} `json:"rp_promotion_url_generate_response"`
}

// OauthRpPromUrlGenerate 生成营销工具推广链接
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.oauth.rp.prom.url.generate
func (c *Client) OauthRpPromUrlGenerate(ctx context.Context, notMustParams ...*gorequest.Params) (response OauthRpPromUrlGenerate, err error) {

	// 参数
	params := NewParamsWithType("pdd.ddk.oauth.rp.prom.url.generate", notMustParams...)

	// 请求
	err = c.request(ctx, params, &response)
	return
}
