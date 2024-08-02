package amap

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type IpResponse struct {
	Status    string `json:"status"`    // 值为0或1,0表示失败；1表示成功
	Info      string `json:"info"`      // 返回状态说明，status为0时，info返回错误原因，否则返回“OK”。
	Infocode  string `json:"infocode"`  // 返回状态说明,10000代表正确,详情参阅info状态表
	Province  string `json:"province"`  // 若为直辖市则显示直辖市名称； 如果在局域网 IP网段内，则返回“局域网”； 非法IP以及国外IP则返回空
	City      string `json:"city"`      // 若为直辖市则显示直辖市名称； 如果为局域网网段内IP或者非法IP或国外IP，则返回空
	Adcode    string `json:"adcode"`    // 城市的adcode编码
	Rectangle string `json:"rectangle"` // 所在城市矩形区域范围
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
// https://lbs.amap.com/api/webservice/guide/api/ipconfig
func (c *Client) Ip(ctx context.Context, ip string, notMustParams ...gorequest.Params) (*IpResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("key", c.GetKey())
	params.Set("ip", ip)
	// 请求
	request, err := c.request(ctx, "ip", params, http.MethodGet)
	if err != nil {
		return newIpResult(IpResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response IpResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newIpResult(response, request.ResponseBody, request), err
}
