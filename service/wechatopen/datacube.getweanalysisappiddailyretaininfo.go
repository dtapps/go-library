package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type DataCubeGetWeAnAlySisAppidDailyRetainInfoResponse struct {
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
func (c *Client) DataCubeGetWeAnAlySisAppidDailyRetainInfo(ctx context.Context, beginDate, endDate string) (*DataCubeGetWeAnAlySisAppidDailyRetainInfoResult, error) {
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
	request, err := c.request(ctx, apiUrl+"/datacube/getweanalysisappiddailyretaininfo?access_token="+c.GetAuthorizerAccessToken(ctx), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response DataCubeGetWeAnAlySisAppidDailyRetainInfoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newDataCubeGetWeAnAlySisAppidDailyRetainInfoResult(response, request.ResponseBody, request), nil
}
