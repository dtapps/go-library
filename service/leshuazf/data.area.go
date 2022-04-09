package leshuazf

import (
	"encoding/json"
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
	Result DataAreaResponse // 结果
	Body   []byte           // 内容
	Err    error            // 错误
}

func NewDataAreaResult(result DataAreaResponse, body []byte, err error) *DataAreaResult {
	return &DataAreaResult{Result: result, Body: body, Err: err}
}

// DataArea 代理商通过地区信息来查地区详细信息
// https://www.yuque.com/leshuazf/doc/dbmxyi#YwJl7
func (app *App) DataArea(notMustParams ...Params) *DataAreaResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request("data/area", params, http.MethodPost)
	// 定义
	var response DataAreaResponse
	err = json.Unmarshal(body, &response)
	return NewDataAreaResult(response, body, err)
}
