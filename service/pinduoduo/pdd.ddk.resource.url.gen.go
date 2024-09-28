package pinduoduo

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
)

type ResourceUrlGenResponse struct {
	ResourceUrlResponse struct {
		MultiUrlList struct {
			ShortUrl string `json:"short_url"` // 频道推广短链接
			Url      string `json:"url"`       // 频道推广长链接
		} `json:"multi_url_list"` // 多人团链接
		Sign          string `json:"sign,omitempty"` // sign
		SingleUrlList struct {
			ShortUrl string `json:"short_url"` // 频道推广短链接
			Url      string `json:"url"`       // 频道推广长链接
		} `json:"single_url_list"` // 单人团链接
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
	} `json:"resource_url_response"`
}

type ResourceUrlGenError struct {
	ErrorResponse struct {
		ErrorMsg string `json:"error_msg"`
		SubMsg   string `json:"sub_msg"`
		SubCode  string `json:"sub_code"`
	} `json:"error_response"`
}

type ResourceUrlGenResult struct {
	Result ResourceUrlGenResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newResourceUrlGenResult(result ResourceUrlGenResponse, body []byte, http gorequest.Response) *ResourceUrlGenResult {
	return &ResourceUrlGenResult{Result: result, Body: body, Http: http}
}

// ResourceUrlGen 生成多多进宝频道推广
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.goods.pid.generate
func (c *Client) ResourceUrlGen(ctx context.Context, notMustParams ...gorequest.Params) (*ResourceUrlGenResult, ResourceUrlGenError, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "pdd.ddk.resource.url.gen")
	defer span.End()

	// 参数
	params := NewParamsWithType("pdd.ddk.resource.url.gen", notMustParams...)
	params.Set("pid", c.GetPid())

	// 请求
	var response ResourceUrlGenResponse
	request, err := c.request(ctx, span, params, &response)
	var responseError ResourceUrlGenError
	err = gojson.Unmarshal(request.ResponseBody, &responseError)
	return newResourceUrlGenResult(response, request.ResponseBody, request), responseError, err
}
