package _map

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type ReverseGeocodingV3Response struct {
	Status int64 `json:"status"`
	Result struct {
		Location struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
		FormattedAddress string `json:"formatted_address"`
		Business         string `json:"business"`
		AddressComponent struct {
			Country         string `json:"country"`
			CountryCode     int64  `json:"country_code"`
			CountryCodeIso  string `json:"country_code_iso"`
			CountryCodeIso2 string `json:"country_code_iso2"`
			Province        string `json:"province"`
			City            string `json:"city"`
			CityLevel       int64  `json:"city_level"`
			District        string `json:"district"`
			Town            string `json:"town"`
			TownCode        string `json:"town_code"`
			Distance        string `json:"distance"`
			Direction       string `json:"direction"`
			Adcode          string `json:"adcode"`
			Street          string `json:"street"`
			StreetNumber    string `json:"street_number"`
		} `json:"addressComponent"`
		Pois               []interface{} `json:"pois"`
		Roads              []interface{} `json:"roads"`
		PoiRegions         []interface{} `json:"poiRegions"`
		SematicDescription string        `json:"sematic_description"`
		CityCode           int64         `json:"cityCode"`
	} `json:"result"`
}

type ReverseGeocodingV3Result struct {
	Result ReverseGeocodingV3Response // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
}

func newReverseGeocodingV3Result(result ReverseGeocodingV3Response, body []byte, http gorequest.Response) *ReverseGeocodingV3Result {
	return &ReverseGeocodingV3Result{Result: result, Body: body, Http: http}
}

// ReverseGeocodingV3 全球逆地理编码服务
// https://lbsyun.baidu.com/faq/api?title=webapi/guide/webservice-geocoding-abroad-base
func (c *Client) ReverseGeocodingV3(ctx context.Context, location string, notMustParams ...gorequest.Params) (*ReverseGeocodingV3Result, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("ak", c.ak)
	params.Set("location", location)
	params.Set("output", "json")
	// 请求
	request, err := c.request(ctx, "reverse_geocoding/v3/", params, http.MethodGet)
	if err != nil {
		return newReverseGeocodingV3Result(ReverseGeocodingV3Response{}, request.ResponseBody, request), err
	}
	// 定义
	var response ReverseGeocodingV3Response
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newReverseGeocodingV3Result(response, request.ResponseBody, request), err
}
