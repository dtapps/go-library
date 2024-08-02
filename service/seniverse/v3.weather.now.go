package seniverse

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"time"
)

type V3WeatherNowResponse struct {
	Results []struct {
		Location struct {
			Id             string `json:"id"`
			Name           string `json:"name"`
			Country        string `json:"country"`
			Path           string `json:"path"`
			Timezone       string `json:"timezone"`
			TimezoneOffset string `json:"timezone_offset"`
		} `json:"location"`
		Now struct {
			Text                string `json:"text"`                  // 天气现象文字
			Code                string `json:"code"`                  // 天气现象代码
			Temperature         string `json:"temperature"`           // 温度，单位为c摄氏度或f华氏度
			FeelsLike           string `json:"feels_like"`            // 体感温度，单位为c摄氏度或f华氏度
			Pressure            string `json:"pressure"`              // 气压，单位为mb百帕或in英寸
			Humidity            string `json:"humidity"`              // 相对湿度，0~100，单位为百分比
			Visibility          string `json:"visibility"`            // 能见度，单位为km公里或mi英里
			WindDirection       string `json:"wind_direction"`        // 风向文字
			WindDirectionDegree string `json:"wind_direction_degree"` // 风向角度，范围0~360
			WindSpeed           string `json:"wind_speed"`            // 风速，单位km/h（当unit=c时）、mph（当unit=f时）
			WindScale           string `json:"wind_scale"`            // 风力等级
			Clouds              string `json:"clouds"`                // 云量，单位%，范围0~100，天空被云覆盖的百分比 #目前不支持中国城市#
			DewPoint            string `json:"dew_point"`             // 露点温度，请参考：http://baike.baidu.com/view/118348.htm #目前不支持中国城市#
		} `json:"now"`
		LastUpdate time.Time `json:"last_update"`
	} `json:"results"`
}

type V3WeatherNowResult struct {
	Result V3WeatherNowResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求·
}

func newV3WeatherNowResult(result V3WeatherNowResponse, body []byte, http gorequest.Response) *V3WeatherNowResult {
	return &V3WeatherNowResult{Result: result, Body: body, Http: http}
}

// WeatherNow 天气实况
// location 所查询的位置
// https://seniverse.yuque.com/hyper_data/api_v3/nyiu3t
func (c *V3Client) WeatherNow(ctx context.Context, location string, notMustParams ...gorequest.Params) (*V3WeatherNowResult, ApiError, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("location", location)
	// 请求
	request, err := c.request(ctx, "weather/now.json", params)
	if err != nil {
		return newV3WeatherNowResult(V3WeatherNowResponse{}, request.ResponseBody, request), ApiError{}, err
	}
	// 定义
	var response V3WeatherNowResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = gojson.Unmarshal(request.ResponseBody, &apiError)
	return newV3WeatherNowResult(response, request.ResponseBody, request), apiError, err
}
