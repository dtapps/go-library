package meituan

import (
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type PoiAreaResponse struct {
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

type PoiAreaResult struct {
	Result PoiAreaResponse    // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
	Err    error              // 错误
}

func newPoiAreaResult(result PoiAreaResponse, body []byte, http gorequest.Response, err error) *PoiAreaResult {
	return &PoiAreaResult{Result: result, Body: body, Http: http, Err: err}
}

// PoiArea 基础数据 - 商圈接口
// https://openapi.meituan.com/#api-0.%E5%9F%BA%E7%A1%80%E6%95%B0%E6%8D%AE-GetHttpsOpenapiMeituanComPoiAreaCityid1
func (c *Client) PoiArea(cityID int) *PoiAreaResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("cityid", cityID)
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(apiUrl+"/poi/area", params, http.MethodGet)
	// 定义
	var response PoiAreaResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newPoiAreaResult(response, request.ResponseBody, request, err)
}
