package caiyunapp

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type DailyResponse struct {
	Status     string    `json:"status"`      // 状态
	ApiVersion string    `json:"api_version"` // api版本
	ApiStatus  string    `json:"api_status"`  // api状态
	Lang       string    `json:"lang"`
	Unit       string    `json:"unit"`
	Tzshift    float64   `json:"tzshift"`
	Timezone   string    `json:"timezone"`
	ServerTime float64   `json:"server_time"`
	Location   []float64 `json:"location"`
	Result     struct {
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
		} `json:"daily"`
		Primary float64 `json:"primary"`
	} `json:"result"`
}

type DailyResult struct {
	Result DailyResponse      // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newDailyResult(result DailyResponse, body []byte, http gorequest.Response) *DailyResult {
	return &DailyResult{Result: result, Body: body, Http: http}
}

// Daily 天级别预报
// https://docs.caiyunapp.com/docs/daily
func (c *Client) Daily(ctx context.Context, locationLatitude, locationLongitude string, notMustParams ...gorequest.Params) (*DailyResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, c.getApiUrl()+fmt.Sprintf("/%s,%s/daily?dailysteps=1", locationLatitude, locationLongitude), params, http.MethodGet)
	if err != nil {
		return newDailyResult(DailyResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response DailyResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newDailyResult(response, request.ResponseBody, request), err
}
