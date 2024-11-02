package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type QuerySchemeResponse struct {
	Errcode    int    `json:"errcode"` // 错误码
	Errmsg     string `json:"errmsg"`  // 错误信息
	SchemeInfo struct {
		Appid      string `json:"appid"`       // 小程序 appid
		Path       string `json:"path"`        // 小程序页面路径
		Query      string `json:"query"`       // 小程序页面query
		CreateTime int    `json:"create_time"` // 创建时间，为 Unix 时间戳
		ExpireTime int    `json:"expire_time"` // 到期失效时间，为 Unix 时间戳，0 表示永久生效
		EnvVersion string `json:"env_version"` // 要打开的小程序版本。正式版为"release"，体验版为"trial"，开发版为"develop"
	} `json:"scheme_info"` // scheme 信息
	QuotaInfo struct {
		RemainVisitQuota string `json:"remain_visit_quota"` // 	URL Scheme（加密+明文）/加密 URL Link 单天剩余访问次数
	} `json:"quota_info"`
}

type QuerySchemeResult struct {
	Result QuerySchemeResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newQuerySchemeResult(result QuerySchemeResponse, body []byte, http gorequest.Response) *QuerySchemeResult {
	return &QuerySchemeResult{Result: result, Body: body, Http: http}
}

// QueryScheme 查询scheme码
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-scheme/queryScheme.html
func (c *Client) QueryScheme(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*QuerySchemeResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response QuerySchemeResponse
	request, err := c.request(ctx, "wxa/queryscheme?access_token="+authorizerAccessToken, params, http.MethodPost, &response)

	return newQuerySchemeResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *QuerySchemeResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 40097:
		return "参数错误"
	case 85403:
		return "scheme/url link不存在"
	default:
		return resp.Result.Errmsg
	}
}
