package pintoto

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

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
}

func newGetShowListResult(result GetShowListResponse, body []byte, http gorequest.Response) *GetShowListResult {
	return &GetShowListResult{Result: result, Body: body, Http: http}
}

// GetShowList 包含某电影的影院 https://www.showdoc.com.cn/1154868044931571/6067372188376779
func (c *Client) GetShowList(ctx context.Context, notMustParams ...*gorequest.Params) (*GetShowListResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/movieapi/movie-info/get-show-list", params)
	if err != nil {
		return newGetShowListResult(GetShowListResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response GetShowListResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGetShowListResult(response, request.ResponseBody, request), err
}
