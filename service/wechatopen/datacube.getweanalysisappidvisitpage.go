package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
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

type DataCubeGetWeAnAlySisAppidVisitPageResult struct {
	Result DataCubeGetWeAnAlySisAppidVisitPageResponse // 结果
	Body   []byte                                      // 内容
	Http   gorequest.Response                          // 请求
}

func newDataCubeGetWeAnAlySisAppidVisitPageResult(result DataCubeGetWeAnAlySisAppidVisitPageResponse, body []byte, http gorequest.Response) *DataCubeGetWeAnAlySisAppidVisitPageResult {
	return &DataCubeGetWeAnAlySisAppidVisitPageResult{Result: result, Body: body, Http: http}
}

// DataCubeGetWeAnAlySisAppidVisitPage 获取访问页面数据
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/others/getVisitPage.html
func (c *Client) DataCubeGetWeAnAlySisAppidVisitPage(ctx context.Context, authorizerAccessToken, beginDate, endDate string, notMustParams ...*gorequest.Params) (*DataCubeGetWeAnAlySisAppidVisitPageResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("begin_date", beginDate)
	params.Set("end_date", endDate)
	// 请求
	request, err := c.request(ctx, apiUrl+"/datacube/getweanalysisappidvisitpage?access_token="+authorizerAccessToken, params, http.MethodPost)
	if err != nil {
		return newDataCubeGetWeAnAlySisAppidVisitPageResult(DataCubeGetWeAnAlySisAppidVisitPageResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response DataCubeGetWeAnAlySisAppidVisitPageResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newDataCubeGetWeAnAlySisAppidVisitPageResult(response, request.ResponseBody, request), err
}
