package pintoto

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type GetCinemaListResponse struct {
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

type GetCinemaListResult struct {
	Result GetCinemaListResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newGetCinemaListResult(result GetCinemaListResponse, body []byte, http gorequest.Response) *GetCinemaListResult {
	return &GetCinemaListResult{Result: result, Body: body, Http: http}
}

// GetCinemaList 影院列表 https://www.showdoc.com.cn/1154868044931571/5866426126744792
func (c *Client) GetCinemaList(ctx context.Context, cityId int, notMustParams ...*gorequest.Params) (*GetCinemaListResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("cityId", cityId)
	// 请求
	request, err := c.request(ctx, apiUrl+"/movieapi/movie-info/get-cinema-list", params)
	if err != nil {
		return newGetCinemaListResult(GetCinemaListResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response GetCinemaListResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGetCinemaListResult(response, request.ResponseBody, request), err
}
