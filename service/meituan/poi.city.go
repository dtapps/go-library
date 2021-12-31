package meituan

import (
	"encoding/json"
	"net/http"
)

type PoiCityResponse struct {
	Code int `json:"code"` // 状态码 0表示请求正常
	Data []struct {
		Pinyin string `json:"pinyin"` // 城市拼音
		Name   string `json:"name"`   // 城市名称
		ID     int    `json:"id"`     // 城市id
	} `json:"data"` // 返回城市列表
}

type PoiCityResult struct {
	Result PoiCityResponse // 结果
	Body   []byte          // 内容
	Err    error           // 错误
}

func NewPoiCityResult(result PoiCityResponse, body []byte, err error) *PoiCityResult {
	return &PoiCityResult{Result: result, Body: body, Err: err}
}

// PoiCity 基础数据 - 开放城市接口
// https://openapi.meituan.com/#api-0.%E5%9F%BA%E7%A1%80%E6%95%B0%E6%8D%AE-GetHttpsOpenapiMeituanComPoiCity
func (app *App) PoiCity() *PoiCityResult {
	// 请求
	body, err := app.request("https://openapi.meituan.com/poi/city", map[string]interface{}{}, http.MethodGet)
	// 定义
	var response PoiCityResponse
	err = json.Unmarshal(body, &response)
	return NewPoiCityResult(response, body, err)
}
