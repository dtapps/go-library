package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type DataCubeGetWeAnAlySisAppidWeeklyRetainInfoResponse struct {
	RefDate    string `json:"ref_date"`
	VisitUvNew []struct {
		Key   int `json:"key"`
		Value int `json:"value"`
	} `json:"visit_uv_new"`
	VisitUv []struct {
		Key   int `json:"key"`
		Value int `json:"value"`
	} `json:"visit_uv"`
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
func (c *Client) DataCubeGetWeAnAlySisAppidWeeklyRetainInfo(ctx context.Context, beginDate, endDate string) (*DataCubeGetWeAnAlySisAppidWeeklyRetainInfoResult, error) {
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
	request, err := c.request(ctx, apiUrl+"/datacube/getweanalysisappidweeklyretaininfo?access_token="+c.GetAuthorizerAccessToken(ctx), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response DataCubeGetWeAnAlySisAppidWeeklyRetainInfoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newDataCubeGetWeAnAlySisAppidWeeklyRetainInfoResult(response, request.ResponseBody, request), nil
}
