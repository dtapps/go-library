package meituan

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotime"
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
	Err    error                  // 错误
}

func newApiMtUnionCityResult(result ApiMtUnionCityResponse, body []byte, http gorequest.Response, err error) *ApiMtUnionCityResult {
	return &ApiMtUnionCityResult{Result: result, Body: body, Http: http, Err: err}
}

// ApiMtUnionCity 城市信息查询（新版）
// https://union.meituan.com/v2/apiDetail?id=29
func (c *Client) ApiMtUnionCity(ctx context.Context, notMustParams ...gorequest.Params) *ApiMtUnionCityResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求时刻10位时间戳(秒级)，有效期60s
	params["ts"] = gotime.Current().Timestamp()
	params["appkey"] = c.GetAppKey()
	params["sign"] = c.getSign(c.GetSecret(), params)
	// 请求
	request, err := c.request(ctx, apiUrl+"/api/getqualityscorebysid", params, http.MethodGet)
	// 定义
	var response ApiMtUnionCityResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newApiMtUnionCityResult(response, request.ResponseBody, request, err)
}
