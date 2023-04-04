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
func (c *Client) Realtime(ctx context.Context, location string, notMustParams ...gorequest.Params) (*RealtimeResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, c.getApiUrl()+fmt.Sprintf("/%s/realtime", location), params, http.MethodGet)
	if err != nil {
		return newRealtimeResult(RealtimeResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RealtimeResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRealtimeResult(response, request.ResponseBody, request), err
}

// GetUltravioletDesc 紫外线 https://docs.caiyunapp.com/docs/tables/lifeindex
func (RealtimeResult) GetUltravioletDesc(ultraviolet float64) string {
	if ultraviolet <= 0 {
		return "无"
	} else if ultraviolet <= 1 {
		return "很弱"
	} else if ultraviolet <= 2 {
		return "很弱"
	} else if ultraviolet <= 3 {
		return "弱"
	} else if ultraviolet <= 4 {
		return "弱"
	} else if ultraviolet <= 5 {
		return "中等"
	} else if ultraviolet <= 6 {
		return "中等"
	} else if ultraviolet <= 7 {
		return "强"
	} else if ultraviolet <= 8 {
		return "强"
	} else if ultraviolet <= 9 {
		return "强"
	} else if ultraviolet <= 10 {
		return "很强"
	} else if ultraviolet <= 11 {
		return "极强"
	}
	return "无"
}
