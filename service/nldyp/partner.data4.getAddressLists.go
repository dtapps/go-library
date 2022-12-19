package nldyp

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
)

type PartnerData4GetAddressListsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Id       int    `json:"id"`        // 地区ID
		AreaName string `json:"area_name"` // 地区名
		CityId   string `json:"city_id"`   // 城市ID
	} `json:"data"`
}

type PartnerData4GetAddressListsResult struct {
	Result PartnerData4GetAddressListsResponse // 结果
	Body   []byte                              // 内容
	Http   gorequest.Response                  // 请求
	Err    error                               // 错误
}

func newPartnerData4GetAddressListsResult(result PartnerData4GetAddressListsResponse, body []byte, http gorequest.Response, err error) *PartnerData4GetAddressListsResult {
	return &PartnerData4GetAddressListsResult{Result: result, Body: body, Http: http, Err: err}
}

// PartnerData4GetAddressLists 获取地区
// https://docs.apipost.cn/preview/fa101f4865dc783f/66e7c2e894fda4a6?target_id=85053df1-09b5-4042-aeed-c7e10f3cdddc
func (c *Client) PartnerData4GetAddressLists(ctx context.Context) *PartnerData4GetAddressListsResult {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(ctx, apiUrl+"/partner/data4/getAddressLists", params)
	// 定义
	var response PartnerData4GetAddressListsResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newPartnerData4GetAddressListsResult(response, request.ResponseBody, request, err)
}
