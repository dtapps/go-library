package ejiaofei

import (
	"encoding/xml"
	"fmt"
	"go.dtapp.net/library/gorequest"
	"net/http"
)

type CheckCostResponse struct {
	XMLName xml.Name `xml:"response"`
	UserID  string   `xml:"userid"`  // 用户账号
	OrderID string   `xml:"orderid"` // 用户提交订单号
	Face    float64  `xml:"face"`    // 官方价格
	Price   float64  `xml:"price"`   // 用户成本价
	Error   int      `xml:"error"`   // 错误提示
}

type CheckCostResult struct {
	Result CheckCostResponse  // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
	Err    error              // 错误
}

func NewCheckCostResult(result CheckCostResponse, body []byte, http gorequest.Response, err error) *CheckCostResult {
	return &CheckCostResult{Result: result, Body: body, Http: http, Err: err}
}

// CheckCost 会员订单成本价查询接口
func (app *App) CheckCost(orderId string) *CheckCostResult {
	// 参数
	param := NewParams()
	param.Set("orderid", orderId)
	params := app.NewParamsWith(param)
	// 签名
	app.signStr = fmt.Sprintf("userid%vpwd%vorderid%v", app.userId, app.pwd, orderId)
	// 请求
	request, err := app.request("http://api.ejiaofei.net:11140/checkCost.do", params, http.MethodGet)
	// 定义
	var response CheckCostResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return NewCheckCostResult(response, request.ResponseBody, request, err)
}
