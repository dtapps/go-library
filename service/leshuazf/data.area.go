package leshuazf

import (
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type DataAreaResponse struct {
	RespCode    string `json:"respCode"`
	RespMsg     string `json:"respMsg"`
	ReqSerialNo string `json:"reqSerialNo"`
	Data        []struct {
		AreaName       string `json:"areaName"`
		AreaCode       string `json:"areaCode"`
		ParentAreaCode string `json:"parentAreaCode"`
	} `json:"data"`
}

type DataAreaResult struct {
	Result DataAreaResponse   // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
	Err    error              // 错误
}

func newDataAreaResult(result DataAreaResponse, body []byte, http gorequest.Response, err error) *DataAreaResult {
	return &DataAreaResult{Result: result, Body: body, Http: http, Err: err}
}

// DataArea 代理商通过地区信息来查地区详细信息
// https://www.yuque.com/leshuazf/doc/dbmxyi#YwJl7
func (c *Client) DataArea(notMustParams ...gorequest.Params) *DataAreaResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request("/data/area", params, http.MethodPost)
	// 定义
	var response DataAreaResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newDataAreaResult(response, request.ResponseBody, request, err)
}
