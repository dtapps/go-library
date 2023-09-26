package amap

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type GeocodeGeoResponse struct {
	Status   string `json:"status"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
	Count    string `json:"count"`
	Geocodes []struct {
		FormattedAddress string        `json:"formatted_address"`
		Country          string        `json:"country"`
		Province         string        `json:"province"`
		Citycode         string        `json:"citycode"`
		City             string        `json:"city"`
		District         string        `json:"district"`
		Township         []interface{} `json:"township"`
		Neighborhood     struct {
			Name []interface{} `json:"name"`
			Type []interface{} `json:"type"`
		} `json:"neighborhood"`
		Building struct {
			Name []interface{} `json:"name"`
			Type []interface{} `json:"type"`
		} `json:"building"`
		Adcode   string `json:"adcode"`
		Street   string `json:"street"`
		Number   string `json:"number"`
		Location string `json:"location"`
		Level    string `json:"level"`
	} `json:"geocodes"`
}

type GeocodeGeoResult struct {
	Result GeocodeGeoResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newGeocodeGeoResult(result GeocodeGeoResponse, body []byte, http gorequest.Response) *GeocodeGeoResult {
	return &GeocodeGeoResult{Result: result, Body: body, Http: http}
}

// GeocodeGeo 地理编码
// https://lbs.amap.com/api/webservice/guide/api/georegeo
func (c *Client) GeocodeGeo(ctx context.Context, address string, notMustParams ...gorequest.Params) (*GeocodeGeoResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("key", c.GetKey())
	params.Set("address", address)
	params.Set("output", "JSON")
	// 请求
	request, err := c.request(ctx, apiUrl+"/geocode/geo", params, http.MethodGet)
	if err != nil {
		return newGeocodeGeoResult(GeocodeGeoResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response GeocodeGeoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGeocodeGeoResult(response, request.ResponseBody, request), err
}
