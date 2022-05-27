package pintoto

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
)

type GetCityListResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		List []GetCityListResponseDataList `json:"list"`
	} `json:"data"`
	Success bool `json:"success"`
}

type GetCityListResponseDataList struct {
	PinYin     string `json:"pinYin"`     // 城市首字母
	RegionName string `json:"regionName"` // 城市名
	CityId     int    `json:"cityId"`     // 城市id
}

type GetCityListResult struct {
	Result GetCityListResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
	Err    error               // 错误
}

func NewGetCityListResult(result GetCityListResponse, body []byte, http gorequest.Response, err error) *GetCityListResult {
	return &GetCityListResult{Result: result, Body: body, Http: http, Err: err}
}

// GetCityList 城市列表
// https://www.showdoc.com.cn/1154868044931571/5865562425538244
func (app *App) GetCityList() *GetCityListResult {
	request, err := app.request("https://movieapi2.pintoto.cn/movieapi/movie-info/get-city-list", map[string]interface{}{})
	var response GetCityListResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewGetCityListResult(response, request.ResponseBody, request, err)
}
