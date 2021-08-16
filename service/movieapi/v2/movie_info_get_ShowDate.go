package v2

import (
	"encoding/json"
)

type GetShowDate struct {
	FilmId int `json:"filmId"` // 影片id，由热映/即将上映接口获得
	CityId int `json:"cityId"` // 城市id，由城市列表接口获得
}

type GetShowDateResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		DateList []string `json:"dateList"`
	} `json:"data"`
	Success bool `json:"success"`
}

// GetShowDate 包含某电影的日期 https://www.showdoc.com.cn/1154868044931571/6091788579441818
func (app *App) GetShowDate(param GetShowDate) (result GetShowDateResult, err error) {
	// api params
	params := map[string]interface{}{}
	b, _ := json.Marshal(&param)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		params[k] = v
	}
	body, err := app.request("movieapi/movie-info/get-show-date", params)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
