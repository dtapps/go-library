package qweather

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type OceanTideResponse struct {
	Code       string `json:"code"`       // 状态码
	UpdateTime string `json:"updateTime"` // 最近更新时间
	FxLink     string `json:"fxLink"`     // 响应式页面
	TideTable  []struct {
		FxTime string `json:"fxTime"` // 满潮或干潮时间
		Height string `json:"height"` // 海水高度，单位：米
		Type   string `json:"type"`   // 满潮（H）或干潮（L）
	} `json:"tideTable"`
	TideHourly []struct {
		FxTime string `json:"fxTime"` // 逐小时预报时间
		Height string `json:"height"` // 海水高度，单位：米。对于一些地点，此数据可能为空
	} `json:"tideHourly"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，或数据源说明，可能为空
		License []string `json:"license"` // 数据许可或版权声明，可能为空
	} `json:"refer"`
}

type OceanTideResult struct {
	Result OceanTideResponse  // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newOceanTideResult(result OceanTideResponse, body []byte, http gorequest.Response) *OceanTideResult {
	return &OceanTideResult{Result: result, Body: body, Http: http}
}

// OceanTide 潮汐
// https://dev.qweather.com/docs/api/ocean/tide/
func (c *Client) OceanTide(ctx context.Context, location string, notMustParams ...gorequest.Params) (*OceanTideResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("location", location)
	params.Set("key", c.key)
	// 请求
	request, err := c.request(ctx, "ocean/tide", params, http.MethodGet)
	if err != nil {
		return newOceanTideResult(OceanTideResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response OceanTideResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newOceanTideResult(response, request.ResponseBody, request), err
}
