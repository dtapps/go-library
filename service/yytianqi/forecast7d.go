package yytianqi

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type Forecast7dResponse struct {
	Code       int    `json:"code"`
	Directions string `json:"directions,omitempty"`
	Msg        string `json:"msg"`
	Counts     int    `json:"counts,omitempty"` // 访问的剩余次数
	Data       struct {
		CityId   string `json:"cityId"`   // 城市id
		CityName string `json:"cityName"` // 城市名称
		Sj       string `json:"sj"`       // 数据更新时间
		List     []struct {
			Tq1    string      `json:"tq1"`    // 白天天气
			Tq2    string      `json:"tq2"`    // 夜间天气，当与白天天气相同时，两者可合并为一个
			Numtq1 string      `json:"numtq1"` // 白天天气编码
			Numtq2 string      `json:"numtq2"` // 夜间天气编码
			Qw1    interface{} `json:"qw1"`    // 白天气温
			Qw2    interface{} `json:"qw2"`    // 夜间气温
			Fl1    string      `json:"fl1"`    // 白天风力
			Fl2    string      `json:"fl2"`    // 夜间风力
			Numfl1 string      `json:"numfl1"` // 白天风力编码
			Numfl2 string      `json:"numfl2"` // 夜间风力编码
			Fx1    string      `json:"fx1"`    // 白天风向
			Fx2    string      `json:"fx2"`    // 夜间风向，如白天风力风向与夜间风力风向一致，可合并为一行
			Numfx1 string      `json:"numfx1"` // 白天风向编码
			Numfx2 string      `json:"numfx2"` // 夜间风向编码
			Date   string      `json:"date"`   // 预报日期
		} `json:"list"`
	} `json:"data,omitempty"`
}

type Forecast7dResult struct {
	Result Forecast7dResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newForecast7dResult(result Forecast7dResponse, body []byte, http gorequest.Response) *Forecast7dResult {
	return &Forecast7dResult{Result: result, Body: body, Http: http}
}

// Forecast7d 7天预报
// http://www.yytianqi.com/api.html
func (c *Client) Forecast7d(ctx context.Context, city string, notMustParams ...gorequest.Params) (*Forecast7dResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("city", city)
	params.Set("key", c.key)
	// 请求
	request, err := c.request(ctx, "forecast7d", params, http.MethodGet)
	if err != nil {
		return newForecast7dResult(Forecast7dResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response Forecast7dResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newForecast7dResult(response, request.ResponseBody, request), err
}
