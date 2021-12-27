package ejiaofei

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type MoneyJkUserResult struct {
	XMLName   xml.Name `xml:"response"`
	LastMoney float64  `xml:"lastMoney"` // 用户余额
	Tag       int      `xml:"tag"`       // 用户状态（0正常 1暂停）
	Error     int      `xml:"error"`     // 错误提示
}

// MoneyJkUser 用户余额查询
func (app *App) MoneyJkUser() (body []byte, err error) {
	// 签名
	app.signStr = fmt.Sprintf("userid%vpwd%v", app.UserID, app.Pwd)
	// 请求
	body, err = app.request("http://api.ejiaofei.net:11140/money_jkuser.do", map[string]interface{}{}, http.MethodGet)
	return body, err
}
