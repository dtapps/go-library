package baidu

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WeatherResponse struct {
	Status int `json:"status"`
	Result struct {
		Location struct {
			Country  string `json:"country"`
			Province string `json:"province"`
			City     string `json:"city"`
			Name     string `json:"name"`
			Id       string `json:"id"`
		} `json:"location"`
		Now struct {
			Text      string `json:"text"`
			Temp      int    `json:"temp"`
			FeelsLike int    `json:"feels_like"`
			Rh        int    `json:"rh"`
			WindClass string `json:"wind_class"`
			WindDir   string `json:"wind_dir"`
			Uptime    string `json:"uptime"`
		} `json:"now"`
		Forecasts []struct {
			TextDay   string `json:"text_day"`
			TextNight string `json:"text_night"`
			High      int    `json:"high"`
			Low       int    `json:"low"`
			WcDay     string `json:"wc_day"`
			WdDay     string `json:"wd_day"`
			WcNight   string `json:"wc_night"`
			WdNight   string `json:"wd_night"`
			Date      string `json:"date"`
			Week      string `json:"week"`
		} `json:"forecasts"`
	} `json:"result"`
	Message string `json:"message"`
}

type WeatherResult struct {
	Result WeatherResponse    // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newWeatherResult(result WeatherResponse, body []byte, http gorequest.Response) *WeatherResult {
	return &WeatherResult{Result: result, Body: body, Http: http}
}

// Weather 国内天气查询服务
// https://lbsyun.baidu.com/index.php?title=webapi/weather
func (c *Client) Weather(ctx context.Context, districtId string, notMustParams ...gorequest.Params) (*WeatherResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("ak", c.GetAk())
	params.Set("district_id", districtId)
	params.Set("output", "json")
	// 请求
	request, err := c.request(ctx, apiUrl+"/weather/v1/", params, http.MethodGet)
	if err != nil {
		return newWeatherResult(WeatherResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WeatherResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWeatherResult(response, request.ResponseBody, request), err
}
