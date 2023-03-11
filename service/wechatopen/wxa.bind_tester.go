package wechatopen

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaBindTesterResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
	Userstr string `json:"userstr"` // 人员对应的唯一字符串
}

type WxaBindTesterResult struct {
	Result WxaBindTesterResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newWxaBindTesterResult(result WxaBindTesterResponse, body []byte, http gorequest.Response) *WxaBindTesterResult {
	return &WxaBindTesterResult{Result: result, Body: body, Http: http}
}

// WxaBindTester 绑定微信用户为体验者
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/Admin.html
func (c *Client) WxaBindTester(ctx context.Context, wechatid string, notMustParams ...gorequest.Params) (*WxaBindTesterResult, error) {
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
	params := gorequest.NewParamsWith(notMustParams...)
	params["wechatid"] = wechatid
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/bind_tester?access_token=%s", c.GetAuthorizerAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response WxaBindTesterResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newWxaBindTesterResult(response, request.ResponseBody, request), nil
}

// ErrcodeInfo 错误描述
func (resp *WxaBindTesterResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 85001:
		return "微信号不存在或微信号设置为不可搜索"
	case 85002:
		return "小程序绑定的体验者数量达到上限"
	case 85003:
		return "微信号绑定的小程序体验者达到上限"
	case 85004:
		return "微信号已经绑定"
	}
	return "系统繁忙"
}
