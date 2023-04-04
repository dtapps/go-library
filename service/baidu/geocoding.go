package baidu

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type GeocodingResponse struct {
	Status int `json:"status"`
	Result struct {
		Location struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
		Precise       int    `json:"precise"`
		Confidence    int    `json:"confidence"`
		Comprehension int    `json:"comprehension"`
		Level         string `json:"level"`
	} `json:"result"`
}

type GeocodingResult struct {
	Result GeocodingResponse  // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newGeocodingResult(result GeocodingResponse, body []byte, http gorequest.Response) *GeocodingResult {
	return &GeocodingResult{Result: result, Body: body, Http: http}
}

// Geocoding 地理编码服务
// https://lbsyun.baidu.com/index.php?title=webapi/guide/webservice-geocoding
func (c *Client) Geocoding(ctx context.Context, address string, notMustParams ...gorequest.Params) (*GeocodingResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("ak", c.GetAk())
	params.Set("address", address)
	params.Set("output", "json")
	// 请求
	request, err := c.request(ctx, apiUrl+"/geocoding/v3/", params, http.MethodGet)
	if err != nil {
		return newGeocodingResult(GeocodingResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response GeocodingResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGeocodingResult(response, request.ResponseBody, request), err
}
