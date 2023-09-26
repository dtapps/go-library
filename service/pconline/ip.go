package pconline

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type IpResponse struct {
	Ip          string `json:"ip"`
	Pro         string `json:"pro"`
	ProCode     string `json:"proCode"`
	City        string `json:"city"`
	CityCode    string `json:"cityCode"`
	Region      string `json:"region"`
	RegionCode  string `json:"regionCode"`
	Addr        string `json:"addr"`
	RegionNames string `json:"regionNames"`
	Err         string `json:"err"`
}

type IpResult struct {
	Result IpResponse         // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newIpResult(result IpResponse, body []byte, http gorequest.Response) *IpResult {
	return &IpResult{Result: result, Body: body, Http: http}
}

// Ip 接口 https://whois.pconline.com.cn/
// ip=xxx
func (c *Client) Ip(ctx context.Context, ip string, notMustParams ...gorequest.Params) (*IpResult, error) { // 参数
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+fmt.Sprintf("/ipJson.jsp?json=true&ip=%s", ip), params)
	if err != nil {
		return newIpResult(IpResponse{}, request.ResponseBody, request), err
	}
	// 转码
	var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(request.ResponseBody)
	// 定义
	var response IpResponse
	err = gojson.Unmarshal(decodeBytes, &response)
	return newIpResult(response, request.ResponseBody, request), err
}
