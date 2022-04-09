package dingdanxia

import (
	"dtapps/dta/library/utils/gohttp"
	"encoding/json"
	"net/http"
)

type WaimaiMeituanOrderidResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Order struct {
			BusinessLine     int         `json:"businessLine"`
			SubBusinessLine  int         `json:"subBusinessLine"`
			ActId            int         `json:"actId"`
			Quantity         int         `json:"quantity"`
			OrderId          string      `json:"orderId"`
			Paytime          string      `json:"paytime"`
			ModTime          string      `json:"modTime"`
			Payprice         string      `json:"payprice"`
			Profit           string      `json:"profit"`
			CpaProfit        string      `json:"cpaProfit"`
			Sid              string      `json:"sid"`
			Appkey           string      `json:"appkey"`
			Smstitle         string      `json:"smstitle"`
			Status           int         `json:"status"`
			TradeTypeList    []int       `json:"tradeTypeList"`
			RiskOrder        interface{} `json:"riskOrder"`
			Refundprofit     interface{} `json:"refundprofit"`
			CpaRefundProfit  interface{} `json:"cpaRefundProfit"`
			RefundInfoList   interface{} `json:"refundInfoList"`
			RefundProfitList interface{} `json:"refundProfitList"`
			Extra            interface{} `json:"extra"`
		} `json:"order"`
	} `json:"data"`
}

type WaimaiMeituanOrderidResult struct {
	Result WaimaiMeituanOrderidResponse // 结果
	Body   []byte                       // 内容
	Http   gohttp.Response              // 请求
	Err    error                        // 错误
}

func NewWaimaiMeituanOrderidResult(result WaimaiMeituanOrderidResponse, body []byte, http gohttp.Response, err error) *WaimaiMeituanOrderidResult {
	return &WaimaiMeituanOrderidResult{Result: result, Body: body, Http: http, Err: err}
}

// WaimaiMeituanOrderid 美团联盟外卖/闪购/优选/酒店订单查询API（订单号版）
// https://www.dingdanxia.com/doc/179/173
func (app *App) WaimaiMeituanOrderid(orderid string) *WaimaiMeituanOrderidResult {
	// 参数
	param := NewParams()
	param.Set("orderid", orderid)
	params := app.NewParamsWith(param)
	// 请求
	request, err := app.request("https://api.tbk.dingdanxia.com/waimai/meituan_orderid", params, http.MethodPost)
	// 定义
	var response WaimaiMeituanOrderidResponse
	err = json.Unmarshal(request.Body, &response)
	return NewWaimaiMeituanOrderidResult(response, request.Body, request, err)
}
