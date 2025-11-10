package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type DataCubeGetWeAnAlySisAppidDailyVisitTrendResponse struct {
	List []struct {
		RefDate         string  `json:"ref_date"`          // 时间
		SessionCnt      int64   `json:"session_cnt"`       // 打开次数（自然周内汇总）
		VisitPv         int64   `json:"visit_pv"`          // 访问次数（自然周内汇总）
		VisitUv         int64   `json:"visit_uv"`          // 访问人数（自然周内去重）
		VisitUvNew      int64   `json:"visit_uv_new"`      // 新用户数（自然周内去重）
		StayTimeUv      float64 `json:"stay_time_uv"`      // 人均停留时长 (浮点型，单位：秒)
		StayTimeSession float64 `json:"stay_time_session"` // 次均停留时长 (浮点型，单位：秒)
		VisitDepth      float64 `json:"visit_depth"`       // 平均访问深度 (浮点型)
	} `json:"list"` // 数据列表
}

// DataCubeGetWeAnAlySisAppidDailyVisitTrend 获取用户访问小程序数据日趋势
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/visit-trend/getDailyVisitTrend.html
func (c *Client) DataCubeGetWeAnAlySisAppidDailyVisitTrend(ctx context.Context, beginDate, endDate string, notMustParams ...*gorequest.Params) (response DataCubeGetWeAnAlySisAppidDailyVisitTrendResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("begin_date", beginDate)
	params.Set("end_date", endDate)

	// 请求
	err = c.request(ctx, "datacube/getweanalysisappiddailyvisittrend?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}
