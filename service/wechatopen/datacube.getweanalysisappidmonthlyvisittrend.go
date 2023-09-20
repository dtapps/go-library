package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type DataCubeGetWeAnAlySisAppidMonthlyVisitTrendResponse struct {
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

type DataCubeGetWeAnAlySisAppidMonthlyVisitTrendResult struct {
	Result DataCubeGetWeAnAlySisAppidMonthlyVisitTrendResponse // 结果
	Body   []byte                                              // 内容
	Http   gorequest.Response                                  // 请求
}

func newDataCubeGetWeAnAlySisAppidMonthlyVisitTrendResult(result DataCubeGetWeAnAlySisAppidMonthlyVisitTrendResponse, body []byte, http gorequest.Response) *DataCubeGetWeAnAlySisAppidMonthlyVisitTrendResult {
	return &DataCubeGetWeAnAlySisAppidMonthlyVisitTrendResult{Result: result, Body: body, Http: http}
}

// DataCubeGetWeAnAlySisAppidMonthlyVisitTrend 获取用户访问小程序数据月趋势
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/visit-trend/getMonthlyVisitTrend.html
func (c *Client) DataCubeGetWeAnAlySisAppidMonthlyVisitTrend(ctx context.Context, authorizerAccessToken, beginDate, endDate string, notMustParams ...*gorequest.Params) (*DataCubeGetWeAnAlySisAppidMonthlyVisitTrendResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("begin_date", beginDate)
	params.Set("end_date", endDate)
	// 请求
	request, err := c.request(ctx, apiUrl+"/datacube/getweanalysisappidmonthlyvisittrend?access_token="+authorizerAccessToken, params, http.MethodPost)
	if err != nil {
		return newDataCubeGetWeAnAlySisAppidMonthlyVisitTrendResult(DataCubeGetWeAnAlySisAppidMonthlyVisitTrendResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response DataCubeGetWeAnAlySisAppidMonthlyVisitTrendResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newDataCubeGetWeAnAlySisAppidMonthlyVisitTrendResult(response, request.ResponseBody, request), err
}
