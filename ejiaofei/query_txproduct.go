package ejiaofei

import (
	"encoding/xml"
	"fmt"
	"go.dtapp.net/library/gorequest"
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

func NewQueryTxProductResult(result QueryTxProductResponse, body []byte, http gorequest.Response, err error) *QueryTxProductResult {
	return &QueryTxProductResult{Result: result, Body: body, Http: http, Err: err}
}

// QueryTxProduct 可充值腾讯产品查询
func (app *App) QueryTxProduct() *QueryTxProductResult {
	// 签名
	app.signStr = fmt.Sprintf("userid%vpwd%v", app.userId, app.pwd)
	// 请求
	request, err := app.request("http://api.ejiaofei.net:11140/queryTXproduct.do", map[string]interface{}{}, http.MethodGet)
	// 定义
	var response QueryTxProductResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return NewQueryTxProductResult(response, request.ResponseBody, request, err)
}
