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
	Err    error              // 错误
}

func newIpResult(result IpResponse, body []byte, http gorequest.Response, err error) *IpResult {
	return &IpResult{Result: result, Body: body, Http: http, Err: err}
}

// Ip 接口 https://whois.pconline.com.cn/
func (c *Client) Ip(ctx context.Context, ip string) *IpResult { // 参数
	param := gorequest.NewParams()
	// 请求
	request, err := c.request(ctx, apiUrl+fmt.Sprintf("/ipJson.jsp?json=true&ip=%s", ip), param)
	if err != nil {
		return newIpResult(IpResponse{}, request.ResponseBody, request, err)
	}
	// 转码
	var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(request.ResponseBody)
	// 定义
	var response IpResponse
	err = gojson.Unmarshal(decodeBytes, &response)
	return newIpResult(response, request.ResponseBody, request, err)
}
