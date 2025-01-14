package qweather

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type Weather72HResponse struct {
	Code       string `json:"code"`       // 状态码
	UpdateTime string `json:"updateTime"` // 最近更新时间
	FxLink     string `json:"fxLink"`     // 响应式页面
	Hourly     []struct {
		FxTime    string `json:"fxTime"`    // 预报时间
		Temp      string `json:"temp"`      // 温度，默认单位：摄氏度
		Icon      string `json:"icon"`      // 天气状况的图标代码
		Text      string `json:"text"`      // 天气状况的文字描述，包括阴晴雨雪等天气状态的描述
		Wind360   string `json:"wind360"`   // 风向360角度
		WindDir   string `json:"windDir"`   // 风向
		WindScale string `json:"windScale"` // 风力等级
		WindSpeed string `json:"windSpeed"` // 风速，公里/小时
		Humidity  string `json:"humidity"`  // 相对湿度，百分比数值
		Precip    string `json:"precip"`    // 当前小时累计降水量，默认单位：毫米
		Pop       string `json:"pop"`       // 逐小时预报降水概率，百分比数值，可能为空
		Pressure  string `json:"pressure"`  // 大气压强，默认单位：百帕
		Cloud     string `json:"cloud"`     // 云量，百分比数值。可能为空
		Dew       string `json:"dew"`       //  露点温度。可能为空
	} `json:"hourly"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，或数据源说明，可能为空
		License []string `json:"license"` // 数据许可或版权声明，可能为空
	} `json:"refer"`
}

type Weather72HResult struct {
	Result Weather72HResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newWeather72HResult(result Weather72HResponse, body []byte, http gorequest.Response) *Weather72HResult {
	return &Weather72HResult{Result: result, Body: body, Http: http}
}

// Weather72H 逐小时天气预报
// https://dev.qweather.com/docs/api/weather/weather-hourly-forecast/
func (c *Client) Weather72H(ctx context.Context, location string, notMustParams ...gorequest.Params) (*Weather72HResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("location", location)
	params.Set("key", c.key)
	// 请求
	request, err := c.request(ctx, "weather/72h", params, http.MethodGet)
	if err != nil {
		return newWeather72HResult(Weather72HResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response Weather72HResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWeather72HResult(response, request.ResponseBody, request), err
}
