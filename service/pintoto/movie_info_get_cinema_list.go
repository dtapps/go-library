package pintoto

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
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
	Err    error                 // 错误
}

func newGetCinemaListResult(result GetCinemaListResponse, body []byte, http gorequest.Response, err error) *GetCinemaListResult {
	return &GetCinemaListResult{Result: result, Body: body, Http: http, Err: err}
}

// GetCinemaList 影院列表 https://www.showdoc.com.cn/1154868044931571/5866426126744792
func (c *Client) GetCinemaList(cityId int) *GetCinemaListResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("cityId", cityId)
	// 转换
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(apiUrl+"/movieapi/movie-info/get-cinema-list", params)
	// 定义
	var response GetCinemaListResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newGetCinemaListResult(response, request.ResponseBody, request, err)
}
