package meituan

import (
	"encoding/json"
	"net/http"
)

type OpenapiPoiCategoryResult struct {
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

// OpenapiPoiCategory 基础数据 - 品类接口 https://openapi.meituan.com/#api-0.%E5%9F%BA%E7%A1%80%E6%95%B0%E6%8D%AE-GetHttpsOpenapiMeituanComPoiDistrictCityid1
func (app *App) OpenapiPoiCategory(cityID int) (result OpenapiPoiCategoryResult, err error) {

	param := NewParams()
	param.Set("cityid", cityID)

	// 参数
	params := app.NewParamsWith(param)

	// 请求
	body, err := app.request("https://openapi.meituan.com/poi/category", params, http.MethodGet)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}

type OpenapiPoiAreaResult struct {
	Code int `json:"code"`
	Data []struct {
		Area []struct {
			Name string `json:"name"` // 商圈名称
			ID   int    `json:"id"`   // 商圈id
		} `json:"area"`
		Name string `json:"name"` // 行政区名称
		ID   int    `json:"id"`   // 行政区id
	} `json:"data"`
}

// OpenapiPoiArea 基础数据 - 商圈接口 https://openapi.meituan.com/#api-0.%E5%9F%BA%E7%A1%80%E6%95%B0%E6%8D%AE-GetHttpsOpenapiMeituanComPoiAreaCityid1
func (app *App) OpenapiPoiArea(cityID int) (result OpenapiPoiAreaResult, err error) {

	param := NewParams()
	param.Set("cityid", cityID)

	// 参数
	params := app.NewParamsWith(param)

	// 请求
	body, err := app.request("https://openapi.meituan.com/poi/area", params, http.MethodGet)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}

type OpenapiPoiDistrictResult struct {
	Code int `json:"code"` // 状态码 0表示请求正常
	Data []struct {
		Name string `json:"name"` // 行政区名称
		ID   int    `json:"id"`   // 行政区id
	} `json:"data"` // 返回行政区列表
}

// OpenapiPoiDistrict 基础数据 - 城市的行政区接口 https://openapi.meituan.com/#api-0.%E5%9F%BA%E7%A1%80%E6%95%B0%E6%8D%AE-GetHttpsOpenapiMeituanComPoiDistrictCityid1
func (app *App) OpenapiPoiDistrict(cityID int) (result OpenapiPoiDistrictResult, err error) {

	param := NewParams()
	param.Set("cityid", cityID)

	// 参数
	params := app.NewParamsWith(param)

	// 请求
	body, err := app.request("https://openapi.meituan.com/poi/district", params, http.MethodGet)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}

type OpenapiPoiCityResult struct {
	Code int `json:"code"` // 状态码 0表示请求正常
	Data []struct {
		Pinyin string `json:"pinyin"` // 城市拼音
		Name   string `json:"name"`   // 城市名称
		ID     int    `json:"id"`     // 城市id
	} `json:"data"` // 返回城市列表
}

// OpenapiPoiCity 基础数据 - 开放城市接口 https://openapi.meituan.com/#api-0.%E5%9F%BA%E7%A1%80%E6%95%B0%E6%8D%AE-GetHttpsOpenapiMeituanComPoiCity
func (app *App) OpenapiPoiCity() (result OpenapiPoiCityResult, err error) {

	// 请求
	body, err := app.request("https://openapi.meituan.com/poi/city", map[string]interface{}{}, http.MethodGet)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
