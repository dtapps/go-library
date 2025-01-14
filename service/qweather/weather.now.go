package qweather

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WeatherNowResponse struct {
	Code       string `json:"code"`       // 状态码
	UpdateTime string `json:"updateTime"` // 最近更新时间
	FxLink     string `json:"fxLink"`     // 响应式页面
	Now        struct {
		ObsTime   string `json:"obsTime"`   // 数据观测时间
		Temp      string `json:"temp"`      // 温度，默认单位：摄氏度
		FeelsLike string `json:"feelsLike"` // 体感温度，默认单位：摄氏度
		Icon      string `json:"icon"`      // 天气状况的图标代码
		Text      string `json:"text"`      // 天气状况的文字描述，包括阴晴雨雪等天气状态的描述
		Wind360   string `json:"wind360"`   // 风向360角度
		WindDir   string `json:"windDir"`   // 风向
		WindScale string `json:"windScale"` // 风力等级
		WindSpeed string `json:"windSpeed"` // 风速，公里/小时
		Humidity  string `json:"humidity"`  // 相对湿度，百分比数值
		Precip    string `json:"precip"`    // 当前小时累计降水量，默认单位：毫米
		Pressure  string `json:"pressure"`  // 大气压强，默认单位：百帕
		Vis       string `json:"vis"`       // 能见度，默认单位：公里
		Cloud     string `json:"cloud"`     // 云量，百分比数值。可能为空
		Dew       string `json:"dew"`       // 露点温度。可能为空
	} `json:"now"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，或数据源说明，可能为空
		License []string `json:"license"` // 数据许可或版权声明，可能为空
	} `json:"refer"`
}

type WeatherNowResult struct {
	Result WeatherNowResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newWeatherNowResult(result WeatherNowResponse, body []byte, http gorequest.Response) *WeatherNowResult {
	return &WeatherNowResult{Result: result, Body: body, Http: http}
}

// WeatherNow 实时天气
// location = 需要查询地区的LocationID或以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位），LocationID可通过GeoAPI获取。例如 location=101010100 或 location=116.41,39.92
// https://dev.qweather.com/docs/api/weather/weather-now/
func (c *Client) WeatherNow(ctx context.Context, location string, notMustParams ...gorequest.Params) (*WeatherNowResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("location", location)
	params.Set("key", c.key)
	// 请求
	request, err := c.request(ctx, "weather/now", params, http.MethodGet)
	if err != nil {
		return newWeatherNowResult(WeatherNowResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WeatherNowResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWeatherNowResult(response, request.ResponseBody, request), err
}
