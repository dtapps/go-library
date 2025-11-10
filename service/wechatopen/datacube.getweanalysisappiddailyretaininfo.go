package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
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

// DataCubeGetWeAnAlySisAppidDailyRetainInfo 获取用户访问小程序日留存
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/visit-retain/getDailyRetain.html
func (c *Client) DataCubeGetWeAnAlySisAppidDailyRetainInfo(ctx context.Context, authorizerAccessToken, beginDate, endDate string, notMustParams ...*gorequest.Params) (response DataCubeGetWeAnAlySisAppidDailyRetainInfoResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("begin_date", beginDate)
	params.Set("end_date", endDate)

	// 请求
	err = c.request(ctx, "datacube/getweanalysisappiddailyretaininfo?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}
