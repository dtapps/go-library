package wechatopen

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
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
func (c *Client) CgiBinShortUrl(ctx context.Context, longUrl string) (*CgiBinShortUrlResult, error) {
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
	params["action"] = "long2short" // 此处填long2short，代表长链接转短链接
	params["long_url"] = longUrl    // 需要转换的长链接，支持http://、https://、weixin://wxpay 格式的url
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/cgi-bin/shorturl?access_token=%s", c.GetAuthorizerAccessToken(ctx)), params, http.MethodPost)
	// 定义
	var response CgiBinShortUrlResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newCgiBinShortUrlResult(response, request.ResponseBody, request), nil
}
