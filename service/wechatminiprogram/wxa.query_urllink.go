package wechatminiprogram

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WxaQueryUrlLinkResponse struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	UrlLinkInfo struct {
		Appid      string `json:"appid"`
		Path       string `json:"path"`
		Query      string `json:"query"`
		CreateTime int    `json:"create_time"`
		ExpireTime int    `json:"expire_time"`
		EnvVersion string `json:"env_version"`
		CloudBase  struct {
			Env           string `json:"env"`
			Doamin        string `json:"doamin"`
			Path          string `json:"path"`
			Query         string `json:"query"`
			ResourceAppid string `json:"resource_appid"`
		} `json:"cloud_base"`
	} `json:"url_link_info"`
	UrlLinkQuota struct {
		LongTimeUsed  int `json:"long_time_used"`
		LongTimeLimit int `json:"long_time_limit"`
	} `json:"url_link_quota"`
}

type WxaQueryUrlLinkResult struct {
	Result WxaQueryUrlLinkResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
	Err    error                   // 错误
}

func newWxaQueryUrlLinkResult(result WxaQueryUrlLinkResponse, body []byte, http gorequest.Response, err error) *WxaQueryUrlLinkResult {
	return &WxaQueryUrlLinkResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaQueryUrlLink 查询小程序 url_link 配置，及长期有效 quota
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.query.html
func (c *Client) WxaQueryUrlLink(urlLink string) *WxaQueryUrlLinkResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("url_link", urlLink)
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(fmt.Sprintf(apiUrl+"/wxa/query_urllink?access_token=%s", c.getAccessToken()), params, http.MethodPost)
	// 定义
	var response WxaQueryUrlLinkResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newWxaQueryUrlLinkResult(response, request.ResponseBody, request, err)
}
