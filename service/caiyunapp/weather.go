package caiyunapp

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WeatherResponse struct {
	Status     string    `json:"status"`
	ApiVersion string    `json:"api_version"`
	ApiStatus  string    `json:"api_status"`
	Lang       string    `json:"lang"`
	Unit       string    `json:"unit"`
	Tzshift    int       `json:"tzshift"`
	Timezone   string    `json:"timezone"`
	ServerTime int       `json:"server_time"`
	Location   []float64 `json:"location"`
	Result     struct {
		Alert struct {
		} `json:"alert"`
		Realtime struct {
		} `json:"realtime"`
		Minutely struct {
		} `json:"minutely"`
		Hourly struct {
		} `json:"hourly"`
		Daily struct {
		} `json:"daily"`
		Primary          int    `json:"primary"`
		ForecastKeypoint string `json:"forecast_keypoint"`
	} `json:"result"`
}

type WeatherResult struct {
	Result WeatherResponse    // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newWeatherResult(result WeatherResponse, body []byte, http gorequest.Response) *WeatherResult {
	return &WeatherResult{Result: result, Body: body, Http: http}
}

// Weather 综合
// https://docs.caiyunapp.com/docs/weather
func (c *Client) Weather(ctx context.Context, locationLongitude, locationLatitude string, notMustParams ...gorequest.Params) (*WeatherResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, c.getApiUrl()+fmt.Sprintf("/%s,%s/weather", locationLatitude, locationLongitude), params, http.MethodGet)
	if err != nil {
		return newWeatherResult(WeatherResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WeatherResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWeatherResult(response, request.ResponseBody, request), err
}
