package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type DataCubeGetWeAnAlySisAppidWeeklyRetainInfoResponse struct {
	RefDate    string `json:"ref_date"` // 日期
	VisitUvNew []struct {
		Key   int64 `json:"key"`
		Value int64 `json:"value"`
	} `json:"visit_uv_new"` // 新增用户留存
	VisitUv []struct {
		Key   int64 `json:"key"`
		Value int64 `json:"value"`
	} `json:"visit_uv"` // 活跃用户留存
}

type DataCubeGetWeAnAlySisAppidWeeklyRetainInfoResult struct {
	Result DataCubeGetWeAnAlySisAppidWeeklyRetainInfoResponse // 结果
	Body   []byte                                             // 内容
	Http   gorequest.Response                                 // 请求
}

func newDataCubeGetWeAnAlySisAppidWeeklyRetainInfoResult(result DataCubeGetWeAnAlySisAppidWeeklyRetainInfoResponse, body []byte, http gorequest.Response) *DataCubeGetWeAnAlySisAppidWeeklyRetainInfoResult {
	return &DataCubeGetWeAnAlySisAppidWeeklyRetainInfoResult{Result: result, Body: body, Http: http}
}

// DataCubeGetWeAnAlySisAppidWeeklyRetainInfo 获取用户访问小程序周留存
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/visit-retain/getWeeklyRetain.html
func (c *Client) DataCubeGetWeAnAlySisAppidWeeklyRetainInfo(ctx context.Context, authorizerAccessToken, beginDate, endDate string, notMustParams ...gorequest.Params) (*DataCubeGetWeAnAlySisAppidWeeklyRetainInfoResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("begin_date", beginDate)
	params.Set("end_date", endDate)

	// 请求
	var response DataCubeGetWeAnAlySisAppidWeeklyRetainInfoResponse
	request, err := c.request(ctx, "datacube/getweanalysisappidweeklyretaininfo?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newDataCubeGetWeAnAlySisAppidWeeklyRetainInfoResult(response, request.ResponseBody, request), err
}
