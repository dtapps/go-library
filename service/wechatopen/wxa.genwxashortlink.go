package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type GenerateShortLinkResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
	Link    string `json:"link"`    // 生成的小程序 Short Link
}

type GenerateShortLinkResult struct {
	Result GenerateShortLinkResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newGenerateShortLinkResult(result GenerateShortLinkResponse, body []byte, http gorequest.Response) *GenerateShortLinkResult {
	return &GenerateShortLinkResult{Result: result, Body: body, Http: http}
}

// GenerateShortLink 获取ShortLink
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/short-link/generateShortLink.html
func (c *Client) GenerateShortLink(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (*GenerateShortLinkResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response GenerateShortLinkResponse
	request, err := c.request(ctx, "wxa/genwxashortlink?access_token="+authorizerAccessToken, params, http.MethodPost, &response)

	return newGenerateShortLinkResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *GenerateShortLinkResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 40001:
		return "获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的公众号调用接口"
	case 40066:
		return "url不存在，即，已发布小程序没有对应url"
	case 40225:
		return "无效的页面标题"
	case 85400:
		return "长期有效Scheme或short link达到生成上限10万，不可再生成。"
	case 45009:
		return "单天生成Short Link数量超过上限100万"
	case 43104:
		return "没有调用权限，目前只开放给电商类目（具体包含以下一级类目：电商平台、商家自营、跨境电商）"
	default:
		return resp.Result.Errmsg
	}
}
