package dingdanxia

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WaiMaiMeituanOrderIdResponse struct {
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

type WaiMaiMeituanOrderIdResult struct {
	Result WaiMaiMeituanOrderIdResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
	Err    error                        // 错误
}

func newWaiMaiMeituanOrderIdResult(result WaiMaiMeituanOrderIdResponse, body []byte, http gorequest.Response, err error) *WaiMaiMeituanOrderIdResult {
	return &WaiMaiMeituanOrderIdResult{Result: result, Body: body, Http: http, Err: err}
}

// WaiMaiMeituanOrderId 美团联盟外卖/闪购/优选/酒店订单查询API（订单号版）
// https://www.dingdanxia.com/doc/179/173
func (c *Client) WaiMaiMeituanOrderId(orderId string) *WaiMaiMeituanOrderIdResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("order_id", orderId)
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(apiUrl+"/waimai/meituan_orderid", params, http.MethodPost)
	// 定义
	var response WaiMaiMeituanOrderIdResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newWaiMaiMeituanOrderIdResult(response, request.ResponseBody, request, err)
}
