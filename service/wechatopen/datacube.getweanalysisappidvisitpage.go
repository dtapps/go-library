package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type DataCubeGetWeAnAlySisAppidVisitPageResponse struct {
	RefDate string `json:"ref_date"` // 日期
	List    []struct {
		PagePath       string  `json:"page_path"`        // 页面路径
		PageVisitPv    int64   `json:"page_visit_pv"`    // 访问次数
		PageVisitUv    int64   `json:"page_visit_uv"`    // 访问人数
		PageStaytimePv float64 `json:"page_staytime_pv"` // 次均停留时长
		EntrypagePv    int64   `json:"entrypage_pv"`     // 进入页次数
		ExitpagePv     int64   `json:"exitpage_pv"`      // 退出页次数
		PageSharePv    int64   `json:"page_share_pv"`    // 转发次数
		PageShareUv    int64   `json:"page_share_uv"`    // 转发人数
	} `json:"list"` // 数据列表
}

// DataCubeGetWeAnAlySisAppidVisitPage 获取访问页面数据
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/others/getVisitPage.html
func (c *Client) DataCubeGetWeAnAlySisAppidVisitPage(ctx context.Context, beginDate string, endDate string, notMustParams ...*gorequest.Params) (response DataCubeGetWeAnAlySisAppidVisitPageResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("begin_date", beginDate)
	params.Set("end_date", endDate)

	// 请求
	err = c.request(ctx, "datacube/getweanalysisappidvisitpage?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}
