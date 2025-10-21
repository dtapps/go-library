package qxwlwagnt

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type V3QueryAccountInfoResponse struct {
	Username         string `json:"username"`         // 用户名
	CompanyNameShort string `json:"companyNameShort"` // 公司简称
	Balance          string `json:"balance"`          // 账户余额(元)
	GuaranteeAmount  string `json:"guaranteeAmount"`  // 保证金额(元)
	PrePayAmount     string `json:"prePayAmount"`     // 预付费金额(元)
	Status           int    `json:"status"`           // 状态 0：停用 1：正常
	PayMent          int    `json:"payMent"`          // 付费模式 0：预付费 1：后付费
}

// V3QueryAccountInfo 账户信息查询
// http://docs.konyun.net/web/#/71/2397
func (c *Client) V3QueryAccountInfo(ctx context.Context, notMustParams ...*gorequest.Params) (response CommonResponse[V3QueryAccountInfoResponse], err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.Request(ctx, "/api/v3/queryAccountInfo", params, http.MethodPost, &response)
	return
}
