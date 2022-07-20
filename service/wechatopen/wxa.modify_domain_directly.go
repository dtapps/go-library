package wechatopen

import (
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaModifyDomainDirectlyResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type WxaModifyDomainDirectlyResult struct {
	Result WxaModifyDomainDirectlyResponse // 结果
	Body   []byte                          // 内容
	Http   gorequest.Response              // 请求
	Err    error                           // 错误
}

func newWxaModifyDomainDirectlyResult(result WxaModifyDomainDirectlyResponse, body []byte, http gorequest.Response, err error) *WxaModifyDomainDirectlyResult {
	return &WxaModifyDomainDirectlyResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaModifyDomainDirectly 快速设置小程序服务器域名
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/modify_domain_directly.html
func (c *Client) WxaModifyDomainDirectly(notMustParams ...gorequest.Params) *WxaModifyDomainDirectlyResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(fmt.Sprintf(apiUrl+"/wxa/modify_domain_directly?access_token=%s", c.GetAuthorizerAccessToken()), params, http.MethodPost)
	// 定义
	var response WxaModifyDomainDirectlyResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newWxaModifyDomainDirectlyResult(response, request.ResponseBody, request, err)
}

// ErrcodeInfo 错误描述
func (resp *WxaModifyDomainDirectlyResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 85015:
		return "该账号不是小程序账号"
	case 86100:
		return "该 URL 的协议头有误"
	case 45082:
		return "域名需要 icp 备案，否则无法添加"
	case 86101:
		return "不支持配置api.weixin.qq.com"
	case 85016:
		return "域名数量超限制"
	case 86102:
		return "每个月只能修改50次，超过域名修改次数限制"
	}
	return "系统繁忙"
}
