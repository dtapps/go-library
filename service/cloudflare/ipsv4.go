package cloudflare

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type IpsV4Response struct{}

type IpsV4Result struct {
	Result IpsV4Response      // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
	Err    error              // 错误
}

func newIpsV4Result(result IpsV4Response, body []byte, http gorequest.Response, err error) *IpsV4Result {
	return &IpsV4Result{Result: result, Body: body, Http: http, Err: err}
}

// IpsV4 ipv4
// https://www.cloudflare.com/ips-v4
func (c *Client) IpsV4() *IpsV4Result {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(apiUrl+"/ips-v4", params, http.MethodPost)
	// 定义
	var response IpsV4Response
	err = json.Unmarshal(request.ResponseBody, &response)
	return newIpsV4Result(response, request.ResponseBody, request, err)
}
