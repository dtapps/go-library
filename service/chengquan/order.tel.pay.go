package chengquan

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type OrderTelPayResponse struct {
	Code    int    `json:"code"`    // 错误代码
	Message string `json:"message"` // 错误信息
	Data    struct {
		AppID          string  `json:"app_id"`          // 商户账号
		OrderNo        string  `json:"order_no"`        // 商户订单号
		RechargeNumber string  `json:"recharge_number"` // 充值账号
		StartTime      string  `json:"start_time"`      // 订单创建时间
		EndTime        string  `json:"end_time"`        // 订单完成时间，如果订单为充值中该时间和创建时间相同，如果成功或者失败该时间为订单完成时间。
		State          string  `json:"state"`           // 订单状态
		ConsumeAmount  float64 `json:"consume_amount"`  // 扣款金额(单位：元)，保留小数点后四位
	} `json:"data"`
}

type OrderTelPayResult struct {
	Result OrderTelPayResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newOrderTelPayResult(result OrderTelPayResponse, body []byte, http gorequest.Response) *OrderTelPayResult {
	return &OrderTelPayResult{Result: result, Body: body, Http: http}
}

// OrderTelPay 话费下单接口
// https://www.chengquan.cn/rechargeInterface/tel.html
func (c *Client) OrderTelPay(ctx context.Context, orderNo string, rechargeNumber string, notMustParams ...gorequest.Params) (*OrderTelPayResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("order_no", orderNo)               // 商户提交的订单号，最长32位(商户保证其唯一性)
	params.Set("recharge_number", rechargeNumber) // 充值手机号码
	// 请求
	request, err := c.request(ctx, "/order/tel/pay", params, http.MethodPost)
	if err != nil {
		return newOrderTelPayResult(OrderTelPayResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response OrderTelPayResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newOrderTelPayResult(response, request.ResponseBody, request), err
}
