package meituan

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
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
	Http   gorequest.Response  // 请求
}

func newPoiDistrictResult(result PoiDistrictResponse, body []byte, http gorequest.Response) *PoiDistrictResult {
	return &PoiDistrictResult{Result: result, Body: body, Http: http}
}

// PoiDistrict 基础数据 - 城市的行政区接口
// https://openapi.meituan.com/#api-0.%E5%9F%BA%E7%A1%80%E6%95%B0%E6%8D%AE-GetHttpsOpenapiMeituanComPoiDistrictCityid1
func (c *Client) PoiDistrict(ctx context.Context, cityID int, notMustParams ...gorequest.Params) (*PoiDistrictResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "poi/district")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("cityid", cityID)

	// 请求
	var response PoiDistrictResponse
	request, err := c.request(ctx, "poi/district", params, http.MethodGet, &response)
	return newPoiDistrictResult(response, request.ResponseBody, request), err
}
