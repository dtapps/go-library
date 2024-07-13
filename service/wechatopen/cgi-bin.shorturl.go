package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type CgiBinShortUrlResponse struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	ShortUrl string `json:"short_url"` // 	短链接。
}

type CgiBinShortUrlResult struct {
	Result CgiBinShortUrlResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newCgiBinShortUrlResult(result CgiBinShortUrlResponse, body []byte, http gorequest.Response) *CgiBinShortUrlResult {
	return &CgiBinShortUrlResult{Result: result, Body: body, Http: http}
}

// CgiBinShortUrl 将二维码长链接转成短链接
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/qrcode/shorturl.html
func (c *Client) CgiBinShortUrl(ctx context.Context, authorizerAccessToken, longUrl string, notMustParams ...gorequest.Params) (*CgiBinShortUrlResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "cgi-bin/shorturl")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("action", "long2short") // 此处填long2short，代表长链接转短链接
	params.Set("long_url", longUrl)    // 需要转换的长链接，支持http://、https://、weixin://wxpay 格式的url

	// 请求
	var response CgiBinShortUrlResponse
	request, err := c.request(ctx, span, "cgi-bin/shorturl?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newCgiBinShortUrlResult(response, request.ResponseBody, request), err
}
