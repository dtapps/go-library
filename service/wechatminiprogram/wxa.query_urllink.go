package wechatminiprogram

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
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
}

func newWxaQueryUrlLinkResult(result WxaQueryUrlLinkResponse, body []byte, http gorequest.Response) *WxaQueryUrlLinkResult {
	return &WxaQueryUrlLinkResult{Result: result, Body: body, Http: http}
}

// WxaQueryUrlLink 查询小程序 url_link 配置，及长期有效 quota
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.query.html
func (c *Client) WxaQueryUrlLink(ctx context.Context, notMustParams ...gorequest.Params) (*WxaQueryUrlLinkResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/query_urllink?access_token=%s", c.getAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return newWxaQueryUrlLinkResult(WxaQueryUrlLinkResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaQueryUrlLinkResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaQueryUrlLinkResult(response, request.ResponseBody, request), err
}
