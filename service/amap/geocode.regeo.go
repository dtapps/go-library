package amap

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type GeocodeRegeoResponse struct {
	Status    string `json:"status"`
	Regeocode struct {
		Roads []struct {
			Id        string `json:"id"`
			Location  string `json:"location"`
			Direction string `json:"direction"`
			Name      string `json:"name"`
			Distance  string `json:"distance"`
		} `json:"roads"`
		Roadinters []struct {
			SecondName string `json:"second_name"`
			FirstId    string `json:"first_id"`
			SecondId   string `json:"second_id"`
			Location   string `json:"location"`
			Distance   string `json:"distance"`
			FirstName  string `json:"first_name"`
			Direction  string `json:"direction"`
		} `json:"roadinters"`
		FormattedAddress string `json:"formatted_address"`
		AddressComponent struct {
			City         []interface{} `json:"city"`
			Province     string        `json:"province"`
			Adcode       string        `json:"adcode"`
			District     string        `json:"district"`
			Towncode     string        `json:"towncode"`
			StreetNumber struct {
				Number    string `json:"number"`
				Location  string `json:"location"`
				Direction string `json:"direction"`
				Distance  string `json:"distance"`
				Street    string `json:"street"`
			} `json:"streetNumber"`
			Country       string `json:"country"`
			Township      string `json:"township"`
			BusinessAreas []struct {
				Location string `json:"location"`
				Name     string `json:"name"`
				Id       string `json:"id"`
			} `json:"businessAreas"`
			Building struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"building"`
			Neighborhood struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"neighborhood"`
			Citycode string `json:"citycode"`
		} `json:"addressComponent"`
		Aois []struct {
			Area     string `json:"area"`
			Type     string `json:"type"`
			Id       string `json:"id"`
			Location string `json:"location"`
			Adcode   string `json:"adcode"`
			Name     string `json:"name"`
			Distance string `json:"distance"`
		} `json:"aois"`
		Pois []struct {
			Id           string      `json:"id"`
			Direction    string      `json:"direction"`
			Businessarea string      `json:"businessarea"`
			Address      string      `json:"address"`
			Poiweight    string      `json:"poiweight"`
			Name         string      `json:"name"`
			Location     string      `json:"location"`
			Distance     string      `json:"distance"`
			Tel          interface{} `json:"tel"`
			Type         string      `json:"type"`
		} `json:"pois"`
	} `json:"regeocode"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
}

type GeocodeRegeoResult struct {
	Result GeocodeRegeoResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
}

func newGeocodeRegeoResult(result GeocodeRegeoResponse, body []byte, http gorequest.Response) *GeocodeRegeoResult {
	return &GeocodeRegeoResult{Result: result, Body: body, Http: http}
}

// GeocodeRegeo 逆地理编码
// https://lbs.amap.com/api/webservice/guide/api/georegeo
func (c *Client) GeocodeRegeo(ctx context.Context, notMustParams ...*gorequest.Params) (*GeocodeRegeoResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("key", c.GetKey())
	params.Set("output", "JSON")
	// 请求
	request, err := c.request(ctx, apiUrl+"/geocode/regeo", params, http.MethodGet)
	if err != nil {
		return newGeocodeRegeoResult(GeocodeRegeoResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response GeocodeRegeoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGeocodeRegeoResult(response, request.ResponseBody, request), err
}
