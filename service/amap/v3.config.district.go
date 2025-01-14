package amap

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type V3ConfigDistrictResponse struct {
	Status     string `json:"status"`   // 值为0或1 1：成功；0：失败
	Info       string `json:"info"`     // 返回的状态信息
	Infocode   string `json:"infocode"` // 返回状态说明,10000代表正确
	Count      string `json:"count"`    // 返回结果总数目
	Suggestion struct {
		Keywords []interface{} `json:"keywords"` // 建议关键字列表
		Cities   []interface{} `json:"cities"`   // 建议城市列表
	} `json:"suggestion"` // 建议结果列表
	Districts []struct {
		//Citycode  []interface{} `json:"citycode,omitempty"` // 城市编码
		Adcode    string `json:"adcode"`             // 区域编码
		Name      string `json:"name"`               // 行政区名称
		Polyline  string `json:"polyline,omitempty"` // 行政区边界坐标点
		Center    string `json:"center"`             // 区域中心点
		Level     string `json:"level"`              // 行政区划级别
		Districts []struct {
			//Citycode  []interface{} `json:"citycode,omitempty"` // 城市编码
			Adcode    string `json:"adcode"`             // 区域编码
			Name      string `json:"name"`               // 行政区名称
			Polyline  string `json:"polyline,omitempty"` // 行政区边界坐标点
			Center    string `json:"center"`             // 区域中心点
			Level     string `json:"level"`              // 行政区划级别
			Districts []struct {
				//Citycode  interface{} `json:"citycode,omitempty"` // 城市编码
				Adcode    string `json:"adcode"`   // 区域编码
				Name      string `json:"name"`     // 行政区名称
				Polyline  string `json:"polyline"` // 行政区边界坐标点
				Center    string `json:"center"`   // 区域中心点
				Level     string `json:"level"`    // 行政区划级别
				Districts []struct {
					//Citycode  interface{}   `json:"citycode,omitempty"`  // 城市编码
					Adcode    string        `json:"adcode"`              // 区域编码
					Name      string        `json:"name"`                // 行政区名称
					Polyline  string        `json:"polyline"`            // 行政区边界坐标点
					Center    string        `json:"center"`              // 区域中心点
					Level     string        `json:"level"`               // 行政区划级别
					Districts []interface{} `json:"districts,omitempty"` // 下级行政区列表
				} `json:"districts"`
			} `json:"districts"`
		} `json:"districts"`
	} `json:"districts"` // 行政区列表
}

type V3ConfigDistrictResult struct {
	Result V3ConfigDistrictResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
}

func newV3ConfigDistrictResult(result V3ConfigDistrictResponse, body []byte, http gorequest.Response) *V3ConfigDistrictResult {
	return &V3ConfigDistrictResult{Result: result, Body: body, Http: http}
}

// V3ConfigDistrict 行政区域查询
// https://lbs.amap.com/api/webservice/guide/api/district
func (c *Client) V3ConfigDistrict(ctx context.Context, notMustParams ...gorequest.Params) (*V3ConfigDistrictResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("key", c.GetKey())
	params.Set("output", "JSON")
	// 请求
	request, err := c.request(ctx, "config/district", params, http.MethodGet)
	if err != nil {
		return newV3ConfigDistrictResult(V3ConfigDistrictResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response V3ConfigDistrictResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newV3ConfigDistrictResult(response, request.ResponseBody, request), err
}
