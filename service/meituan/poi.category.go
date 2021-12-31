package meituan

import (
	"encoding/json"
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
	Err    error               // 错误
}

func NewPoiCategoryResult(result PoiCategoryResponse, body []byte, err error) *PoiCategoryResult {
	return &PoiCategoryResult{Result: result, Body: body, Err: err}
}

// PoiCategory 基础数据 - 品类接口
// https://openapi.meituan.com/#api-0.%E5%9F%BA%E7%A1%80%E6%95%B0%E6%8D%AE-GetHttpsOpenapiMeituanComPoiDistrictCityid1
func (app *App) PoiCategory(cityID int) *PoiCategoryResult {
	// 参数
	param := NewParams()
	param.Set("cityid", cityID)
	params := app.NewParamsWith(param)
	// 请求
	body, err := app.request("https://openapi.meituan.com/poi/category", params, http.MethodGet)
	// 定义
	var response PoiCategoryResponse
	err = json.Unmarshal(body, &response)
	return NewPoiCategoryResult(response, body, err)
}
