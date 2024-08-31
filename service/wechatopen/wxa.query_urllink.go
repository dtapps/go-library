package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type QueryUrlLinkResponse struct {
	Errcode     int    `json:"errcode"` // 错误码
	Errmsg      string `json:"errmsg"`  // 错误信息
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

type QueryUrlLinkResult struct {
	Result QueryUrlLinkResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
}

func newQueryUrlLinkResult(result QueryUrlLinkResponse, body []byte, http gorequest.Response) *QueryUrlLinkResult {
	return &QueryUrlLinkResult{Result: result, Body: body, Http: http}
}

// QueryUrlLink 查询加密URLLink
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-link/queryUrlLink.html
func (c *Client) QueryUrlLink(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*QueryUrlLinkResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxa/query_urllink")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response QueryUrlLinkResponse
	request, err := c.request(ctx, span, "wxa/query_urllink?access_token="+authorizerAccessToken, params, http.MethodPost, &response)

	return newQueryUrlLinkResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *QueryUrlLinkResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 40097:
		return "参数错误"
	case 85403:
		return "scheme/url link不存在"
	default:
		return resp.Result.Errmsg
	}
}
