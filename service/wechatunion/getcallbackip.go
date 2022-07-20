package wechatunion

import (
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
)

type GetCallBackIpResponse struct {
	IpList []string `json:"ip_list"`
}

type GetCallBackIpResult struct {
	Result GetCallBackIpResponse // 结果
	Byte   []byte                // 内容
	Http   gorequest.Response    // 请求
	Err    error                 // 错误
}

func newGetCallBackIpResult(result GetCallBackIpResponse, byte []byte, http gorequest.Response, err error) *GetCallBackIpResult {
	return &GetCallBackIpResult{Result: result, Byte: byte, Http: http, Err: err}
}

// GetCallBackIp 获取微信callback IP地址
// callback IP即微信调用开发者服务器所使用的出口IP。
// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html#2.%20%E8%8E%B7%E5%8F%96%E5%BE%AE%E4%BF%A1callback%20IP%E5%9C%B0%E5%9D%80
func (c *Client) GetCallBackIp() *GetCallBackIpResult {
	// 请求
	request, err := c.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=%s", c.getAccessToken()), map[string]interface{}{}, "GET")
	// 定义
	var response GetCallBackIpResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newGetCallBackIpResult(response, request.ResponseBody, request, err)
}
