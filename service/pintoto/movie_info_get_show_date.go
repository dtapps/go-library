package pintoto

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
func (app *App) GetShowDate(cityId, filmId int) (result GetShowDateResult, err error) {
	// 参数
	param := NewParams()
	param.Set("cityId", cityId)
	param.Set("filmId", filmId)
	// 转换
	params := app.NewParamsWith(param)
	// 请求
	body, err := app.request("https://movieapi2.pintoto.cn/movieapi/movie-info/get-show-date", params)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
