package qq

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type IpResponse struct {
	Status    int    `json:"status"`  // 状态码，0为正常，其它为异常
	Message   string `json:"message"` // 对status的描述
	RequestId string `json:"request_id"`
	Result    struct {
		Ip       string `json:"ip"` // 用于定位的IP地址
		Location struct {
			Lat float64 `json:"lat"` // 纬度
			Lng float64 `json:"lng"` // 经度
		} `json:"location"` // 定位坐标。注：IP定位服务精确到市级，该位置为IP地址所属的行政区划政府坐标
		AdInfo struct {
			Nation     string `json:"nation"`      // 国家
			Province   string `json:"province"`    // 国家代码（ISO3166标准3位数字码）
			City       string `json:"city"`        // 省
			District   string `json:"district"`    // 市
			Adcode     int    `json:"adcode"`      // 区
			NationCode int    `json:"nation_code"` // 行政区划代码

		} `json:"ad_info"` // 定位行政区划信息
	} `json:"result"` // IP定位结果
}

type IpResult struct {
	Result IpResponse         // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newIpResult(result IpResponse, body []byte, http gorequest.Response) *IpResult {
	return &IpResult{Result: result, Body: body, Http: http}
}

// Ip IP定位
// https://lbs.qq.com/service/webService/webServiceGuide/webServiceIp
func (c *Client) Ip(ctx context.Context, ip string, notMustParams ...gorequest.Params) (*IpResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("key", c.GetKey())
	params.Set("ip", ip)
	params.Set("output", "JSON")
	// 请求
	request, err := c.request(ctx, apiUrl+"/ws/location/v1/ip", params, http.MethodGet)
	if err != nil {
		return newIpResult(IpResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response IpResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newIpResult(response, request.ResponseBody, request), err
}
