package qxwlwagnt

import (
	"context"
	"net/http"
	"strings"

	"go.dtapp.net/library/utils/gorequest"
)

type V3BatchQueryFlowHistoryResponse struct {
	Iccid       string `json:"iccid"` // 物联网号码的ICCID
	IccidResult struct {
		Iccid        string `json:"iccid"`        // 物联网号码的ICCID
		BillingMonth string `json:"billingMonth"` // 月份，格式：yyyyMM
		UsageTotal   string `json:"usageTotal"`   // 数据使用量，单位 M
		Status       string `json:"status"`       // Y:成功，N:数据未同步
	} `json:"iccidResult"` // 物联网号码的ICCID相关数据
}

// V3 BatchQueryFlowHistory 批量历史流量查询
// http://docs.konyun.net/web/#/71/2401
func (c *Client) V3BatchQueryFlowHistory(ctx context.Context, iccids []string, billingMonth string, notMustParams ...*gorequest.Params) (response CommonResponse[[]V3BatchQueryFlowHistoryResponse], err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("iccids", strings.Join(iccids, "_")) // 集成电路卡识别码，IC 卡的唯一识别号码(多个iccid之间_分割,最多不超过200个)
	params.Set("billingMonth", billingMonth)        // 月份，格式：yyyyMM, 例如202111

	// 请求
	err = c.Request(ctx, "/api/v3/batchQueryFlowHistory", params, http.MethodGet, &response)
	return
}
