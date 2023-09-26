package baidu

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type ReverseGeocodingResponse struct {
	Status int `json:"status"`
	Result struct {
		Location struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
		FormattedAddress string `json:"formatted_address"`
		Business         string `json:"business"`
		AddressComponent struct {
			Country         string `json:"country"`
			CountryCode     int    `json:"country_code"`
			CountryCodeIso  string `json:"country_code_iso"`
			CountryCodeIso2 string `json:"country_code_iso2"`
			Province        string `json:"province"`
			City            string `json:"city"`
			CityLevel       int    `json:"city_level"`
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
		CityCode           int           `json:"cityCode"`
	} `json:"result"`
}

type ReverseGeocodingResult struct {
	Result ReverseGeocodingResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
}

func newReverseGeocodingResult(result ReverseGeocodingResponse, body []byte, http gorequest.Response) *ReverseGeocodingResult {
	return &ReverseGeocodingResult{Result: result, Body: body, Http: http}
}

// ReverseGeocoding 全球逆地理编码服务
// https://lbsyun.baidu.com/index.php?title=webapi/guide/webservice-geocoding-abroad
func (c *Client) ReverseGeocoding(ctx context.Context, location string, notMustParams ...gorequest.Params) (*ReverseGeocodingResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("ak", c.GetAk())
	params.Set("location", location)
	params.Set("output", "json")
	// 请求
	request, err := c.request(ctx, apiUrl+"/reverse_geocoding/v3/", params, http.MethodGet)
	if err != nil {
		return newReverseGeocodingResult(ReverseGeocodingResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ReverseGeocodingResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newReverseGeocodingResult(response, request.ResponseBody, request), err
}
