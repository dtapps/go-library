package meituan

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PoiCategoryResponse struct {
	Code int `json:"code"`
	Data []struct {
		Name    string `json:"name"`
		Subcate []struct {
			Name string `json:"name"` // 品类名称
			ID   int    `json:"id"`   // 品类id
		} `json:"subcate"`
		ID int `json:"id"`
	} `json:"data"`
}

type PoiCategoryResult struct {
	Result PoiCategoryResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newPoiCategoryResult(result PoiCategoryResponse, body []byte, http gorequest.Response) *PoiCategoryResult {
	return &PoiCategoryResult{Result: result, Body: body, Http: http}
}

// PoiCategory 基础数据 - 品类接口
// https://openapi.meituan.com/#api-0.%E5%9F%BA%E7%A1%80%E6%95%B0%E6%8D%AE-GetHttpsOpenapiMeituanComPoiDistrictCityid1
func (c *Client) PoiCategory(ctx context.Context, cityID int, notMustParams ...*gorequest.Params) (*PoiCategoryResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("cityid", cityID)

	// 请求
	var response PoiCategoryResponse
	request, err := c.request(ctx, "poi/category", params, http.MethodGet, &response)
	return newPoiCategoryResult(response, request.ResponseBody, request), err
}
