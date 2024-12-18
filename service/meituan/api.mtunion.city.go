package meituan

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
	"net/http"
)

type ApiMtUnionCityResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		DataList []struct {
			CityId   float64 `json:"cityId"`   // 城市ID
			CityName string  `json:"cityName"` // 城市名称
		} `json:"dataList"`
		Total int64 `json:"total"` // 查询总数
	} `json:"data"`
}
type ApiMtUnionCityResult struct {
	Result ApiMtUnionCityResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newApiMtUnionCityResult(result ApiMtUnionCityResponse, body []byte, http gorequest.Response) *ApiMtUnionCityResult {
	return &ApiMtUnionCityResult{Result: result, Body: body, Http: http}
}

// ApiMtUnionCity 城市信息查询（新版）
// https://union.meituan.com/v2/apiDetail?id=29
func (c *Client) ApiMtUnionCity(ctx context.Context, notMustParams ...*gorequest.Params) (*ApiMtUnionCityResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求时刻10位时间戳(秒级)，有效期60s
	params.Set("ts", gotime.Current().Timestamp())
	params.Set("appkey", c.GetAppKey())
	params.Set("sign", c.getSign(c.GetSecret(), params))

	// 请求
	var response ApiMtUnionCityResponse
	request, err := c.request(ctx, "api/getqualityscorebysid", params, http.MethodGet, &response)
	return newApiMtUnionCityResult(response, request.ResponseBody, request), err
}
