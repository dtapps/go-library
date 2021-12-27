package wikeyun

import (
	"encoding/json"
	"errors"
)

// RechargeQueryResponse 返回参数
type RechargeQueryResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderNumber string `json:"order_number"`
		Status      int    `json:"status"`
		Mobile      string `json:"mobile"`
		Amount      int    `json:"amount"`
		OrderNo     string `json:"order_no"`
	} `json:"data"`
}

// RechargeQuery 查询接口
func (app *App) RechargeQuery(orderNumber string) (result RechargeQueryResponse, err error) {
	// 请求
	body, err := app.request("https://router.wikeyun.cn/rest/Recharge/query", map[string]interface{}{
		"order_number": orderNumber, // 官方订单号
	})
	if err != nil {
		return result, errors.New(err.Error())
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, errors.New(err.Error())
	}
	return result, nil
}
