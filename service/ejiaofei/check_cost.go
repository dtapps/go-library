package ejiaofei

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type CheckCostResult struct {
	XMLName xml.Name `xml:"response"`
	UserID  string   `xml:"userid"`  // 用户账号
	OrderID string   `xml:"orderid"` // 用户提交订单号
	Face    float64  `xml:"face"`    // 官方价格
	Price   float64  `xml:"price"`   // 用户成本价
	Error   int      `xml:"error"`   // 错误提示
}

// CheckCost 会员订单成本价查询接口
func (app *App) CheckCost(orderID string) (body []byte, err error) {
	// 参数
	param := NewParams()
	param.Set("orderid", orderID)
	params := app.NewParamsWith(param)
	// 签名
	app.signStr = fmt.Sprintf("userid%vpwd%vorderid%v", app.UserID, app.Pwd, orderID)
	// 请求
	body, err = app.request("http://api.ejiaofei.net:11140/checkCost.do", params, http.MethodGet)
	return body, err
}
