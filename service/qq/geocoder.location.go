package qq

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type GeocoderLocationResponse struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
	Result    struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		Address            string `json:"address"`
		FormattedAddresses struct {
			Recommend string `json:"recommend"`
			Rough     string `json:"rough"`
		} `json:"formatted_addresses"`
		AddressComponent struct {
			Nation       string `json:"nation"`
			Province     string `json:"province"`
			City         string `json:"city"`
			District     string `json:"district"`
			Street       string `json:"street"`
			StreetNumber string `json:"street_number"`
		} `json:"address_component"`
		AdInfo struct {
			NationCode    string `json:"nation_code"`
			Adcode        string `json:"adcode"`
			PhoneAreaCode string `json:"phone_area_code"`
			CityCode      string `json:"city_code"`
			Name          string `json:"name"`
			Location      struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			Nation   string `json:"nation"`
			Province string `json:"province"`
			City     string `json:"city"`
			District string `json:"district"`
		} `json:"ad_info"`
		AddressReference struct {
			BusinessArea struct {
				Id       string `json:"id"`
				Title    string `json:"title"`
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Distance int    `json:"_distance"`
				DirDesc  string `json:"_dir_desc"`
			} `json:"business_area"`
			FamousArea struct {
				Id       string `json:"id"`
				Title    string `json:"title"`
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Distance int    `json:"_distance"`
				DirDesc  string `json:"_dir_desc"`
			} `json:"famous_area"`
			Crossroad struct {
				Id       string `json:"id"`
				Title    string `json:"title"`
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Distance float64 `json:"_distance"`
				DirDesc  string  `json:"_dir_desc"`
			} `json:"crossroad"`
			Town struct {
				Id       string `json:"id"`
				Title    string `json:"title"`
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Distance int    `json:"_distance"`
				DirDesc  string `json:"_dir_desc"`
			} `json:"town"`
			StreetNumber struct {
				Id       string `json:"id"`
				Title    string `json:"title"`
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Distance float64 `json:"_distance"`
				DirDesc  string  `json:"_dir_desc"`
			} `json:"street_number"`
			Street struct {
				Id       string `json:"id"`
				Title    string `json:"title"`
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Distance float64 `json:"_distance"`
				DirDesc  string  `json:"_dir_desc"`
			} `json:"street"`
			LandmarkL2 struct {
				Id       string `json:"id"`
				Title    string `json:"title"`
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Distance int    `json:"_distance"`
				DirDesc  string `json:"_dir_desc"`
			} `json:"landmark_l2"`
		} `json:"address_reference"`
		PoiCount int `json:"poi_count"`
		Pois     []struct {
			Id       string `json:"id"`
			Title    string `json:"title"`
			Address  string `json:"address"`
			Category string `json:"category"`
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			AdInfo struct {
				Adcode   string `json:"adcode"`
				Province string `json:"province"`
				City     string `json:"city"`
				District string `json:"district"`
			} `json:"ad_info"`
			Distance float64 `json:"_distance"`
			DirDesc  string  `json:"_dir_desc,omitempty"`
		} `json:"pois"`
	} `json:"result"`
}

type GeocoderLocationResult struct {
	Result GeocoderLocationResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
}

func newGeocoderLocationResult(result GeocoderLocationResponse, body []byte, http gorequest.Response) *GeocoderLocationResult {
	return &GeocoderLocationResult{Result: result, Body: body, Http: http}
}

// GeocoderLocation 逆地址解析（坐标位置描述）
// https://lbs.qq.com/service/webService/webServiceGuide/webServiceGcoder
func (c *Client) GeocoderLocation(ctx context.Context, location string, notMustParams ...*gorequest.Params) (*GeocoderLocationResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("key", c.GetKey())
	params.Set("location", location)
	params.Set("output", "JSON")
	// 请求
	request, err := c.request(ctx, apiUrl+"/ws/geocoder/v1/", params, http.MethodGet)
	if err != nil {
		return newGeocoderLocationResult(GeocoderLocationResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response GeocoderLocationResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGeocoderLocationResult(response, request.ResponseBody, request), err
}
