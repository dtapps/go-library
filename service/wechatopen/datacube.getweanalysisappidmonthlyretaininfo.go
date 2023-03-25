package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResponse struct {
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
func (c *Client) DataCubeGetWeAnAlySisAppidMonthlyRetainInfo(ctx context.Context, beginDate, endDate string) (*DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResult, error) {
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
	request, err := c.request(ctx, apiUrl+"/datacube/getweanalysisappidmonthlyretaininfo?access_token="+c.GetAuthorizerAccessToken(ctx), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response DataCubeGetWeAnAlySisAppidMonthlyRetainInfoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newDataCubeGetWeAnAlySisAppidMonthlyRetainInfoResult(response, request.ResponseBody, request), nil
}
