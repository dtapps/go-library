package pintoto

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
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
}

func newGetCityAreaResult(result GetCityAreaResponse, body []byte, http gorequest.Response) *GetCityAreaResult {
	return &GetCityAreaResult{Result: result, Body: body, Http: http}
}

// GetCityArea 城市下区域
// https://www.showdoc.com.cn/1154868044931571/6243539682553126
func (c *Client) GetCityArea(ctx context.Context, cityId int, notMustParams ...*gorequest.Params) (*GetCityAreaResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("cityId", cityId)
	// 请求
	request, err := c.request(ctx, apiUrl+"/movieapi/movie-info/get-city-area", params)
	if err != nil {
		return newGetCityAreaResult(GetCityAreaResponse{}, request.ResponseBody, request), err
	}
	var response GetCityAreaResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGetCityAreaResult(response, request.ResponseBody, request), err
}
