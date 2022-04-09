package ejiaofei

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type QueryTxProductResponse struct {
	XMLName xml.Name `xml:"response"`
	Error   string   `xml:"error"` // 错误提示
}

type QueryTxProductResult struct {
	Result QueryTxProductResponse // 结果
	Body   []byte                 // 内容
	Err    error                  // 错误
}

func NewQueryTxProductResult(result QueryTxProductResponse, body []byte, err error) *QueryTxProductResult {
	return &QueryTxProductResult{Result: result, Body: body, Err: err}
}

// QueryTxProduct 可充值腾讯产品查询
func (app *App) QueryTxProduct() *QueryTxProductResult {
	// 签名
	app.signStr = fmt.Sprintf("userid%vpwd%v", app.UserID, app.Pwd)
	// 请求
	body, err := app.request("http://api.ejiaofei.net:11140/queryTXproduct.do", map[string]interface{}{}, http.MethodGet)
	// 定义
	var response QueryTxProductResponse
	err = xml.Unmarshal(body, &response)
	return NewQueryTxProductResult(response, body, err)
}
