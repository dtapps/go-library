package pinduoduo

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
)

type RpPromUrlGenerateResponse struct {
	RpPromotionUrlGenerateResponse struct {
		ResourceList []struct {
			Desc string `json:"desc"` // 活动描述
			Url  string `json:"url"`  // 活动地址
		} `json:"resource_list"`
		UrlList []struct {
			MobileShortUrl           string `json:"mobile_short_url"`             // 推广移动短链接，对应出参mobile_url的短链接，与mobile_url功能一致。
			MobileUrl                string `json:"mobile_url"`                   // 推广移动链接，用户安装拼多多APP的情况下会唤起APP，否则唤起H5页面
			MultiGroupMobileShortUrl string `json:"multi_group_mobile_short_url"` // 推广多人团移动短链接
			MultiGroupMobileUrl      string `json:"multi_group_mobile_url"`       // 推广多人团移动链接，用户安装拼多多APP的情况下会唤起APP，否则唤起H5页面
			MultiGroupShortUrl       string `json:"multi_group_short_url"`        // 推广多人团短链接
			MultiGroupUrl            string `json:"multi_group_url"`              // 推广多人团链接，唤起H5页面
			QqAppInfo                struct {
				AppId             string `json:"app_id"`              // 拼多多小程序id
				BannerUrl         string `json:"banner_url"`          // Banner图
				Desc              string `json:"desc"`                // 描述
				PagePath          string `json:"page_path"`           // 小程序path值
				QqAppIconUrl      string `json:"qq_app_icon_url"`     // 小程序icon
				SourceDisplayName string `json:"source_display_name"` // 来源名
				Title             string `json:"title"`               // 小程序标题
				UserName          string `json:"user_name"`           // 用户名
			} `json:"qq_app_info"` // qq小程序信息
			SchemaUrl string `json:"schema_url"` // schema链接，用户安装拼多多APP的情况下会唤起APP（需客户端支持schema跳转协议）
			ShortUrl  string `json:"short_url"`  // 推广短链接，对应出参url的短链接，与url功能一致
			Url       string `json:"url"`        // 普通推广长链接，唤起H5页面
			WeAppInfo struct {
				AppId             string `json:"app_id"`              // 小程序id
				BannerUrl         string `json:"banner_url"`          // Banner图
				Desc              string `json:"desc"`                // 描述
				PagePath          string `json:"page_path"`           // 小程序path值
				SourceDisplayName string `json:"source_display_name"` // 来源名
				Title             string `json:"title"`               // 小程序标题
				UserName          string `json:"user_name"`           // 用户名
				WeAppIconUrl      string `json:"we_app_icon_url"`     // 小程序icon
			} `json:"we_app_info"` // 拼多多福利券微信小程序信息
		} `json:"url_list"`
	} `json:"rp_promotion_url_generate_response"`
}
type RpPromUrlGenerateError struct {
	ErrorResponse struct {
		ErrorMsg string `json:"error_msg"`
		SubMsg   string `json:"sub_msg"`
		SubCode  string `json:"sub_code"`
	} `json:"error_response"`
}

type RpPromUrlGenerateResult struct {
	Result RpPromUrlGenerateResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newRpPromUrlGenerateResult(result RpPromUrlGenerateResponse, body []byte, http gorequest.Response) *RpPromUrlGenerateResult {
	return &RpPromUrlGenerateResult{Result: result, Body: body, Http: http}
}

// RpPromUrlGenerate 生成营销工具推广链接
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.rp.prom.url.generate
func (c *Client) RpPromUrlGenerate(ctx context.Context, notMustParams ...gorequest.Params) (*RpPromUrlGenerateResult, RpPromUrlGenerateError, error) {

	// 参数
	params := NewParamsWithType("pdd.ddk.rp.prom.url.generate", notMustParams...)
	params.Set("p_id_list", []string{c.GetPid()})

	// 请求
	var response RpPromUrlGenerateResponse
	request, err := c.request(ctx, params, &response)
	var responseError RpPromUrlGenerateError
	err = gojson.Unmarshal(request.ResponseBody, &responseError)
	return newRpPromUrlGenerateResult(response, request.ResponseBody, request), responseError, err
}
