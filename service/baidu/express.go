package baidu

import (
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type ExpressResponse struct{}

type ExpressResult struct {
	Result ExpressResponse    // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
	Err    error              // 错误
}

func newExpressResult(result ExpressResponse, body []byte, http gorequest.Response, err error) *ExpressResult {
	return &ExpressResult{Result: result, Body: body, Http: http, Err: err}
}

// Express ipv1
// https://www.cloudflare.com/ips-v4
func (c *Client) Express() *ExpressResult {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request("https://m.baidu.com/s?word=快递查询&ts={$ts}&t_kt=0&ie=utf-8&rsv_iqid=&rsv_t=&sa=&rsv_pq=&rsv_sug4=&tj=1&inputT={$input}&sugid=&ss=", params, http.MethodPost)
	// 定义
	var response ExpressResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newExpressResult(response, request.ResponseBody, request, err)
}
