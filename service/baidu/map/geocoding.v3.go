package _map

import (
	"go.dtapp.net/library/utils/gorequest"
)

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"net/http"
)

type GeocodingV3Response struct {
	Status int64 `json:"status"`
	Result struct {
		Location struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
		Precise       int64  `json:"precise"`
		Confidence    int64  `json:"confidence"`
		Comprehension int64  `json:"comprehension"`
		Level         string `json:"level"`
	} `json:"result"`
}

type GeocodingV3Result struct {
	Result GeocodingV3Response // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newGeocodingV3Result(result GeocodingV3Response, body []byte, http gorequest.Response) *GeocodingV3Result {
	return &GeocodingV3Result{Result: result, Body: body, Http: http}
}

// GeocodingV3 地理编码服务
// https://lbsyun.baidu.com/faq/api?title=webapi/guide/webservice-geocoding-base
func (c *Client) GeocodingV3(ctx context.Context, address string, notMustParams ...gorequest.Params) (*GeocodingV3Result, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("ak", c.ak)
	params.Set("address", address)
	params.Set("output", "json")
	// 请求
	request, err := c.request(ctx, "geocoding/v3/", params, http.MethodGet)
	if err != nil {
		return newGeocodingV3Result(GeocodingV3Response{}, request.ResponseBody, request), err
	}
	// 定义
	var response GeocodingV3Response
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGeocodingV3Result(response, request.ResponseBody, request), err
}
