package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type DataCubeGetWeAnAlySisAppidVisitPageResponse struct {
	RefDate string `json:"ref_date"`
	List    []struct {
		PagePath       string  `json:"page_path"`
		PageVisitPv    int     `json:"page_visit_pv"`
		PageVisitUv    int     `json:"page_visit_uv"`
		PageStaytimePv float64 `json:"page_staytime_pv"`
		EntrypagePv    int     `json:"entrypage_pv"`
		ExitpagePv     int     `json:"exitpage_pv"`
		PageSharePv    int     `json:"page_share_pv"`
		PageShareUv    int     `json:"page_share_uv"`
	} `json:"list"`
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
func (c *Client) DataCubeGetWeAnAlySisAppidVisitPage(ctx context.Context, beginDate, endDate string) (*DataCubeGetWeAnAlySisAppidVisitPageResult, error) {
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
	request, err := c.request(ctx, apiUrl+"/datacube/getweanalysisappidvisitpage?access_token="+c.GetAuthorizerAccessToken(ctx), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response DataCubeGetWeAnAlySisAppidVisitPageResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newDataCubeGetWeAnAlySisAppidVisitPageResult(response, request.ResponseBody, request), nil
}
