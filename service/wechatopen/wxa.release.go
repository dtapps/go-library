package wechatopen

import (
	"context"
	"encoding/json"
	"fmt"
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
func (c *Client) WxaRelease(ctx context.Context) (*WxaReleaseResult, error) {
	// 检查
	err := c.checkComponentIsConfig()
	if err != nil {
		return nil, err
	}
	err = c.checkAuthorizerIsConfig()
	if err != nil {
		return nil, err
	}
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/release?access_token=%s", c.GetAuthorizerAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response WxaReleaseResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newWxaReleaseResult(response, request.ResponseBody, request), nil
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
