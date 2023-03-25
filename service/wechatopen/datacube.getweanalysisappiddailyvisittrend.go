package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
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

type DataCubeGetWeAnAlySisAppidDailyVisitTrendResult struct {
	Result DataCubeGetWeAnAlySisAppidDailyVisitTrendResponse // 结果
	Body   []byte                                            // 内容
	Http   gorequest.Response                                // 请求
}

func newDataCubeGetWeAnAlySisAppidDailyVisitTrendResult(result DataCubeGetWeAnAlySisAppidDailyVisitTrendResponse, body []byte, http gorequest.Response) *DataCubeGetWeAnAlySisAppidDailyVisitTrendResult {
	return &DataCubeGetWeAnAlySisAppidDailyVisitTrendResult{Result: result, Body: body, Http: http}
}

// DataCubeGetWeAnAlySisAppidDailyVisitTrend 获取用户访问小程序数据日趋势
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/visit-trend/getDailyVisitTrend.html
func (c *Client) DataCubeGetWeAnAlySisAppidDailyVisitTrend(ctx context.Context, beginDate, endDate string) (*DataCubeGetWeAnAlySisAppidDailyVisitTrendResult, error) {
	// 检查
	err := c.checkComponentIsConfig()
	if err != nil {
		return nil, err
	}
	err = c.checkAuthorizerIsConfig()
	if err != nil {
		return nil, err
	}
	// 参数
	params := gorequest.NewParams()
	params.Set("begin_date", beginDate)
	params.Set("end_date", endDate)
	// 请求
	request, err := c.request(ctx, apiUrl+"/datacube/getweanalysisappiddailyvisittrend?access_token="+c.GetAuthorizerAccessToken(ctx), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response DataCubeGetWeAnAlySisAppidDailyVisitTrendResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newDataCubeGetWeAnAlySisAppidDailyVisitTrendResult(response, request.ResponseBody, request), nil
}
