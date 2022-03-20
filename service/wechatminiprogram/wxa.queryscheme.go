package wechatminiprogram

import (
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gohttp"
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
	Http   gohttp.Response        // 请求
	Err    error                  // 错误
}

func NewWxaQuerySchemeResult(result WxaQuerySchemeResponse, body []byte, http gohttp.Response, err error) *WxaQuerySchemeResult {
	return &WxaQuerySchemeResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaQueryScheme 查询小程序 scheme 码，及长期有效 quota
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.query.html
func (app *App) WxaQueryScheme(notMustParams ...Params) *WxaQuerySchemeResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/queryscheme?access_token=%s", app.AccessToken), params, http.MethodPost)
	// 定义
	var response WxaQuerySchemeResponse
	err = json.Unmarshal(request.Body, &response)
	return NewWxaQuerySchemeResult(response, request.Body, request, err)
}
