package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResponse struct {
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

type DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResult struct {
	Result DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResponse // 结果
	Body   []byte                                              // 内容
	Http   gorequest.Response                                  // 请求
}

func newDataCubeGetWeAnAlySisAppidMonthlyRetainInfoResult(result DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResponse, body []byte, http gorequest.Response) *DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResult {
	return &DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResult{Result: result, Body: body, Http: http}
}

// DataCubeGetWeAnAlySisAppidMonthlyRetainInfo 获取用户访问小程序月留存
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/visit-retain/getMonthlyRetain.html
func (c *Client) DataCubeGetWeAnAlySisAppidMonthlyRetainInfo(ctx context.Context, beginDate, endDate string, notMustParams ...gorequest.Params) (*DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResult, error) {
	// 检查
	if err := c.checkAuthorizerConfig(ctx); err != nil {
		return newDataCubeGetWeAnAlySisAppidMonthlyRetainInfoResult(DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResponse{}, []byte{}, gorequest.Response{}), err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("begin_date", beginDate)
	params.Set("end_date", endDate)
	// 请求
	request, err := c.request(ctx, apiUrl+"/datacube/getweanalysisappidmonthlyretaininfo?access_token="+GetAuthorizerAccessToken(ctx, c), params, http.MethodPost)
	if err != nil {
		return newDataCubeGetWeAnAlySisAppidMonthlyRetainInfoResult(DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newDataCubeGetWeAnAlySisAppidMonthlyRetainInfoResult(response, request.ResponseBody, request), err
}
