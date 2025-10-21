package qxwlwagnt

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type V3CardOperatorResponse struct {
	Iccid string `json:"iccid"` // 物联网号码的ICCID
}

// V3CardOperator 业务状态变更
// http://docs.konyun.net/web/#/71/2390
func (c *Client) V3CardOperator(ctx context.Context, iccid string, action string, notMustParams ...*gorequest.Params) (response CommonResponse[V3CardOperatorResponse], err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("iccid", iccid)   // 集成电路卡识别码，IC 卡的唯一识别号码
	params.Set("action", action) // 动作（具体参数请联系接口提供方索取）

	// 请求
	err = c.Request(ctx, "/api/v3/cardOperator", params, http.MethodGet, &response)
	return
}
