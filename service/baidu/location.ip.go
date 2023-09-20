package baidu

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type LocationIpResponse struct {
	Address string `json:"address"` // 详细地址信息
	Content struct {
		AddressDetail struct {
			Province     string `json:"province"` // 省份
			City         string `json:"city"`     // 城市
			District     string `json:"district"`
			Street       string `json:"street"`
			StreetNumber string `json:"street_number"`
			CityCode     int    `json:"city_code"` // 百度城市代码
			Adcode       string `json:"adcode"`
		} `json:"address_detail"`
		Address string `json:"address"` // 简要地址信息
		Point   struct {
			X string `json:"x"` // 当前城市中心点经度
			Y string `json:"y"` // 当前城市中心点纬度
		} `json:"point"`
	} `json:"content"`
	Status int `json:"status"`
}

type LocationIpResult struct {
	Result LocationIpResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newLocationIpResult(result LocationIpResponse, body []byte, http gorequest.Response) *LocationIpResult {
	return &LocationIpResult{Result: result, Body: body, Http: http}
}

// LocationIp 普通IP定位
// https://lbsyun.baidu.com/index.php?title=webapi/ip-api
func (c *Client) LocationIp(ctx context.Context, ip string, notMustParams ...*gorequest.Params) (*LocationIpResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("ak", c.GetAk())
	params.Set("ip", ip)
	// 请求
	request, err := c.request(ctx, apiUrl+"/location/ip", params, http.MethodGet)
	if err != nil {
		return newLocationIpResult(LocationIpResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response LocationIpResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newLocationIpResult(response, request.ResponseBody, request), err
}
