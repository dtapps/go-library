package wechatopen

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type GetCallBackIpResponse struct {
	IpList []string `json:"ip_list"`
}

type GetCallBackIpResult struct {
	Result GetCallBackIpResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
	Err    error                 // 错误
}

func NewGetCallBackIpResult(result GetCallBackIpResponse, body []byte, http gorequest.Response, err error) *GetCallBackIpResult {
	return &GetCallBackIpResult{Result: result, Body: body, Http: http, Err: err}
}

// CgiBinGetApiDomainIp 获取微信服务器IP地址
// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
func (c *Client) CgiBinGetApiDomainIp(ctx context.Context, componentAccessToken string) *GetCallBackIpResult {
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/cgi-bin/get_api_domain_ip?access_token=%s", componentAccessToken), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response GetCallBackIpResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewGetCallBackIpResult(response, request.ResponseBody, request, err)
}
