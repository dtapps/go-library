package pintoto

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
)

type GetCityAreaResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			AreaId   int    `json:"areaId"`   // 区域id
			AreaName string `json:"areaName"` // 区域名
		} `json:"list"`
	} `json:"data"`
	Success bool `json:"success"`
}

type GetCityAreaResult struct {
	Result GetCityAreaResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
	Err    error               // 错误
}

func newGetCityAreaResult(result GetCityAreaResponse, body []byte, http gorequest.Response, err error) *GetCityAreaResult {
	return &GetCityAreaResult{Result: result, Body: body, Http: http, Err: err}
}

// GetCityArea 城市下区域
// https://www.showdoc.com.cn/1154868044931571/6243539682553126
func (c *Client) GetCityArea(ctx context.Context, cityId int) *GetCityAreaResult {
	// 测试
	param := gorequest.NewParams()
	param.Set("cityId", cityId)
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(ctx, apiUrl+"/movieapi/movie-info/get-city-area", params)
	var response GetCityAreaResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newGetCityAreaResult(response, request.ResponseBody, request, err)
}
