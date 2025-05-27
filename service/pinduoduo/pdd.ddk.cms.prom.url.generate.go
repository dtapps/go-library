package pinduoduo

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type CmsPromUrlGenerateUrlList struct {
	SingleUrlList struct {
		TzSchemaUrl          string `json:"tz_schema_url"`
		MobileUrl            string `json:"mobile_url"`
		SchemaUrl            string `json:"schema_url"`
		MobileShortUrl       string `json:"mobile_short_url"`
		WeAppWebViewUrl      string `json:"we_app_web_view_url"`
		Url                  string `json:"url"`
		ShortUrl             string `json:"short_url"`
		WeAppWebViewShortUrl string `json:"we_app_web_view_short_url"`
	} `json:"single_url_list"`
	MobileUrl string `json:"mobile_url"`
	Sign      string `json:"sign"`
	WeAppInfo struct {
		WeAppIconUrl      string `json:"we_app_icon_url"`
		UserName          string `json:"user_name"`
		PagePath          string `json:"page_path"`
		SourceDisplayName string `json:"source_display_name"`
		Title             string `json:"title"`
		AppId             string `json:"app_id"`
		Desc              string `json:"desc"`
	} `json:"we_app_info"`
	MobileShortUrl       string `json:"mobile_short_url"`
	WeAppWebViewUrl      string `json:"we_app_web_view_url"`
	Url                  string `json:"url"`
	ShortUrl             string `json:"short_url"`
	WeAppWebViewShortUrl string `json:"we_app_web_view_short_url"`
}

type CmsPromUrlGenerate struct {
	CmsPromotionUrlGenerateResponse struct {
		Total     int64                       `json:"total"`
		UrlList   []CmsPromUrlGenerateUrlList `json:"url_list"`
		RequestId string                      `json:"request_id"`
	} `json:"cms_promotion_url_generate_response"`
}

// CmsPromUrlGenerate 生成商城-频道推广链接
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.cms.prom.url.generate
func (c *Client) CmsPromUrlGenerate(ctx context.Context, notMustParams ...*gorequest.Params) (response CmsPromUrlGenerate, apiErr ApiError, err error) {

	// 参数
	params := NewParamsWithType("pdd.ddk.cms.prom.url.generate", notMustParams...)
	SetPidList(params, []string{c.GetPid()})

	// 请求
	err = c.requestAndErr(ctx, params, &response, &apiErr)
	return
}
