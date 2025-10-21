package chengquan

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type UserBalanceGetResponse struct {
	Code    int    `json:"code"`    // 错误代码
	Message string `json:"message"` // 错误信息
	Data    struct {
		AppID   string  `json:"app_id"`  // 商户账号
		Balance float64 `json:"balance"` // 商户余额(单位：元)
	} `json:"data"`
}

// UserBalanceGet 账号余额查询接口
// https://chengquan.cn/basicData/queryBalance.html
func (c *Client) UserBalanceGet(ctx context.Context, notMustParams ...*gorequest.Params) (response UserBalanceGetResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "user/balance/get", params, http.MethodPost, &response)
	return
}
