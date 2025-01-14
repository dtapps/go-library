package cma

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type ApiWeatherViewResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data struct {
		Location struct {
			Id        string  `json:"id"`
			Name      string  `json:"name"`
			Path      string  `json:"path"`
			Longitude float64 `json:"longitude"`
			Latitude  float64 `json:"latitude"`
			Timezone  int     `json:"timezone"`
		} `json:"location"`
		Daily []struct {
			Date               string  `json:"date"`
			High               float64 `json:"high"`
			DayText            string  `json:"dayText"`
			DayCode            float64 `json:"dayCode"`
			DayWindDirection   string  `json:"dayWindDirection"`
			DayWindScale       string  `json:"dayWindScale"`
			Low                float64 `json:"low"`
			NightText          string  `json:"nightText"`
			NightCode          float64 `json:"nightCode"`
			NightWindDirection string  `json:"nightWindDirection"`
			NightWindScale     string  `json:"nightWindScale"`
		} `json:"daily"`
		Now struct {
			Precipitation       float64 `json:"precipitation"`
			Temperature         float64 `json:"temperature"`
			Pressure            float64 `json:"pressure"`
			Humidity            float64 `json:"humidity"`
			WindDirection       string  `json:"windDirection"`
			WindDirectionDegree float64 `json:"windDirectionDegree"`
			WindSpeed           float64 `json:"windSpeed"`
			WindScale           string  `json:"windScale"`
		} `json:"now"`
		Alarm      []interface{} `json:"alarm"`
		LastUpdate string        `json:"lastUpdate"`
	} `json:"data"`
}

type ApiWeatherViewResult struct {
	Result ApiWeatherViewResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newApiWeatherViewResult(result ApiWeatherViewResponse, body []byte, http gorequest.Response) *ApiWeatherViewResult {
	return &ApiWeatherViewResult{Result: result, Body: body, Http: http}
}

func (c *Client) ApiWeatherView(ctx context.Context, stationid string, notMustParams ...gorequest.Params) (*ApiWeatherViewResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("stationid", stationid)
	// 请求
	request, err := c.request(ctx, "api/weather/view", params, http.MethodGet)
	if err != nil {
		return newApiWeatherViewResult(ApiWeatherViewResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ApiWeatherViewResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newApiWeatherViewResult(response, request.ResponseBody, request), err
}
