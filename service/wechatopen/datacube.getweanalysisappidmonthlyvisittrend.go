package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type DataCubeGetWeAnAlySisAppidMonthlyVisitTrendResponse struct {
	List []struct {
		RefDate         string  `json:"ref_date"`
		SessionCnt      int     `json:"session_cnt"`
		VisitPv         int     `json:"visit_pv"`
		VisitUv         int     `json:"visit_uv"`
		VisitUvNew      int     `json:"visit_uv_new"`
		StayTimeSession float64 `json:"stay_time_session"`
		VisitDepth      float64 `json:"visit_depth"`
	} `json:"list"`
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
func (c *Client) DataCubeGetWeAnAlySisAppidMonthlyVisitTrend(ctx context.Context, beginDate, endDate string) (*DataCubeGetWeAnAlySisAppidMonthlyVisitTrendResult, error) {
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
	request, err := c.request(ctx, apiUrl+"/datacube/getweanalysisappidmonthlyvisittrend?access_token="+c.GetAuthorizerAccessToken(ctx), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response DataCubeGetWeAnAlySisAppidMonthlyVisitTrendResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newDataCubeGetWeAnAlySisAppidMonthlyVisitTrendResult(response, request.ResponseBody, request), nil
}
