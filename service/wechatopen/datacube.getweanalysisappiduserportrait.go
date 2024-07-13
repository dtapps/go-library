package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type DataCubeGetWeAnAlySisAppidUserPortraitResponse struct {
	RefDate    string `json:"ref_date"` // 时间范围
	VisitUvNew struct {
		Province []struct {
			Id    int64  `json:"id"`
			Name  string `json:"name"`
			Value int64  `json:"value"`
		} `json:"province"` // 分布类型
		City []struct {
			Id    int64  `json:"id"`
			Name  string `json:"name"`
			Value int64  `json:"value"`
		} `json:"city"` // 省份
		Genders []struct {
			Id    int64  `json:"id"`
			Name  string `json:"name"`
			Value int64  `json:"value"`
		} `json:"genders"` // 城市
		Platforms []struct {
			Id    int64  `json:"id"`
			Name  string `json:"name"`
			Value int64  `json:"value"`
		} `json:"platforms"` // 性别
		Devices []struct {
			Name  string `json:"name"`
			Value int64  `json:"value"`
		} `json:"devices"` // 终端类型
		Ages []struct {
			Id    int64  `json:"id"`
			Name  string `json:"name"`
			Value int64  `json:"value"`
		} `json:"ages"` // 年龄
	} `json:"visit_uv_new"` // 新用户画像
	VisitUv struct {
		Province []struct {
			Id    int64  `json:"id"`
			Name  string `json:"name"`
			Value int64  `json:"value"`
		} `json:"province"` // 分布类型
		City []struct {
			Id    int64  `json:"id"`
			Name  string `json:"name"`
			Value int64  `json:"value"`
		} `json:"city"` // 省份
		Genders []struct {
			Id    int64  `json:"id"`
			Name  string `json:"name"`
			Value int64  `json:"value"`
		} `json:"genders"` // 城市
		Platforms []struct {
			Id    int64  `json:"id"`
			Name  string `json:"name"`
			Value int64  `json:"value"`
		} `json:"platforms"` // 性别
		Devices []struct {
			Name  string `json:"name"`
			Value int64  `json:"value"`
		} `json:"devices"` // 终端类型
		Ages []struct {
			Id    int64  `json:"id"`
			Name  string `json:"name"`
			Value int64  `json:"value"`
		} `json:"ages"` // 年龄
	} `json:"visit_uv"` // 活跃用户画像
}

type DataCubeGetWeAnAlySisAppidUserPortraitResult struct {
	Result DataCubeGetWeAnAlySisAppidUserPortraitResponse // 结果
	Body   []byte                                         // 内容
	Http   gorequest.Response                             // 请求
}

func newDataCubeGetWeAnAlySisAppidUserPortraitResult(result DataCubeGetWeAnAlySisAppidUserPortraitResponse, body []byte, http gorequest.Response) *DataCubeGetWeAnAlySisAppidUserPortraitResult {
	return &DataCubeGetWeAnAlySisAppidUserPortraitResult{Result: result, Body: body, Http: http}
}

// DataCubeGetWeAnAlySisAppidUserPortrait 获取小程序用户画像分布
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/others/getUserPortrait.html
func (c *Client) DataCubeGetWeAnAlySisAppidUserPortrait(ctx context.Context, authorizerAccessToken, beginDate, endDate string, notMustParams ...gorequest.Params) (*DataCubeGetWeAnAlySisAppidUserPortraitResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "datacube/getweanalysisappiduserportrait")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("begin_date", beginDate)
	params.Set("end_date", endDate)

	// 请求
	var response DataCubeGetWeAnAlySisAppidUserPortraitResponse
	request, err := c.request(ctx, span, "datacube/getweanalysisappiduserportrait?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newDataCubeGetWeAnAlySisAppidUserPortraitResult(response, request.ResponseBody, request), err
}
