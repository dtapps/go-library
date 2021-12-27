package pintoto

import (
	"encoding/json"
)

type GetHotList struct {
	CityId int `json:"cityId,omitempty"` // 传入cityId时，会显示当前城市下的相关电影。 如果不传，则默认显示北京的电影
}

type GetHotListResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		HasMore int `json:"hasMore"`
		List    []struct {
			Director     string `json:"director"`     // 导演
			PublishDate  string `json:"publishDate"`  // 影片上映日期
			VersionTypes string `json:"versionTypes"` // 上映类型
			Language     string `json:"language"`     // 语言
			ShowStatus   int    `json:"showStatus"`   // 放映状态：1 正在热映。2 即将上映
			Pic          string `json:"pic"`          // 海报URL地址
			FilmTypes    string `json:"filmTypes"`    // 影片类型
			LikeNum      int    `json:"likeNum"`      // 想看人数
			Duration     int    `json:"duration"`     // 时长，分钟
			Cast         string `json:"cast"`         // 主演
			FilmId       int    `json:"filmId"`       // 影片id
			Grade        string `json:"grade"`        // 评分
			Intro        string `json:"intro"`        // 简介
			Name         string `json:"name"`         // 影片名
		} `json:"list"`
	} `json:"data"`
	Success bool `json:"success"`
}

// GetHotList 正在热映 https://www.showdoc.com.cn/1154868044931571/5866125707634369
func (app *App) GetHotList(param GetHotList) (result GetHotListResult, err error) {
	// api params
	params := map[string]interface{}{}
	b, _ := json.Marshal(&param)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		params[k] = v
	}
	body, err := app.request("https://movieapi2.pintoto.cn/movieapi/movie-info/get-hot-list", params)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
