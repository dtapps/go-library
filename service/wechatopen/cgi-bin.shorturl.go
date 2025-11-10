package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type CgiBinShortUrlResponse struct {
	APIResponse        // 错误
	ShortUrl    string `json:"short_url"` // 	短链接。
}

// CgiBinShortUrl 将二维码长链接转成短链接
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/qrcode/shorturl.html
func (c *Client) CgiBinShortUrl(ctx context.Context, longUrl string, notMustParams ...*gorequest.Params) (response CgiBinShortUrlResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("action", "long2short") // 此处填long2short，代表长链接转短链接
	params.Set("long_url", longUrl)    // 需要转换的长链接，支持http://、https://、weixin://wxpay 格式的url

	// 请求
	err = c.request(ctx, "cgi-bin/shorturl?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}
