package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaMemberAuthResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
	Members []struct {
		Userstr string `json:"userstr"` // 人员对应的唯一字符串
	} `json:"members"` // 人员信息列表
}

type WxaMemberAuthResult struct {
	Result WxaMemberAuthResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newWxaMemberAuthResult(result WxaMemberAuthResponse, body []byte, http gorequest.Response) *WxaMemberAuthResult {
	return &WxaMemberAuthResult{Result: result, Body: body, Http: http}
}

// WxaMemberAuth 获取体验者列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/memberauth.html
func (c *Client) WxaMemberAuth(ctx context.Context, notMustParams ...gorequest.Params) (*WxaMemberAuthResult, error) {
	// 检查
	if err := c.checkAuthorizerConfig(ctx); err != nil {
		return newWxaMemberAuthResult(WxaMemberAuthResponse{}, []byte{}, gorequest.Response{}), err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("action", "get_experiencer")
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/memberauth?access_token="+c.GetAuthorizerAccessToken(ctx), params, http.MethodPost)
	if err != nil {
		return newWxaMemberAuthResult(WxaMemberAuthResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaMemberAuthResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaMemberAuthResult(response, request.ResponseBody, request), err
}
