package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type DataCubeGetWeAnAlysIsAppidVisitDistributionResponse struct {
	RefDate string `json:"ref_date"` // 日期
	List    []struct {
		Index    string `json:"index"` // 分布类型。枚举值为：access_source_session_cnt（访问来源分布）、access_staytime_info（访问时长分布）、access_depth_info（访问深度的分布 ）
		ItemList []struct {
			Key   int64 `json:"key"`   // 场景 id，定义在各个 index 下不同，具体参见下方表格
			Value int64 `json:"value"` // 该场景 id 访问 pv
		} `json:"item_list"` // 分布数据列表
	} `json:"list"` // 数据列表
}

type DataCubeGetWeAnAlysIsAppidVisitDistributionResult struct {
	Result DataCubeGetWeAnAlysIsAppidVisitDistributionResponse // 结果
	Body   []byte                                              // 内容
	Http   gorequest.Response                                  // 请求
}

func newDataCubeGetWeAnAlysIsAppidVisitDistributionResult(result DataCubeGetWeAnAlysIsAppidVisitDistributionResponse, body []byte, http gorequest.Response) *DataCubeGetWeAnAlysIsAppidVisitDistributionResult {
	return &DataCubeGetWeAnAlysIsAppidVisitDistributionResult{Result: result, Body: body, Http: http}
}

// DataCubeGetWeAnAlysIsAppidVisitDistribution 获取用户小程序访问分布数据
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/others/getVisitDistribution.html
func (c *Client) DataCubeGetWeAnAlysIsAppidVisitDistribution(ctx context.Context, authorizerAccessToken, beginDate, endDate string, notMustParams ...*gorequest.Params) (*DataCubeGetWeAnAlysIsAppidVisitDistributionResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("begin_date", beginDate)
	params.Set("end_date", endDate)
	// 请求
	request, err := c.request(ctx, apiUrl+"/datacube/getweanalysisappidvisitdistribution?access_token="+authorizerAccessToken, params, http.MethodPost)
	if err != nil {
		return newDataCubeGetWeAnAlysIsAppidVisitDistributionResult(DataCubeGetWeAnAlysIsAppidVisitDistributionResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response DataCubeGetWeAnAlysIsAppidVisitDistributionResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newDataCubeGetWeAnAlysIsAppidVisitDistributionResult(response, request.ResponseBody, request), err
}
