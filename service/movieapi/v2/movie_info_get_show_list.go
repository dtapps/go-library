package v2

import (
	"encoding/json"
)

type GetShowList struct {
	Page      int     `json:"page"`      // 页码，默认1
	Limit     int     `json:"limit"`     // 条数，默认 10
	FilmId    int     `json:"filmId"`    // 影片id，由热映/即将上映接口获得
	CityId    int     `json:"cityId"`    // 城市id，由城市列表接口获得
	Area      string  `json:"area"`      // 区域名，由区域列表接口获得
	Date      string  `json:"date"`      // 日期，例：2020-01-01，不传默认当天
	Latitude  float64 `json:"latitude"`  // 纬度，不传则无距离排序
	Longitude float64 `json:"longitude"` // 经度，不传则无距离排序
}

type GetShowListResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		HasMore int `json:"hasMore"`
		List    []struct {
			Address    string  `json:"address"`
			ShowId     string  `json:"showId"`
			Distance   string  `json:"distance"`
			CinemaId   int     `json:"cinemaId"`
			CinemaName string  `json:"cinemaName"`
			Latitude   float64 `json:"latitude"`
			ShowTime   string  `json:"showTime"`
			HallName   string  `json:"hallName"`
			Longitude  float64 `json:"longitude"`
		} `json:"list"`
	} `json:"data"`
	Success bool `json:"success"`
}

// GetShowList 包含某电影的影院 https://www.showdoc.com.cn/1154868044931571/6067372188376779
func (app *App) GetShowList(param GetShowList) (result GetShowListResult, err error) {
	// api params
	params := map[string]interface{}{}
	b, _ := json.Marshal(&param)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		params[k] = v
	}
	body, err := app.request("movieapi/movie-info/get-show-list", params)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
