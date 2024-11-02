package wikeyun

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type RestMovieAllCityResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		CityId      int64  `json:"cityId"`      // 城市ID
		CityName    string `json:"cityName"`    // 城市名称
		Firstletter string `json:"firstletter"` // 城市名称首字母
		Ishot       int64  `json:"ishot"`       // 是否热门 1是 0否
	} `json:"data"` // 城市列表
}

type RestMovieAllCityResult struct {
	Result RestMovieAllCityResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
}

func newRestMovieAllCityResult(result RestMovieAllCityResponse, body []byte, http gorequest.Response) *RestMovieAllCityResult {
	return &RestMovieAllCityResult{Result: result, Body: body, Http: http}
}

// RestMovieAllCity 定位--获取全国所有城市（支持字母汉字搜索）
// keyword = 关键词搜索
// https://open.wikeyun.cn/#/apiDocument/4/document/302
func (c *Client) RestMovieAllCity(ctx context.Context, notMustParams ...gorequest.Params) (*RestMovieAllCityResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.GetStoreId()) // 店铺ID

	// 请求
	var response RestMovieAllCityResponse
	request, err := c.request(ctx, "rest/movie/allCity", params, &response)
	return newRestMovieAllCityResult(response, request.ResponseBody, request), err
}
