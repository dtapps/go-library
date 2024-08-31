package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type GetVersionInfoResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	ExpInfo struct {
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

type GetVersionInfoResult struct {
	Result GetVersionInfoResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newGetVersionInfoResult(result GetVersionInfoResponse, body []byte, http gorequest.Response) *GetVersionInfoResult {
	return &GetVersionInfoResult{Result: result, Body: body, Http: http}
}

// GetVersionInfo 查询小程序版本信息
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/getVersionInfo.html
func (c *Client) GetVersionInfo(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*GetVersionInfoResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxa/getversioninfo")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response GetVersionInfoResponse
	request, err := c.request(ctx, span, "wxa/getversioninfo?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newGetVersionInfoResult(response, request.ResponseBody, request), err
}
