package wechatopen

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetCallBackIpResponse struct {
	IpList []string `json:"ip_list"`
}

type GetCallBackIpResult struct {
	Result GetCallBackIpResponse // 结果
	Body   []byte                // 内容
	Err    error                 // 错误
}

func NewGetCallBackIpResult(result GetCallBackIpResponse, body []byte, err error) *GetCallBackIpResult {
	return &GetCallBackIpResult{Result: result, Body: body, Err: err}
}

// CgiBinGetApiDomainIp 获取微信服务器IP地址
// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
func (app *App) CgiBinGetApiDomainIp(componentAccessToken string) *GetCallBackIpResult {
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/get_api_domain_ip?access_token=%s", componentAccessToken), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response GetCallBackIpResponse
	err = json.Unmarshal(body, &response)
	return NewGetCallBackIpResult(response, body, err)
}