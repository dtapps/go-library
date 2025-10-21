package qxwlwagnt

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type V3QueryFlowHistoryResponse struct {
	Iccid        string `json:"iccid"`        // 物联网号码的ICCID
	BillingMonth string `json:"billingMonth"` // 月份，格式：yyyyMM
	UsageTotal   string `json:"usageTotal"`   // 数据使用量，单位 M
	Status       string `json:"status"`       // Y:成功，N:数据未同步
}

// V3QueryFlowHistory 历史流量查询
// http://docs.konyun.net/web/#/71/2400
func (c *Client) V3QueryFlowHistory(ctx context.Context, iccid string, billingMonth string, notMustParams ...*gorequest.Params) (response CommonResponse[V3QueryFlowHistoryResponse], err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("iccid", iccid)               // 集成电路卡识别码，IC 卡的唯一识别号码
	params.Set("billingMonth", billingMonth) // 月份，格式：yyyyMM, 例如202111

	// 请求
	err = c.Request(ctx, "/api/v3/queryFlowHistory", params, http.MethodGet, &response)
	return
}
