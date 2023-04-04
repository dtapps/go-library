package caiyunapp

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type HourlyResponse struct {
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
		Primary          float64 `json:"primary"`
		ForecastKeypoint string  `json:"forecast_keypoint"`
	} `json:"result"`
}

type HourlyResult struct {
	Result HourlyResponse     // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newHourlyResult(result HourlyResponse, body []byte, http gorequest.Response) *HourlyResult {
	return &HourlyResult{Result: result, Body: body, Http: http}
}

// Hourly 小时级别预报
// https://docs.caiyunapp.com/docs/hourly
func (c *Client) Hourly(ctx context.Context, location string, notMustParams ...gorequest.Params) (*HourlyResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, c.getApiUrl()+fmt.Sprintf("/%s/hourly", location), params, http.MethodGet)
	if err != nil {
		return newHourlyResult(HourlyResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response HourlyResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newHourlyResult(response, request.ResponseBody, request), err
}
