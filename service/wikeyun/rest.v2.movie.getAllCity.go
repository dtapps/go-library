package wikeyun

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type RestV2MovieGetAllCityResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		HotCity struct {
			CityId     int64  `json:"cityId"`     // 城市ID
			CityName   string `json:"cityName"`   // 城市名称
			CityPinyin string `json:"cityPinyin"` // 首字母
		} `json:"hotCity"`
		CityList struct {
			CityId     int64  `json:"cityId"`     // 城市ID
			CityName   string `json:"cityName"`   // 城市名称
			CityPinyin string `json:"cityPinyin"` // 首字母
		} `json:"cityList"`
	} `json:"data"` // 城市列表
}

type RestV2MovieGetAllCityResult struct {
	Result RestV2MovieGetAllCityResponse // 结果
	Body   []byte                        // 内容
	Http   gorequest.Response            // 请求
}

func newRestV2MovieGetAllCityResult(result RestV2MovieGetAllCityResponse, body []byte, http gorequest.Response) *RestV2MovieGetAllCityResult {
	return &RestV2MovieGetAllCityResult{Result: result, Body: body, Http: http}
}

// RestV2MovieGetAllCity 定位--获取全国所有城市（支持字母汉字搜索）
// keyword = 关键词搜索
// https://open.wikeyun.cn/#/apiDocument/4/document/510
func (c *Client) RestV2MovieGetAllCity(ctx context.Context, notMustParams ...gorequest.Params) (*RestV2MovieGetAllCityResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "rest/v2/movie/getAllCity")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.GetStoreId()) // 店铺ID

	// 请求
	var response RestV2MovieGetAllCityResponse
	request, err := c.request(ctx, "rest/v2/movie/getAllCity", params, &response)
	return newRestV2MovieGetAllCityResult(response, request.ResponseBody, request), err
}
