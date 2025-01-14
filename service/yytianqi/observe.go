package yytianqi

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type ObserveResponse struct {
	Code       int    `json:"code"`
	Directions string `json:"directions,omitempty"`
	Msg        string `json:"msg"`
	Counts     int    `json:"counts,omitempty"` // 访问的剩余次数
	Data       struct {
		CityId     string      `json:"cityId"`     // 城市id
		CityName   string      `json:"cityName"`   // 城市名称
		LastUpdate string      `json:"lastUpdate"` // 实况更新时间
		Tq         string      `json:"tq"`         // 天气现象
		Numtq      interface{} `json:"numtq"`      // 天气现象编码
		Qw         string      `json:"qw"`         // 当前气温
		Fl         string      `json:"fl"`         // 当前风力
		Numfl      interface{} `json:"numfl"`      // 当前风力编码
		Fx         string      `json:"fx"`         // 当前风向
		Numfx      interface{} `json:"numfx"`      // 当前风向编码
		Sd         string      `json:"sd"`         // 相对湿度，直接在此数值后添加%即可
	} `json:"data,omitempty"`
}

type ObserveResult struct {
	Result ObserveResponse    // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newObserveResult(result ObserveResponse, body []byte, http gorequest.Response) *ObserveResult {
	return &ObserveResult{Result: result, Body: body, Http: http}
}

// Observe 天气实况
// city 城市ID / IP地址或"ip"两个字母 / 经纬度
// http://www.yytianqi.com/api.html
func (c *Client) Observe(ctx context.Context, city string, notMustParams ...gorequest.Params) (*ObserveResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("city", city)
	params.Set("key", c.key)
	// 请求
	request, err := c.request(ctx, "observe", params, http.MethodGet)
	if err != nil {
		return newObserveResult(ObserveResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ObserveResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newObserveResult(response, request.ResponseBody, request), err
}
