package _map

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WeatherV1Response struct {
	Status int64 `json:"status"`
	Result struct {
		Location struct {
			Country  string `json:"country"`  // 国家名称
			Province string `json:"province"` // 省份名称
			City     string `json:"city"`     // 城市名称
			Name     string `json:"name"`     // 区县名称
			Id       string `json:"id"`       // 区县id
		} `json:"location"` // 地理位置信息
		Now struct {
			Text      string  `json:"text"`       // 天气现象
			Temp      int64   `json:"temp"`       // 温度（℃）
			FeelsLike int64   `json:"feels_like"` // 体感温度(℃)
			Rh        int64   `json:"rh"`         // 相对湿度(%)
			WindClass string  `json:"wind_class"` // 风力等级
			WindDir   string  `json:"wind_dir"`   // 风向描述
			Prec1h    float64 `json:"prec_1h"`    // 1小时累计降水量(mm)
			Clouds    int64   `json:"clouds"`     // 云量(%)
			Vis       int64   `json:"vis"`        // 能见度(m)
			Aqi       int64   `json:"aqi"`        // 空气质量指数数值
			Pm25      int64   `json:"pm25"`       // pm2.5浓度(μg/m3)
			Pm10      int64   `json:"pm10"`       // pm10浓度(μg/m3)
			No2       int64   `json:"no2"`        // 二氧化氮浓度(μg/m3)
			So2       int64   `json:"so2"`        // 二氧化硫浓度(μg/m3)
			O3        int64   `json:"o3"`         // 臭氧浓度(μg/m3)
			Co        float64 `json:"co"`         // 一氧化碳浓度(mg/m3)
			Uptime    string  `json:"uptime"`     // 数据更新时间，北京时间
		} `json:"now"` // 实况数据
		Alert []struct {
			Type  string `json:"type"`  // 预警事件类型
			Level string `json:"level"` // 预警事件等级
			Title string `json:"title"` // 预警标题
			Desc  string `json:"desc"`  // 预警详细提示信息
		} `json:"alert"` // 气象预警数据
		Indexes []struct {
			Name   string `json:"name"`   // 生活指数中文名称
			Brief  string `json:"brief"`  // 生活指数概要说明
			Detail string `json:"detail"` // 生活指数详细说明
		} `json:"indexes"` // 生活指数数据
		Forecasts []struct {
			Date      string `json:"date"`       // 日期，北京时区
			Week      string `json:"week"`       // 星期，北京时区
			High      int64  `json:"high"`       // 最高温度(℃)
			Low       int64  `json:"low"`        // 最低温度(℃)
			WcDay     string `json:"wc_day"`     // 白天风力
			WcNight   string `json:"wc_night"`   // 晚上风力
			WdDay     string `json:"wd_day"`     // 白天风向
			WdNight   string `json:"wd_night"`   // 晚上风向
			TextDay   string `json:"text_day"`   // 白天天气现象
			TextNight string `json:"text_night"` // 晚上天气现象
		} `json:"forecasts"` // 预报数据
		ForecastHours []struct {
			Text      string  `json:"text"`       // 天气现象
			TempFc    int64   `json:"temp_fc"`    // 温度(℃)
			WindClass string  `json:"wind_class"` // 风力等级
			WindDir   string  `json:"wind_dir"`   // 风向描述
			Rh        int64   `json:"rh"`         // 相对湿度
			Prec1h    float64 `json:"prec_1h"`    // 1小时累计降水量(mm)
			Clouds    int64   `json:"clouds"`     // 云量(%)
			DataTime  string  `json:"data_time"`  // 数据时间
		} `json:"forecast_hours"` // 未来24小时逐小时预报
	} `json:"result"`
	Message string `json:"message"`
}

type WeatherV1Result struct {
	Result WeatherV1Response  // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newWeatherV1Result(result WeatherV1Response, body []byte, http gorequest.Response) *WeatherV1Result {
	return &WeatherV1Result{Result: result, Body: body, Http: http}
}

// WeatherV1 国内天气查询服务
// https://lbsyun.baidu.com/index.php?title=webapi/weather
func (c *Client) WeatherV1(ctx context.Context, notMustParams ...gorequest.Params) (*WeatherV1Result, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("ak", c.ak)       // 开发者密钥
	params.Set("output", "json") // 返回格式，目前支持json/xml
	// 请求
	request, err := c.request(ctx, "weather/v1/", params, http.MethodGet)
	if err != nil {
		return newWeatherV1Result(WeatherV1Response{}, request.ResponseBody, request), err
	}
	// 定义
	var response WeatherV1Response
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWeatherV1Result(response, request.ResponseBody, request), err
}
