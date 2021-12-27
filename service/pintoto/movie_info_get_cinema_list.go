package pintoto

import (
	"encoding/json"
)

type GetCinemaList struct {
	CityId int `json:"cityId"` // 城市id
}

type GetCinemaListResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			CinemaId          int     `json:"cinemaId"`          // 影院id
			CityId            int     `json:"cityId"`            // 城市id
			CinemaName        string  `json:"cinemaName"`        // 影院名称
			Address           string  `json:"address"`           // 影院地址
			Latitude          float64 `json:"latitude"`          // 纬度
			Longitude         float64 `json:"longitude"`         // 经度
			Phone             string  `json:"phone"`             // 影院电话
			RegionName        string  `json:"regionName"`        // 地区名称
			IsAcceptSoonOrder int     `json:"isAcceptSoonOrder"` // 是否支持秒出票，0为不支持，1为支持
			NetPrice          int     `json:"netPrice"`          // 当前影院最低价的排期
		} `json:"list"`
	} `json:"data"`
	Success bool `json:"success"`
}

// GetCinemaList 影院列表 https://www.showdoc.com.cn/1154868044931571/5866426126744792
func (app *App) GetCinemaList(param GetCinemaList) (result GetCinemaListResult, err error) {
	// api params
	params := map[string]interface{}{}
	b, _ := json.Marshal(&param)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		params[k] = v
	}
	body, err := app.request("https://movieapi2.pintoto.cn/movieapi/movie-info/get-cinema-list", params)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
