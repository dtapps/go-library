package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
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

// DataCubeGetWeAnAlySisAppidUserPortrait 获取小程序用户画像分布
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/others/getUserPortrait.html
func (c *Client) DataCubeGetWeAnAlySisAppidUserPortrait(ctx context.Context, beginDate, endDate string, notMustParams ...*gorequest.Params) (response DataCubeGetWeAnAlySisAppidUserPortraitResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("begin_date", beginDate)
	params.Set("end_date", endDate)

	// 请求
	err = c.request(ctx, "datacube/getweanalysisappiduserportrait?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}
