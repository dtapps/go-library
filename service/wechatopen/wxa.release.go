package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type ReleaseResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type ReleaseResult struct {
	Result ReleaseResponse    // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newReleaseResult(result ReleaseResponse, body []byte, http gorequest.Response) *ReleaseResult {
	return &ReleaseResult{Result: result, Body: body, Http: http}
}

// Release 发布已通过审核的小程序
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/release.html
func (c *Client) Release(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (*ReleaseResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response ReleaseResponse
	request, err := c.request(ctx, "wxa/release?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newReleaseResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *ReleaseResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 85019:
		return "没有审核版本"
	case 85020:
		return "审核状态未满足发布"
	default:
		return resp.Result.Errmsg
	}
}
