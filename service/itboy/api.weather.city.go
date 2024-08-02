package itboy

import (
	"context"
	"go.dtapp.net/library/utils/godecimal"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
	"regexp"
)

type ApiWeatherCityResponse struct {
	Message  string `json:"message"` // 返回message
	Status   int    `json:"status"`  // 返回状态
	Date     string `json:"date"`    // 当前天气的当天日期
	Time     string `json:"time"`    // 系统更新时间
	CityInfo struct {
		City       string `json:"city"`       // 请求城市
		Citykey    string `json:"citykey"`    // 请求ID
		Parent     string `json:"parent"`     // 上级，一般是省份
		UpdateTime string `json:"updateTime"` // 天气更新时间
	} `json:"cityInfo"`
	Data struct {
		Shidu    string  `json:"shidu"`   // 湿度
		Pm25     float64 `json:"pm25"`    // pm2.5
		Pm10     float64 `json:"pm10"`    // pm10
		Quality  string  `json:"quality"` // 空气质量
		Wendu    string  `json:"wendu"`   // 温度
		Ganmao   string  `json:"ganmao"`  // 感冒提醒（指数）
		Forecast []struct {
			Date    string `json:"date"`    // 日期
			High    string `json:"high"`    // 高温
			Low     string `json:"low"`     // 低温
			Ymd     string `json:"ymd"`     // 年月日
			Week    string `json:"week"`    // 星期
			Sunrise string `json:"sunrise"` // 日出
			Sunset  string `json:"sunset"`  // 日落
			Aqi     int64  `json:"aqi"`     // 空气质量指数
			Fx      string `json:"fx"`      // 风向
			Fl      string `json:"fl"`      // 风级
			Type    string `json:"type"`    // 类型
			Notice  string `json:"notice"`  // 注意
		} `json:"forecast"` // 预测
		Yesterday struct {
			Date    string `json:"date"`    // 日期
			High    string `json:"high"`    // 高温
			Low     string `json:"low"`     // 低温
			Ymd     string `json:"ymd"`     // 年月日
			Week    string `json:"week"`    // 星期
			Sunrise string `json:"sunrise"` // 日出
			Sunset  string `json:"sunset"`  // 日落
			Aqi     int64  `json:"aqi"`     // 空气质量指数
			Fx      string `json:"fx"`      // 风向
			Fl      string `json:"fl"`      // 风级
			Type    string `json:"type"`    // 类型
			Notice  string `json:"notice"`  // 注意
		} `json:"yesterday"` // 昨天
	} `json:"data"`
}

type ApiWeatherCityResult struct {
	Result ApiWeatherCityResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newApiWeatherCityResult(result ApiWeatherCityResponse, body []byte, http gorequest.Response) *ApiWeatherCityResult {
	return &ApiWeatherCityResult{Result: result, Body: body, Http: http}
}

// ApiWeatherCity 国内天气
// https://www.sojson.com/blog/305.html
func (c *Client) ApiWeatherCity(ctx context.Context, cityCode string, notMustParams ...gorequest.Params) (*ApiWeatherCityResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, "api/weather/city/"+cityCode, params, http.MethodGet)
	if err != nil {
		return newApiWeatherCityResult(ApiWeatherCityResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ApiWeatherCityResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newApiWeatherCityResult(response, request.ResponseBody, request), err
}

func ApiWeatherCityLow(low string) float64 {
	re := regexp.MustCompile(`-?\d+`)
	result := re.FindString(low)
	return godecimal.NewString(result).Float64()
}

func ApiWeatherCityHigh(high string) float64 {
	re := regexp.MustCompile(`-?\d+`)
	result := re.FindString(high)
	return godecimal.NewString(result).Float64()
}
