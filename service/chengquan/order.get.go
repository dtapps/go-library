package chengquan

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type OrderGetResponse struct {
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

// OrderGet 订单查询接口
// order_no = 商户提交的订单号，最长32位(商户保证其唯一性)
// https://www.chengquan.cn/rechargeInterface/queryOrder.html
func (c *Client) OrderGet(ctx context.Context, orderNo string, notMustParams ...*gorequest.Params) (response OrderGetResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("order_no", orderNo) // 商户提交的订单号，最长32位(商户保证其唯一性)
	if c.config.version != "" {
		params.Set("version", c.config.version) // 版本号
	}

	// 请求
	err = c.request(ctx, "order/get", params, http.MethodPost, &response)
	return
}
