package leshuazf

import (
	"encoding/json"
	"net/http"
)

type DataMccResponse struct {
	RespCode    string `json:"respCode"`
	RespMsg     string `json:"respMsg"`
	ReqSerialNo string `json:"reqSerialNo"`
	Data        []struct {
		GrandpaCode string `json:"grandpaCode"`
		GrandpaText string `json:"grandpaText"`
	} `json:"data"`
}

type DataMccResult struct {
	Result DataMccResponse // 结果
	Body   []byte          // 内容
	Err    error           // 错误
}

func NewDataMccResult(result DataMccResponse, body []byte, err error) *DataMccResult {
	return &DataMccResult{Result: result, Body: body, Err: err}
}

// DataMcc 代理商通过MccCode来查商户类别明细
// https://www.yuque.com/leshuazf/doc/dbmxyi#jRTHN
func (app *App) DataMcc(notMustParams ...Params) *DataMccResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request("data/mcc", params, http.MethodPost)
	// 定义
	var response DataMccResponse
	err = json.Unmarshal(body, &response)
	return NewDataMccResult(response, body, err)
}