package pinduoduo

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
)

type CmsPromUrlGenerateResponse struct {
	CmsPromotionUrlGenerateResponse struct {
		Total   int `json:"total"`
		UrlList []struct {
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
		} `json:"url_list"`
		RequestId string `json:"request_id"`
	} `json:"cms_promotion_url_generate_response"`
}

type CmsPromUrlGenerateError struct {
	ErrorResponse struct {
		ErrorMsg string `json:"error_msg"`
		SubMsg   string `json:"sub_msg"`
		SubCode  string `json:"sub_code"`
	} `json:"error_response"`
}

type CmsPromUrlGenerateResult struct {
	Result CmsPromUrlGenerateResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
	Err    error                      // 错误
	Error  CmsPromUrlGenerateError    // 错误结果
}

func newCmsPromUrlGenerateResult(result CmsPromUrlGenerateResponse, body []byte, http gorequest.Response, err error, error CmsPromUrlGenerateError) *CmsPromUrlGenerateResult {
	return &CmsPromUrlGenerateResult{Result: result, Body: body, Http: http, Err: err, Error: error}
}

// CmsPromUrlGenerate 生成商城-频道推广链接
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.cms.prom.url.generate
func (c *Client) CmsPromUrlGenerate(ctx context.Context, notMustParams ...Params) *CmsPromUrlGenerateResult {
	// 参数
	params := NewParamsWithType("pdd.ddk.cms.prom.url.generate", notMustParams...)
	params.Set("p_id_list", []string{c.GetPid()})
	// 请求
	request, err := c.request(ctx, params)
	// 定义
	var response CmsPromUrlGenerateResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	var responseError CmsPromUrlGenerateError
	err = json.Unmarshal(request.ResponseBody, &responseError)
	return newCmsPromUrlGenerateResult(response, request.ResponseBody, request, err, responseError)
}
