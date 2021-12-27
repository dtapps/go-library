package pintoto

import (
	"encoding/json"
)

type GetCityArea struct {
	CityId int `json:"cityId"` // 城市id
}

type GetCityAreaResult struct {
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

// GetCityArea 城市下区域 https://www.showdoc.com.cn/1154868044931571/6243539682553126
func (app *App) GetCityArea(param GetCityArea) (result GetCityAreaResult, err error) {
	// api params
	params := map[string]interface{}{}
	b, _ := json.Marshal(&param)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		params[k] = v
	}
	body, err := app.request("https://movieapi2.pintoto.cn/movieapi/movie-info/get-city-area", params)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
