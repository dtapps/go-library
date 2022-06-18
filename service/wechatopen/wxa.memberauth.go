package wechatopen

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
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
	Err    error                 // 错误
}

func NewWxaMemberAuthResult(result WxaMemberAuthResponse, body []byte, http gorequest.Response, err error) *WxaMemberAuthResult {
	return &WxaMemberAuthResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaMemberAuth 获取体验者列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/memberauth.html
func (c *Client) WxaMemberAuth() *WxaMemberAuthResult {
	accessToken := c.GetAuthorizerAccessToken()
	if accessToken == "" {
		return NewWxaMemberAuthResult(WxaMemberAuthResponse{}, nil, gorequest.Response{}, errors.New("访问令牌为空"))
	}
	// 参数
	params := NewParams()
	params["action"] = "get_experiencer"
	// 请求
	request, err := c.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/memberauth?access_token=%s", accessToken), params, http.MethodPost)
	// 定义
	var response WxaMemberAuthResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewWxaMemberAuthResult(response, request.ResponseBody, request, err)
}
