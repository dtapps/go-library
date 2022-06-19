package ejiaofei

import (
	"encoding/xml"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type QueryTxProductResponse struct {
	XMLName xml.Name `xml:"response"`
	Error   string   `xml:"error"` // 错误提示
}

type QueryTxProductResult struct {
	Result QueryTxProductResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
	Err    error                  // 错误
}

func newQueryTxProductResult(result QueryTxProductResponse, body []byte, http gorequest.Response, err error) *QueryTxProductResult {
	return &QueryTxProductResult{Result: result, Body: body, Http: http, Err: err}
}

// QueryTxProduct 可充值腾讯产品查询
func (c *Client) QueryTxProduct() *QueryTxProductResult {
	// 签名
	c.signStr = fmt.Sprintf("userid%vpwd%v", c.getUserId(), c.getPwd())
	// 请求
	request, err := c.request(apiUrl+"/queryTXproduct.do", map[string]interface{}{}, http.MethodGet)
	// 定义
	var response QueryTxProductResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return newQueryTxProductResult(response, request.ResponseBody, request, err)
}
