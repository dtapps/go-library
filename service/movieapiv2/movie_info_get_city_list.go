package movieapiv2

import (
	"encoding/json"
)

type GetCityListResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			PinYin     string `json:"pinYin"`     // 城市首字母
			RegionName string `json:"regionName"` // 城市名
			CityId     int    `json:"cityId"`     // 城市id
		} `json:"list"`
	} `json:"data"`
	Success bool `json:"success"`
}

// GetCityList 城市列表 https://www.showdoc.com.cn/1154868044931571/5865562425538244
func (app *App) GetCityList() (result GetCityListResult, err error) {
	body, err := app.request("movieapi/movie-info/get-city-list", map[string]interface{}{})
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
