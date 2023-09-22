package amap

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WeatherWeatherInfoResponse struct {
	Status   string `json:"status"`   // 值为0或1 1：成功；0：失败
	Count    string `json:"count"`    // 返回结果总数目
	Info     string `json:"info"`     // 返回的状态信息
	Infocode string `json:"infocode"` // 返回状态说明,10000代表正确
	Lives    []struct {
		Province         string `json:"province"`          // 省份名
		City             string `json:"city"`              // 城市名
		Adcode           string `json:"adcode"`            // 区域编码
		Weather          string `json:"weather"`           // 天气现象（汉字描述）
		Temperature      string `json:"temperature"`       // 实时气温，单位：摄氏度
		Winddirection    string `json:"winddirection"`     // 风向描述
		Windpower        string `json:"windpower"`         // 风力级别，单位：级
		Humidity         string `json:"humidity"`          // 空气湿度
		Reporttime       string `json:"reporttime"`        // 数据发布的时间
		TemperatureFloat string `json:"temperature_float"` // 温度
		HumidityFloat    string `json:"humidity_float"`    // 湿度
	} `json:"lives"` // 实况天气数据信息
	Forecasts []struct {
		City       string `json:"city"`       // 城市名称
		Adcode     string `json:"adcode"`     // 城市编码
		Province   string `json:"province"`   // 省份名称
		Reporttime string `json:"reporttime"` // 预报发布时间
		Casts      []struct {
			Date           string `json:"date"`            // 日期
			Week           string `json:"week"`            // 星期几
			Dayweather     string `json:"dayweather"`      // 白天天气现象
			Nightweather   string `json:"nightweather"`    // 晚上天气现象
			Daytemp        string `json:"daytemp"`         // 白天温度
			Nighttemp      string `json:"nighttemp"`       // 晚上温度
			Daywind        string `json:"daywind"`         // 白天风向
			Nightwind      string `json:"nightwind"`       // 晚上风向
			Daypower       string `json:"daypower"`        // 白天风力
			Nightpower     string `json:"nightpower"`      // 晚上风力
			DaytempFloat   string `json:"daytemp_float"`   // 日温度
			NighttempFloat string `json:"nighttemp_float"` // 夜间温度
		} `json:"casts"` // 预报数据list结构，元素cast,按顺序为当天、第二天、第三天的预报数据
	} `json:"forecasts"` // 预报天气信息数据
}

type WeatherWeatherInfoResult struct {
	Result WeatherWeatherInfoResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
}

func newWeatherWeatherInfoResult(result WeatherWeatherInfoResponse, body []byte, http gorequest.Response) *WeatherWeatherInfoResult {
	return &WeatherWeatherInfoResult{Result: result, Body: body, Http: http}
}

// WeatherWeatherInfo 天气查询
// https://lbs.amap.com/api/webservice/guide/api/weatherinfo
func (c *Client) WeatherWeatherInfo(ctx context.Context, notMustParams ...*gorequest.Params) (*WeatherWeatherInfoResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("key", c.GetKey())
	params.Set("output", "JSON")
	// 请求
	request, err := c.request(ctx, apiUrl+"/weather/weatherInfo", params, http.MethodGet)
	if err != nil {
		return newWeatherWeatherInfoResult(WeatherWeatherInfoResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WeatherWeatherInfoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWeatherWeatherInfoResult(response, request.ResponseBody, request), err
}
