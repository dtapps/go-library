package qweather

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type Weather3DResponse struct {
	Code       string `json:"code"`       // 状态码
	UpdateTime string `json:"updateTime"` // 最近更新时间
	FxLink     string `json:"fxLink"`     // 响应式页面
	Daily      []struct {
		FxDate         string `json:"fxDate"`         // 预报日期
		Sunrise        string `json:"sunrise"`        // 日出时间，在高纬度地区可能为空
		Sunset         string `json:"sunset"`         // 日落时间，在高纬度地区可能为空
		Moonrise       string `json:"moonrise"`       // 当天月升时间，可能为空
		Moonset        string `json:"moonset"`        // 当天月落时间，可能为空
		MoonPhase      string `json:"moonPhase"`      // 月相名称
		MoonPhaseIcon  string `json:"moonPhaseIcon"`  // 月相图标代码
		TempMax        string `json:"tempMax"`        // 预报当天最高温度
		TempMin        string `json:"tempMin"`        // 预报当天最低温度
		IconDay        string `json:"iconDay"`        // 预报白天天气状况的图标代码
		TextDay        string `json:"textDay"`        // 预报白天天气状况文字描述
		IconNight      string `json:"iconNight"`      // 预报夜间天气状况的图标代码
		TextNight      string `json:"textNight"`      // 预报晚间天气状况文字描述
		Wind360Day     string `json:"wind360Day"`     // 预报白天风向360角度
		WindDirDay     string `json:"windDirDay"`     // 预报白天风向
		WindScaleDay   string `json:"windScaleDay"`   // 预报白天风力等级
		WindSpeedDay   string `json:"windSpeedDay"`   // 预报白天风速，公里/小时
		Wind360Night   string `json:"wind360Night"`   // 预报夜间风向360角度
		WindDirNight   string `json:"windDirNight"`   // 预报夜间当天风向
		WindScaleNight string `json:"windScaleNight"` // 预报夜间风力等级
		WindSpeedNight string `json:"windSpeedNight"` // 预报夜间风速，公里/小时
		Precip         string `json:"precip"`         // 预报当天总降水量，默认单位：毫米
		UvIndex        string `json:"uvIndex"`        // 紫外线强度指数
		Humidity       string `json:"humidity"`       // 相对湿度，百分比数值
		Pressure       string `json:"pressure"`       // 大气压强，默认单位：百帕
		Vis            string `json:"vis"`            // 能见度，默认单位：公里
		Cloud          string `json:"cloud"`          // 云量，百分比数值。可能为空
	} `json:"daily"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，或数据源说明，可能为空
		License []string `json:"license"` // 数据许可或版权声明，可能为空
	} `json:"refer"`
}

type Weather3DResult struct {
	Result Weather3DResponse  // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newWeather3DResult(result Weather3DResponse, body []byte, http gorequest.Response) *Weather3DResult {
	return &Weather3DResult{Result: result, Body: body, Http: http}
}

// Weather3D 每日天气预报
// https://dev.qweather.com/docs/api/weather/weather-daily-forecast/
func (c *Client) Weather3D(ctx context.Context, location string, notMustParams ...gorequest.Params) (*Weather3DResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("location", location)
	params.Set("key", c.key)
	// 请求
	request, err := c.request(ctx, "weather/3d", params, http.MethodGet)
	if err != nil {
		return newWeather3DResult(Weather3DResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response Weather3DResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWeather3DResult(response, request.ResponseBody, request), err
}
