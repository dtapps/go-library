package chengquan

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type OrderDirectChargeResponse struct {
	Code    int    `json:"code"`    // 错误代码
	Message string `json:"message"` // 错误信息
	Data    struct {
		AppID          string  `json:"app_id"`          // 商户账号
		OrderNo        string  `json:"order_no"`        // 商户订单号
		RechargeNumber string  `json:"recharge_number"` // 充值手机号码
		StartTime      string  `json:"start_time"`      // 订单创建时间
		EndTime        string  `json:"end_time"`        // 订单完成时间，如果订单为充值中该时间和创建时间相同，如果成功或者失败该时间为订单完成时间。
		State          string  `json:"state"`           // 订单状态
		ConsumeAmount  float64 `json:"consume_amount"`  // 扣款金额(单位：元)，保留小数点后四位
	} `json:"data"`
}

type OrderDirectChargeResult struct {
	Result OrderDirectChargeResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newOrderDirectChargeResult(result OrderDirectChargeResponse, body []byte, http gorequest.Response) *OrderDirectChargeResult {
	return &OrderDirectChargeResult{Result: result, Body: body, Http: http}
}

// OrderDirectCharge 直充下单接口
// https://chengquan.cn/rechargeInterface/directCharge.html
func (c *Client) OrderDirectCharge(ctx context.Context, notMustParams ...gorequest.Params) (*OrderDirectChargeResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("version", version) // 版本号
	// 请求
	request, err := c.request(ctx, "/rder/directCharge", params, http.MethodPost)
	if err != nil {
		return newOrderDirectChargeResult(OrderDirectChargeResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response OrderDirectChargeResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newOrderDirectChargeResult(response, request.ResponseBody, request), err
}
