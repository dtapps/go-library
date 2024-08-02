package seniverse

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"time"
)

type V3WeatherDailyResponse struct {
	Results []struct {
		Location struct {
			Id             string `json:"id"`
			Name           string `json:"name"`
			Country        string `json:"country"`
			Path           string `json:"path"`
			Timezone       string `json:"timezone"`
			TimezoneOffset string `json:"timezone_offset"`
		} `json:"location"`
		Daily []struct {
			Date                string `json:"date"`                  // 日期（该城市的本地时间）
			TextDay             string `json:"text_day"`              // 白天天气现象文字
			CodeDay             string `json:"code_day"`              // 白天天气现象代码
			TextNight           string `json:"text_night"`            // 晚间天气现象文字
			CodeNight           string `json:"code_night"`            // 晚间天气现象代码
			High                string `json:"high"`                  // 当天最高温度
			Low                 string `json:"low"`                   // 当天最低温度
			Precip              string `json:"precip"`                // 降水概率，范围0~100，单位百分比（目前仅支持国外城市）
			WindDirection       string `json:"wind_direction"`        // 风向文字
			WindDirectionDegree string `json:"wind_direction_degree"` // 风向角度，范围0~360
			WindSpeed           string `json:"wind_speed"`            // 风速，单位km/h（当unit=c时）、mph（当unit=f时）
			WindScale           string `json:"wind_scale"`            // 风力等级
			Rainfall            string `json:"rainfall"`              // 降水量，单位mm
			Humidity            string `json:"humidity"`              // 相对湿度，0~100，单位为百分比
		} `json:"daily"`
		LastUpdate time.Time `json:"last_update"`
	} `json:"results"`
}

type V3WeatherDailyResult struct {
	Result V3WeatherDailyResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求·
}

func newV3WeatherDailyResult(result V3WeatherDailyResponse, body []byte, http gorequest.Response) *V3WeatherDailyResult {
	return &V3WeatherDailyResult{Result: result, Body: body, Http: http}
}

// WeatherDaily 未来15天逐日天气预报和昨日天气
// https://seniverse.yuque.com/hyper_data/api_v3/sl6gvt
func (c *V3Client) WeatherDaily(ctx context.Context, location string, notMustParams ...gorequest.Params) (*V3WeatherDailyResult, ApiError, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("location", location)
	// 请求
	request, err := c.request(ctx, "weather/daily.json", params)
	if err != nil {
		return newV3WeatherDailyResult(V3WeatherDailyResponse{}, request.ResponseBody, request), ApiError{}, err
	}
	// 定义
	var response V3WeatherDailyResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = gojson.Unmarshal(request.ResponseBody, &apiError)
	return newV3WeatherDailyResult(response, request.ResponseBody, request), apiError, err
}
