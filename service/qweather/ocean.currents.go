package qweather

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type OceanCurrentsResponse struct {
	Code          string `json:"code"`       // 状态码
	UpdateTime    string `json:"updateTime"` // 最近更新时间
	FxLink        string `json:"fxLink"`     // 响应式页面
	CurrentsTable []struct {
		FxTime   string `json:"fxTime"`   // 潮流最大流速时间
		SpeedMax string `json:"speedMax"` // 潮流最大流速，单位：厘米/秒
		Dir360   string `json:"dir360"`   // 潮流360度方向
	} `json:"currentsTable"`
	CurrentsHourly []struct {
		FxTime string `json:"fxTime"` // 逐小时预报时间
		Speed  string `json:"speed"`  // 潮流流速，单位：厘米/秒
		Dir360 string `json:"dir360"` // 潮流360度方向
	} `json:"currentsHourly"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，或数据源说明，可能为空
		License []string `json:"license"` // 数据许可或版权声明，可能为空
	} `json:"refer"`
}

type OceanCurrentsResult struct {
	Result OceanCurrentsResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newOceanCurrentsResult(result OceanCurrentsResponse, body []byte, http gorequest.Response) *OceanCurrentsResult {
	return &OceanCurrentsResult{Result: result, Body: body, Http: http}
}

// OceanCurrents 潮汐
// https://dev.qweather.com/docs/api/ocean/tide/
func (c *Client) OceanCurrents(ctx context.Context, location string, notMustParams ...gorequest.Params) (*OceanCurrentsResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("location", location)
	params.Set("key", c.key)
	// 请求
	request, err := c.request(ctx, "ocean/currents", params, http.MethodGet)
	if err != nil {
		return newOceanCurrentsResult(OceanCurrentsResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response OceanCurrentsResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newOceanCurrentsResult(response, request.ResponseBody, request), err
}
