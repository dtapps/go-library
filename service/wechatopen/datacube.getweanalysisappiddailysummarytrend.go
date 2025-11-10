package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type DataCubeGetWeAnAlySisAppidDailySummaryTrendResponse struct {
	List []struct {
		RefDate    string `json:"ref_date"`    // 日期
		VisitTotal int64  `json:"visit_total"` // 累计用户数
		SharePv    int64  `json:"share_pv"`    // 转发次数
		ShareUv    int64  `json:"share_uv"`    // 转发人数
	} `json:"list"` // 数据列表
}

// DataCubeGetWeAnAlySisAppidDailySummaryTrend 获取用户访问小程序数据概况
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/others/getDailySummary.html
func (c *Client) DataCubeGetWeAnAlySisAppidDailySummaryTrend(ctx context.Context, authorizerAccessToken, beginDate, endDate string, notMustParams ...*gorequest.Params) (response DataCubeGetWeAnAlySisAppidDailySummaryTrendResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("begin_date", beginDate)
	params.Set("end_date", endDate)

	// 请求
	err = c.request(ctx, "datacube/getweanalysisappiddailysummarytrend?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}
