package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type DataCubeGetWeAnAlySisAppidDailyRetainInfoResponse struct {
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

type DataCubeGetWeAnAlySisAppidDailyRetainInfoResult struct {
	Result DataCubeGetWeAnAlySisAppidDailyRetainInfoResponse // 结果
	Body   []byte                                            // 内容
	Http   gorequest.Response                                // 请求
}

func newDataCubeGetWeAnAlySisAppidDailyRetainInfoResult(result DataCubeGetWeAnAlySisAppidDailyRetainInfoResponse, body []byte, http gorequest.Response) *DataCubeGetWeAnAlySisAppidDailyRetainInfoResult {
	return &DataCubeGetWeAnAlySisAppidDailyRetainInfoResult{Result: result, Body: body, Http: http}
}

// DataCubeGetWeAnAlySisAppidDailyRetainInfo 获取用户访问小程序日留存
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/visit-retain/getDailyRetain.html
func (c *Client) DataCubeGetWeAnAlySisAppidDailyRetainInfo(ctx context.Context, authorizerAccessToken, beginDate, endDate string, notMustParams ...gorequest.Params) (*DataCubeGetWeAnAlySisAppidDailyRetainInfoResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("begin_date", beginDate)
	params.Set("end_date", endDate)

	// 请求
	var response DataCubeGetWeAnAlySisAppidDailyRetainInfoResponse
	request, err := c.request(ctx, "datacube/getweanalysisappiddailyretaininfo?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newDataCubeGetWeAnAlySisAppidDailyRetainInfoResult(response, request.ResponseBody, request), err
}
