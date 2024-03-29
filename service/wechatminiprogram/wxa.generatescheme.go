package wechatminiprogram

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaGenerateSchemeResponse struct {
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
	UrlLink interface{} `json:"url_link"`
}

type WxaGenerateSchemeResult struct {
	Result WxaGenerateSchemeResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newWxaGenerateSchemeResult(result WxaGenerateSchemeResponse, body []byte, http gorequest.Response) *WxaGenerateSchemeResult {
	return &WxaGenerateSchemeResult{Result: result, Body: body, Http: http}
}

// WxaGenerateScheme 获取小程序 scheme 码，适用于短信、邮件、外部网页、微信内等拉起小程序的业务场景。通过该接口，可以选择生成到期失效和永久有效的小程序码，有数量限制，目前仅针对国内非个人主体的小程序开放
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.generate.html
func (c *Client) WxaGenerateScheme(ctx context.Context, notMustParams ...gorequest.Params) (*WxaGenerateSchemeResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/generatescheme?access_token=%s", c.getAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return newWxaGenerateSchemeResult(WxaGenerateSchemeResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaGenerateSchemeResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaGenerateSchemeResult(response, request.ResponseBody, request), err
}
