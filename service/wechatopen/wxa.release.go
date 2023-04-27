package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaReleaseResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type WxaReleaseResult struct {
	Result WxaReleaseResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newWxaReleaseResult(result WxaReleaseResponse, body []byte, http gorequest.Response) *WxaReleaseResult {
	return &WxaReleaseResult{Result: result, Body: body, Http: http}
}

// WxaRelease 发布已通过审核的小程序
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/release.html
func (c *Client) WxaRelease(ctx context.Context, notMustParams ...gorequest.Params) (*WxaReleaseResult, error) {
	// 检查
	if err := c.checkAuthorizerConfig(ctx); err != nil {
		return newWxaReleaseResult(WxaReleaseResponse{}, []byte{}, gorequest.Response{}), err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/release?access_token="+c.GetAuthorizerAccessToken(ctx), params, http.MethodPost)
	if err != nil {
		return newWxaReleaseResult(WxaReleaseResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaReleaseResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaReleaseResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaReleaseResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 85019:
		return "没有审核版本"
	case 85020:
		return "审核状态未满足发布"
	}
	return "系统繁忙"
}
