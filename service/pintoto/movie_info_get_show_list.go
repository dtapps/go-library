package pintoto

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type GetShowList struct {
	Page      int     `json:"page,omitempty"`      // 页码，默认1
	Limit     int     `json:"limit,omitempty"`     // 条数，默认 10
	FilmId    int     `json:"filmId"`              // 影片id，由热映/即将上映接口获得
	CityId    int     `json:"cityId"`              // 城市id，由城市列表接口获得
	Area      string  `json:"area,omitempty"`      // 区域名，由区域列表接口获得
	Date      string  `json:"date,omitempty"`      // 日期，例：2020-01-01，不传默认当天
	Latitude  float64 `json:"latitude,omitempty"`  // 纬度，不传则无距离排序
	Longitude float64 `json:"longitude,omitempty"` // 经度，不传则无距离排序
}

type GetShowListResponse struct {
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

type GetShowListResult struct {
	Result GetShowListResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
	Err    error               // 错误
}

func newGetShowListResult(result GetShowListResponse, body []byte, http gorequest.Response, err error) *GetShowListResult {
	return &GetShowListResult{Result: result, Body: body, Http: http, Err: err}
}

// GetShowList 包含某电影的影院 https://www.showdoc.com.cn/1154868044931571/6067372188376779
func (c *Client) GetShowList(ctx context.Context, param GetShowList) *GetShowListResult {
	// api params
	params := map[string]interface{}{}
	b, _ := gojson.Marshal(&param)
	var m map[string]interface{}
	_ = gojson.Unmarshal(b, &m)
	for k, v := range m {
		params[k] = v
	}
	request, err := c.request(ctx, apiUrl+"/movieapi/movie-info/get-show-list", params)
	// 定义
	var response GetShowListResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGetShowListResult(response, request.ResponseBody, request, err)
}
