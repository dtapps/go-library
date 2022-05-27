package dingdanxia

import (
	"encoding/json"
	"net/http"
)

// JdOrderDetails2 【官方不维护】 京东联盟订单行查询
// https://www.dingdanxia.com/doc/180/94
func (app *App) JdOrderDetails2(notMustParams ...Params) *JdJyOrderDetailsResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request("https://api.tbk.dingdanxia.com/jd/order_details2", params, http.MethodPost)
	// 定义
	var response JdJyOrderDetailsResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewJdJyOrderDetailsResult(response, request.ResponseBody, request, err)
}
