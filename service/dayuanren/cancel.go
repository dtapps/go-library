package dayuanren

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type Cancel struct {
	Errno  int64    `json:"errno"`  // 错误码，0代表成功，非0代表失败
	Errmsg string   `json:"errmsg"` // 错误描述
	Data   struct{} `json:"data"`
}

// Cancel 退单申请
// out_trade_num = 商户订单号；多个用英文,分割
// https://www.kancloud.cn/boyanyun/boyanyun_huafei/3182909
func (c *Client) Cancel(ctx context.Context, outTradeNums string, notMustParams ...*gorequest.Params) (response Cancel, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("userid", c.GetUserID())        // 账户ID
	params.Set("out_trade_nums", outTradeNums) // 商户订单号；多个用英文,分割

	// 请求
	err = c.request(ctx, "index/cancel", params, &response)
	return
}
