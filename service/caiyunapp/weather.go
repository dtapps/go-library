package caiyunapp

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
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
			Status  string `json:"status"`
			Content []struct {
				Pubtimestamp  int       `json:"pubtimestamp"` // 发布时间，单位是 Unix 时间戳
				AlertID       string    `json:"alertId"`      // 预警 ID
				Status        string    `json:"status"`       // 预警信息的状态
				Adcode        string    `json:"adcode"`       // 区域代码
				Location      string    `json:"location"`     // 位置
				Province      string    `json:"province"`     // 省
				City          string    `json:"city"`         // 市
				County        string    `json:"county"`       // 县
				Code          string    `json:"code"`         // 预警代码
				Source        string    `json:"source"`       // 发布单位
				Title         string    `json:"title"`        // 标题
				Description   string    `json:"description"`  // 描述
				RegionID      string    `json:"regionId"`
				Latlon        []float64 `json:"latlon"`
				RequestStatus string    `json:"request_status"`
			} `json:"content"`
			Adcodes []struct {
				Adcode int    `json:"adcode"`
				Name   string `json:"name"`
			} `json:"adcodes"` // 行政区划层级信息
		} `json:"alert"` // 预警数据
		Realtime struct {
			Temperature float64 `json:"temperature"` // 地表2米气温
			Humidity    float64 `json:"humidity"`    // 地表2米湿度相对湿度(%)
			Cloudrate   float64 `json:"cloudrate"`   // 总云量(0.0-1.0)
			Skycon      string  `json:"skycon"`      // 天气现象
			SkyconDesc  string  `json:"skycon_desc"` // 天气现象
			Visibility  float64 `json:"visibility"`  // 地表水平能见度
			Dswrf       float64 `json:"dswrf"`       // 向下短波辐射通量(W/M2)
			Wind        struct {
				Speed     float64 `json:"speed"`     // 地表 10 米风速
				Direction float64 `json:"direction"` // 地表 10 米风向
			} `json:"wind"`
			Pressure            float64 `json:"pressure"`             // 地面气压
			ApparentTemperature float64 `json:"apparent_temperature"` // 体感温度
			Precipitation       struct {
				Local struct {
					Status     string  `json:"status"`
					Datasource string  `json:"datasource"` // 本地降水带与本地的距离
					Intensity  float64 `json:"intensity"`  // 本地降水处的降水强度
				} `json:"local"` // 本地降水强度
				Nearest struct {
					Status    string  `json:"status"`
					Distance  float64 `json:"distance"`  // 最近降水带与本地的距离
					Intensity float64 `json:"intensity"` // 最近降水处的降水强度
				} `json:"nearest"` // 最近降水强度
			} `json:"precipitation"`
			AirQuality struct {
				Pm25 float64 `json:"pm25"` // PM25 浓度(μg/m3)
				Pm10 float64 `json:"pm10"` // PM10 浓度(μg/m3)
				O3   float64 `json:"o3"`   // 臭氧浓度(μg/m3)
				So2  float64 `json:"so2"`  // 二氧化氮浓度(μg/m3)
				No2  float64 `json:"no2"`  // 二氧化硫浓度(μg/m3)
				Co   float64 `json:"co"`   // 一氧化碳浓度(mg/m3)
				Aqi  struct {
					Chn float64 `json:"chn"` // 国标 AQI
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
					Desc  string  `json:"desc"` // 生活指数
				} `json:"ultraviolet"`
				Comfort struct {
					Index float64 `json:"index"`
					Desc  string  `json:"desc"` // 生活指数
				} `json:"comfort"`
			} `json:"life_index"`
		} `json:"realtime"` // 实况
		Minutely struct {
			Status          string    `json:"status"`
			Datasource      string    `json:"datasource"`
			Precipitation2H []float64 `json:"precipitation_2h"`
			Precipitation   []float64 `json:"precipitation"`
			Probability     []float64 `json:"probability"`
			Description     string    `json:"description"`
		} `json:"minutely"` // 分钟级预报
		Hourly struct {
			Status        string `json:"status"`
			Description   string `json:"description"`
			Precipitation []struct {
				Datetime    string  `json:"datetime"`
				Value       float64 `json:"value"`
				Probability float64 `json:"probability"`
			} `json:"precipitation"`
			Temperature []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"` // 地表 2 米气温
			} `json:"temperature"`
			ApparentTemperature []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"` // 体感温度
			} `json:"apparent_temperature"`
			Wind []struct {
				Datetime  string  `json:"datetime"`
				Speed     float64 `json:"speed"`
				Direction float64 `json:"direction"`
			} `json:"wind"`
			Humidity []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"` // 地表 2 米相对湿度(%)
			} `json:"humidity"`
			Cloudrate []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"`
			} `json:"cloudrate"`
			Skycon []struct {
				Datetime string `json:"datetime"`
				Value    string `json:"value"`
			} `json:"skycon"`
			Pressure []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"` // 地面气压
			} `json:"pressure"`
			Visibility []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"`
			} `json:"visibility"`
			Dswrf []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"`
			} `json:"dswrf"`
			AirQuality struct {
				Aqi []struct {
					Datetime string `json:"datetime"`
					Value    struct {
						Chn float64 `json:"chn"`
						Usa float64 `json:"usa"`
					} `json:"value"`
				} `json:"aqi"`
				Pm25 []struct {
					Datetime string  `json:"datetime"`
					Value    float64 `json:"value"`
				} `json:"pm25"`
			} `json:"air_quality"`
		} `json:"hourly"` // 小时级别预报
		Daily struct {
			Status string `json:"status"`
			Astro  []struct {
				Date    string `json:"date"`
				Sunrise struct {
					Time string `json:"time"`
				} `json:"sunrise"`
				Sunset struct {
					Time string `json:"time"`
				} `json:"sunset"`
			} `json:"astro"` // 日出日落时间，当地时区的时刻，tzshift 不作用在这个变量)
			Precipitation08H20H []struct {
				Date        string  `json:"date"`
				Max         float64 `json:"max"`
				Min         float64 `json:"min"`
				Avg         float64 `json:"avg"`
				Probability float64 `json:"probability"`
			} `json:"precipitation_08h_20h"` // 白天降水数据
			Precipitation20H32H []struct {
				Date        string  `json:"date"`
				Max         float64 `json:"max"`
				Min         float64 `json:"min"`
				Avg         float64 `json:"avg"`
				Probability float64 `json:"probability"`
			} `json:"precipitation_20h_32h"` // 夜晚降水数据
			Precipitation []struct {
				Date        string  `json:"date"`
				Max         float64 `json:"max"`
				Min         float64 `json:"min"`
				Avg         float64 `json:"avg"`
				Probability float64 `json:"probability"`
			} `json:"precipitation"` // 降水数据
			Temperature []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"temperature"` // 全天地表 2 米气温
			Temperature08H20H []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"temperature_08h_20h"` // 白天地表 2 米气温
			Temperature20H32H []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"temperature_20h_32h"` // 夜晚地表 2 米气温
			Wind []struct {
				Date string `json:"date"`
				Max  struct {
					Speed     float64 `json:"speed"` // 全天地表 10 米风速
					Direction float64 `json:"direction"`
				} `json:"max"`
				Min struct {
					Speed     float64 `json:"speed"` // 全天地表 10 米风速
					Direction float64 `json:"direction"`
				} `json:"min"`
				Avg struct {
					Speed     float64 `json:"speed"` // 全天地表 10 米风速
					Direction float64 `json:"direction"`
				} `json:"avg"`
			} `json:"wind"`
			Wind08H20H []struct {
				Date string `json:"date"`
				Max  struct {
					Speed     float64 `json:"speed"` // 白天地表 10 米风速
					Direction float64 `json:"direction"`
				} `json:"max"`
				Min struct {
					Speed     float64 `json:"speed"` // 白天地表 10 米风速
					Direction float64 `json:"direction"`
				} `json:"min"`
				Avg struct {
					Speed     float64 `json:"speed"` // 白天地表 10 米风速
					Direction float64 `json:"direction"`
				} `json:"avg"`
			} `json:"wind_08h_20h"`
			Wind20H32H []struct {
				Date string `json:"date"`
				Max  struct {
					Speed     float64 `json:"speed"` // 夜晚地表 10 米风速
					Direction float64 `json:"direction"`
				} `json:"max"`
				Min struct {
					Speed     float64 `json:"speed"` // 夜晚地表 10 米风速
					Direction float64 `json:"direction"`
				} `json:"min"`
				Avg struct {
					Speed     float64 `json:"speed"` // 夜晚地表 10 米风速
					Direction float64 `json:"direction"`
				} `json:"avg"`
			} `json:"wind_20h_32h"`
			Humidity []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"humidity"` // 地表 2 米相对湿度(%)
			Cloudrate []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"cloudrate"` // 云量(0.0-1.0)
			Pressure []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"pressure"` // 地面气压
			Visibility []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"visibility"` // 地表水平能见度
			Dswrf []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"dswrf"` // 向下短波辐射通量(W/M2)
			AirQuality struct {
				Aqi []struct {
					Date string `json:"date"`
					Max  struct {
						Chn float64 `json:"chn"`
						Usa float64 `json:"usa"`
					} `json:"max"`
					Avg struct {
						Chn float64 `json:"chn"`
						Usa float64 `json:"usa"`
					} `json:"avg"`
					Min struct {
						Chn float64 `json:"chn"`
						Usa float64 `json:"usa"`
					} `json:"min"`
				} `json:"aqi"` // 国标 AQI
				Pm25 []struct {
					Date string  `json:"date"`
					Max  float64 `json:"max"`
					Avg  float64 `json:"avg"`
					Min  float64 `json:"min"`
				} `json:"pm25"` // PM2.5 浓度(μg/m3)
			} `json:"air_quality"`
			Skycon []struct {
				Date  string `json:"date"`
				Value string `json:"value"` // 全天主要 天气现象
			} `json:"skycon"`
			Skycon08H20H []struct {
				Date  string `json:"date"`
				Value string `json:"value"` // 白天主要 天气现象
			} `json:"skycon_08h_20h"`
			Skycon20H32H []struct {
				Date  string `json:"date"`
				Value string `json:"value"` // 夜晚主要 天气现象
			} `json:"skycon_20h_32h"`
			LifeIndex struct {
				Ultraviolet []struct {
					Date  string `json:"date"`
					Index string `json:"index"`
					Desc  string `json:"desc"` // 紫外线指数自然语言
				} `json:"ultraviolet"`
				CarWashing []struct {
					Date  string `json:"date"`
					Index string `json:"index"`
					Desc  string `json:"desc"` // 洗车指数自然语言
				} `json:"carWashing"`
				Dressing []struct {
					Date  string `json:"date"`
					Index string `json:"index"`
					Desc  string `json:"desc"` // 穿衣指数自然语言
				} `json:"dressing"`
				Comfort []struct {
					Date  string `json:"date"`
					Index string `json:"index"`
					Desc  string `json:"desc"` // 舒适度指数自然语言
				} `json:"comfort"`
				ColdRisk []struct {
					Date  string `json:"date"`
					Index string `json:"index"`
					Desc  string `json:"desc"` // 感冒指数自然语言
				} `json:"coldRisk"`
			} `json:"life_index"`
		} `json:"daily"` // 天级别预报
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
func (c *Client) Weather(ctx context.Context, location string, notMustParams ...gorequest.Params) (*WeatherResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	url := apiUrl + c.token + "/weather"
	if location != "" {
		url = apiUrl + c.token + "/" + location + "/weather"
	}
	request, err := c.request(ctx, url, params, http.MethodGet)
	if err != nil {
		return newWeatherResult(WeatherResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WeatherResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWeatherResult(response, request.ResponseBody, request), err
}
