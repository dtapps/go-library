package wechatminiprogram

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaGenerateUrlLinkResponse struct {
	Errcode  int         `json:"errcode"`
	Errmsg   string      `json:"errmsg"`
	Openlink interface{} `json:"openlink"`
}

type WxaGenerateUrlLinkResult struct {
	Result WxaGenerateUrlLinkResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
}

func newWxaGenerateUrlLinkResult(result WxaGenerateUrlLinkResponse, body []byte, http gorequest.Response) *WxaGenerateUrlLinkResult {
	return &WxaGenerateUrlLinkResult{Result: result, Body: body, Http: http}
}

// WxaGenerateUrlLink 获取小程序 URL Link，适用于短信、邮件、网页、微信内等拉起小程序的业务场景。通过该接口，可以选择生成到期失效和永久有效的小程序链接，有数量限制，目前仅针对国内非个人主体的小程序开放
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html
func (c *Client) WxaGenerateUrlLink(ctx context.Context, notMustParams ...gorequest.Params) (*WxaGenerateUrlLinkResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/generate_urllink?access_token=%s", c.getAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return newWxaGenerateUrlLinkResult(WxaGenerateUrlLinkResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaGenerateUrlLinkResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaGenerateUrlLinkResult(response, request.ResponseBody, request), err
}
