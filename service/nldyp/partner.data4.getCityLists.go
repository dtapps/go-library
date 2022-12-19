package nldyp

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
)

type PartnerData4GetCityListsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Id       string `json:"id"`        // 城市ID
		CityName string `json:"city_name"` // 城市名
		Letter   string `json:"letter"`    // 首字母
		Hot      int    `json:"hot"`       // 是否热门：0 否 1 是
	} `json:"data"`
}

type PartnerData4GetCityListsResult struct {
	Result PartnerData4GetCityListsResponse // 结果
	Body   []byte                           // 内容
	Http   gorequest.Response               // 请求
	Err    error                            // 错误
}

func newPartnerData4GetCityListsResult(result PartnerData4GetCityListsResponse, body []byte, http gorequest.Response, err error) *PartnerData4GetCityListsResult {
	return &PartnerData4GetCityListsResult{Result: result, Body: body, Http: http, Err: err}
}

// PartnerData4GetCityLists 获取城市
// https://docs.apipost.cn/preview/fa101f4865dc783f/66e7c2e894fda4a6?target_id=d8842641-00bd-4bb4-a031-fb6d89908742
func (c *Client) PartnerData4GetCityLists(ctx context.Context) *PartnerData4GetCityListsResult {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(ctx, apiUrl+"/partner/data4/getCityLists", params)
	// 定义
	var response PartnerData4GetCityListsResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newPartnerData4GetCityListsResult(response, request.ResponseBody, request, err)
}
