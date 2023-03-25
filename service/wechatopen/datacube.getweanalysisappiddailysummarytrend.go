package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type DataCubeGetWeAnAlySisAppidDailySummaryTrendResponse struct {
	List []struct {
		RefDate    string `json:"ref_date"`
		VisitTotal int    `json:"visit_total"`
		SharePv    int    `json:"share_pv"`
		ShareUv    int    `json:"share_uv"`
	} `json:"list"`
}

type DataCubeGetWeAnAlySisAppidDailySummaryTrendResult struct {
	Result DataCubeGetWeAnAlySisAppidDailySummaryTrendResponse // 结果
	Body   []byte                                              // 内容
	Http   gorequest.Response                                  // 请求
}

func newDataCubeGetWeAnAlySisAppidDailySummaryTrendResult(result DataCubeGetWeAnAlySisAppidDailySummaryTrendResponse, body []byte, http gorequest.Response) *DataCubeGetWeAnAlySisAppidDailySummaryTrendResult {
	return &DataCubeGetWeAnAlySisAppidDailySummaryTrendResult{Result: result, Body: body, Http: http}
}

// DataCubeGetWeAnAlySisAppidDailySummaryTrend 获取用户访问小程序数据概况
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/others/getDailySummary.html
func (c *Client) DataCubeGetWeAnAlySisAppidDailySummaryTrend(ctx context.Context, beginDate, endDate string) (*DataCubeGetWeAnAlySisAppidDailySummaryTrendResult, error) {
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
	request, err := c.request(ctx, apiUrl+"/datacube/getweanalysisappiddailysummarytrend?access_token="+c.GetAuthorizerAccessToken(ctx), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response DataCubeGetWeAnAlySisAppidDailySummaryTrendResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newDataCubeGetWeAnAlySisAppidDailySummaryTrendResult(response, request.ResponseBody, request), nil
}
