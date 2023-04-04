package qq

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type GeocoderAddressResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  struct {
		Title    string `json:"title"`
		Location struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
		AdInfo struct {
			Adcode string `json:"adcode"`
		} `json:"ad_info"`
		AddressComponents struct {
			Province     string `json:"province"`
			City         string `json:"city"`
			District     string `json:"district"`
			Street       string `json:"street"`
			StreetNumber string `json:"street_number"`
		} `json:"address_components"`
		Similarity  float64 `json:"similarity"`
		Deviation   int     `json:"deviation"`
		Reliability int     `json:"reliability"`
		Level       int     `json:"level"`
	} `json:"result"`
}

type GeocoderAddressResult struct {
	Result GeocoderAddressResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
}

func newGeocoderAddressResult(result GeocoderAddressResponse, body []byte, http gorequest.Response) *GeocoderAddressResult {
	return &GeocoderAddressResult{Result: result, Body: body, Http: http}
}

// GeocoderAddress 地址解析（地址转坐标）
// https://lbs.qq.com/service/webService/webServiceGuide/webServiceGeocoder
func (c *Client) GeocoderAddress(ctx context.Context, address string, notMustParams ...gorequest.Params) (*GeocoderAddressResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("key", c.GetKey())
	params.Set("address", address)
	params.Set("output", "JSON")
	// 请求
	request, err := c.request(ctx, apiUrl+"/ws/geocoder/v1/", params, http.MethodGet)
	if err != nil {
		return newGeocoderAddressResult(GeocoderAddressResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response GeocoderAddressResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGeocoderAddressResult(response, request.ResponseBody, request), err
}
