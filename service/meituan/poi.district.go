package meituan

import (
	"encoding/json"
	"net/http"
)

type PoiDistrictResponse struct {
	Code int `json:"code"` // 状态码 0表示请求正常
	Data []struct {
		Name string `json:"name"` // 行政区名称
		ID   int    `json:"id"`   // 行政区id
	} `json:"data"` // 返回行政区列表
}

type PoiDistrictResult struct {
	Result PoiDistrictResponse // 结果
	Body   []byte              // 内容
	Err    error               // 错误
}

func NewPoiDistrictResult(result PoiDistrictResponse, body []byte, err error) *PoiDistrictResult {
	return &PoiDistrictResult{Result: result, Body: body, Err: err}
}

// PoiDistrict 基础数据 - 城市的行政区接口
// https://openapi.meituan.com/#api-0.%E5%9F%BA%E7%A1%80%E6%95%B0%E6%8D%AE-GetHttpsOpenapiMeituanComPoiDistrictCityid1
func (app *App) PoiDistrict(cityID int) *PoiDistrictResult {
	// 参数
	param := NewParams()
	param.Set("cityid", cityID)
	params := app.NewParamsWith(param)
	// 请求
	body, err := app.request("https://openapi.meituan.com/poi/district", params, http.MethodGet)
	// 定义
	var response PoiDistrictResponse
	err = json.Unmarshal(body, &response)
	return NewPoiDistrictResult(response, body, err)
}
