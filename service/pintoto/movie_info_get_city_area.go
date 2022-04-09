package pintoto

import (
	"encoding/json"
)

type GetCityAreaResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			AreaId   int    `json:"areaId"`   // 区域id
			AreaName string `json:"areaName"` // 区域名
		} `json:"list"`
	} `json:"data"`
	Success bool `json:"success"`
}

type GetCityAreaResult struct {
	Result GetCityAreaResponse // 结果
	Body   []byte              // 内容
	Err    error               // 错误
}

func NewGetCityAreaResult(result GetCityAreaResponse, body []byte, err error) *GetCityAreaResult {
	return &GetCityAreaResult{Result: result, Body: body, Err: err}
}

// GetCityArea 城市下区域
// https://www.showdoc.com.cn/1154868044931571/6243539682553126
func (app *App) GetCityArea(cityId int) *GetCityAreaResult {
	// 测试
	param := NewParams()
	param.Set("cityId", cityId)
	params := app.NewParamsWith(param)
	// 请求
	body, err := app.request("https://movieapi2.pintoto.cn/movieapi/movie-info/get-city-area", params)
	var response GetCityAreaResponse
	err = json.Unmarshal(body, &response)
	return NewGetCityAreaResult(response, body, err)
}
