package ejiaofei

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type QueryTxProductResult struct {
	XMLName xml.Name `xml:"response"`
	Error   string   `xml:"error"` // 错误提示
}

// QueryTxProduct 可充值腾讯产品查询
func (app *App) QueryTxProduct() (body []byte, err error) {
	// 签名
	app.signStr = fmt.Sprintf("userid%vpwd%v", app.UserID, app.Pwd)
	// 请求
	body, err = app.request("http://api.ejiaofei.net:11140/queryTXproduct.do", map[string]interface{}{}, http.MethodGet)
	return body, err
}
