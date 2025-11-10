package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
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

// DataCubeGetWeAnAlysIsAppidVisitDistribution 获取用户小程序访问分布数据
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/data-analysis/others/getVisitDistribution.html
func (c *Client) DataCubeGetWeAnAlysIsAppidVisitDistribution(ctx context.Context, beginDate, endDate string, notMustParams ...*gorequest.Params) (response DataCubeGetWeAnAlysIsAppidVisitDistributionResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("begin_date", beginDate)
	params.Set("end_date", endDate)

	// 请求
	err = c.request(ctx, "datacube/getweanalysisappidvisitdistribution?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}
