package caiyunapp

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type RealtimeResponse struct {
	Status     string    `json:"status"`
	ApiVersion string    `json:"api_version"`
	ApiStatus  string    `json:"api_status"`
	Lang       string    `json:"lang"`
	Unit       string    `json:"unit"`
	Tzshift    float64   `json:"tzshift"`
	Timezone   string    `json:"timezone"`
	ServerTime float64   `json:"server_time"`
	Location   []float64 `json:"location"`
	Result     struct {
		Realtime struct {
			Status      string  `json:"status"`
			Temperature float64 `json:"temperature"`
			Humidity    float64 `json:"humidity"`
			Cloudrate   float64 `json:"cloudrate"`
			Skycon      string  `json:"skycon"`
			Visibility  float64 `json:"visibility"`
			Dswrf       float64 `json:"dswrf"`
			Wind        struct {
				Speed     float64 `json:"speed"`
				Direction float64 `json:"direction"`
			} `json:"wind"`
			Pressure            float64 `json:"pressure"`
			ApparentTemperature float64 `json:"apparent_temperature"`
			Precipitation       struct {
				Local struct {
					Status     string  `json:"status"`
					Datasource string  `json:"datasource"`
					Intensity  float64 `json:"intensity"`
				} `json:"local"`
				Nearest struct {
					Status    string  `json:"status"`
					Distance  float64 `json:"distance"`
					Intensity float64 `json:"intensity"`
				} `json:"nearest"`
			} `json:"precipitation"`
			AirQuality struct {
				Pm25 float64 `json:"pm25"`
				Pm10 float64 `json:"pm10"`
				O3   float64 `json:"o3"`
				So2  float64 `json:"so2"`
				No2  float64 `json:"no2"`
				Co   float64 `json:"co"`
				Aqi  struct {
					Chn float64 `json:"chn"`
					Usa float64 `json:"usa"`
				} `json:"aqi"`
				Description struct {
					Chn string `json:"chn"`
					Usa string `json:"usa"`
				} `json:"description"`
			} `json:"air_quality"`
			LifeIndex struct {
				Ultraviolet struct {
					Index float64 `json:"index"`
					Desc  string  `json:"desc"`
				} `json:"ultraviolet"`
				Comfort struct {
					Index float64 `json:"index"`
					Desc  string  `json:"desc"`
				} `json:"comfort"`
			} `json:"life_index"`
		} `json:"realtime"`
		Primary float64 `json:"primary"`
	} `json:"result"`
}

type RealtimeResult struct {
	Result RealtimeResponse   // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newRealtimeResult(result RealtimeResponse, body []byte, http gorequest.Response) *RealtimeResult {
	return &RealtimeResult{Result: result, Body: body, Http: http}
}

// Realtime 实况
// https://docs.caiyunapp.com/docs/realtime
func (c *Client) Realtime(ctx context.Context, locationLongitude, locationLatitude string, notMustParams ...gorequest.Params) (*RealtimeResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, c.getApiUrl()+fmt.Sprintf("/%s,%s/realtime", locationLatitude, locationLongitude), params, http.MethodGet)
	if err != nil {
		return newRealtimeResult(RealtimeResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RealtimeResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRealtimeResult(response, request.ResponseBody, request), err
}
