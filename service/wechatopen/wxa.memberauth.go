package wechatopen

import (
	"context"
	"encoding/json"
	"fmt"
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
func (c *Client) WxaMemberAuth(ctx context.Context) (*WxaMemberAuthResult, error) {
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
	params["action"] = "get_experiencer"
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/memberauth?access_token=%s", c.GetAuthorizerAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response WxaMemberAuthResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newWxaMemberAuthResult(response, request.ResponseBody, request), nil
}
