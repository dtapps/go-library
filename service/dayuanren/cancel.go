package dayuanren

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type CancelResponse struct {
	Errno  int64    `json:"errno"`  // 错误码，0代表成功，非0代表失败
	Errmsg string   `json:"errmsg"` // 错误描述
	Data   struct{} `json:"data"`
}

type CancelResult struct {
	Result CancelResponse     // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newCancelResult(result CancelResponse, body []byte, http gorequest.Response) *CancelResult {
	return &CancelResult{Result: result, Body: body, Http: http}
}

// Cancel 退单申请
// out_trade_num = 商户订单号；多个用英文,分割
// https://www.kancloud.cn/boyanyun/boyanyun_huafei/3182909
func (c *Client) Cancel(ctx context.Context, outTradeNums string, notMustParams ...*gorequest.Params) (*CancelResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("userid", c.GetUserID())        // 账户ID
	params.Set("out_trade_nums", outTradeNums) // 商户订单号；多个用英文,分割

	// 请求
	var response CancelResponse
	request, err := c.request(ctx, "index/cancel", params, &response)
	return newCancelResult(response, request.ResponseBody, request), err
}
