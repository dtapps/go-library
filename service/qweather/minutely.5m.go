package qweather

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type Minutely5MResponse struct {
	Code       string `json:"code"`       // 状态码
	UpdateTime string `json:"updateTime"` // 最近更新时间
	FxLink     string `json:"fxLink"`     // 响应式页面
	Summary    string `json:"summary"`    // 分钟降水描述
	Minutely   []struct {
		FxTime string `json:"fxTime"` // 预报时间
		Precip string `json:"precip"` // 5分钟累计降水量，单位毫米
		Type   string `json:"type"`   // 降水类型：rain = 雨，snow = 雪
	} `json:"minutely"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，或数据源说明，可能为空
		License []string `json:"license"` // 数据许可或版权声明，可能为空
	} `json:"refer"`
}

type Minutely5MResult struct {
	Result Minutely5MResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newMinutely5MResult(result Minutely5MResponse, body []byte, http gorequest.Response) *Minutely5MResult {
	return &Minutely5MResult{Result: result, Body: body, Http: http}
}

// Minutely5M 分钟级降水
// https://dev.qweather.com/docs/api/weather/weather-daily-forecast/
func (c *Client) Minutely5M(ctx context.Context, location string, notMustParams ...gorequest.Params) (*Minutely5MResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("location", location)
	params.Set("key", c.key)
	// 请求
	request, err := c.request(ctx, "minutely/5m", params, http.MethodGet)
	if err != nil {
		return newMinutely5MResult(Minutely5MResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response Minutely5MResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newMinutely5MResult(response, request.ResponseBody, request), err
}
