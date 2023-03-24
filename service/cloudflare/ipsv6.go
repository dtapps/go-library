package cloudflare

import (
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type IpsV6Response struct{}

type IpsV6Result struct {
	Result IpsV6Response      // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
	Err    error              // 错误
}

func newIpsV6Result(result IpsV6Response, body []byte, http gorequest.Response, err error) *IpsV6Result {
	return &IpsV6Result{Result: result, Body: body, Http: http, Err: err}
}

// IpsV6 ipv6
// https://www.cloudflare.com/ips-v6
func (c *Client) IpsV6() *IpsV6Result {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(apiUrl+"/ips-v6", params, http.MethodPost)
	// 定义
	var response IpsV6Response
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newIpsV6Result(response, request.ResponseBody, request, err)
}
