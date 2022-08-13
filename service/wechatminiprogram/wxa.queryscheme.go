package wechatminiprogram

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaQuerySchemeResponse struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	SchemeInfo struct {
		Appid      string `json:"appid"`
		Path       string `json:"path"`
		Query      string `json:"query"`
		CreateTime int    `json:"create_time"`
		ExpireTime int    `json:"expire_time"`
		EnvVersion string `json:"env_version"`
	} `json:"scheme_info"`
	SchemeQuota struct {
		LongTimeUsed  int `json:"long_time_used"`
		LongTimeLimit int `json:"long_time_limit"`
	} `json:"scheme_quota"`
}

type WxaQuerySchemeResult struct {
	Result WxaQuerySchemeResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
	Err    error                  // 错误
}

func newWxaQuerySchemeResult(result WxaQuerySchemeResponse, body []byte, http gorequest.Response, err error) *WxaQuerySchemeResult {
	return &WxaQuerySchemeResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaQueryScheme 查询小程序 scheme 码，及长期有效 quota
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.query.html
func (c *Client) WxaQueryScheme(ctx context.Context, notMustParams ...gorequest.Params) *WxaQuerySchemeResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/queryscheme?access_token=%s", c.getAccessToken(ctx)), params, http.MethodPost)
	// 定义
	var response WxaQuerySchemeResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newWxaQuerySchemeResult(response, request.ResponseBody, request, err)
}
