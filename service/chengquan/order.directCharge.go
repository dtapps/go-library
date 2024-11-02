package chengquan

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
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
// order_no = 商户提交的订单号，最长32位(商户保证其唯一性)
// recharge_number = 充值账号
// product_id = 充值产品编号
// amount = 充值数量（加油卡，视频业务默认为1，其它业务按照实际情况传递）。数量范围1-99999
// ip = 充值ip（仅腾讯业务需要，根据实际情况进行传递）
// oil_phone_account = 加油卡充值时用户的手机号
// notify_url = 橙券主动通知订单结果地址
// https://chengquan.cn/rechargeInterface/directCharge.html
func (c *Client) OrderDirectCharge(ctx context.Context, orderNo string, rechargeNumber string, productID int64, amount int64, notMustParams ...gorequest.Params) (*OrderDirectChargeResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("order_no", orderNo)               // 商户提交的订单号，最长32位(商户保证其唯一性)
	params.Set("recharge_number", rechargeNumber) // 充值账号
	params.Set("product_id", productID)           // 充值产品编号
	params.Set("amount", amount)                  // 充值数量（加油卡，视频业务默认为1，其它业务按照实际情况传递）。数量范围1-99999
	params.Set("version", c.config.version)       // 版本号

	// 请求
	var response OrderDirectChargeResponse
	request, err := c.request(ctx, "rder/directCharge", params, http.MethodPost, &response)
	return newOrderDirectChargeResult(response, request.ResponseBody, request), err
}
