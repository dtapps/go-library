package taobao

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
)

type TbkSpreadGetResponse struct {
	TbkSpreadGetResponse struct {
		Results struct {
			TbkSpread []struct {
				Content string `json:"content"`
				ErrMsg  string `json:"err_msg"`
			} `json:"tbk_spread"`
		} `json:"results"`
		TotalResults int    `json:"total_results"`
		RequestId    string `json:"request_id"`
	} `json:"tbk_spread_get_response"`
}

type TbkSpreadGetResult struct {
	Result TbkSpreadGetResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
	Err    error                // 错误
}

func NewTbkSpreadGetResult(result TbkSpreadGetResponse, body []byte, http gorequest.Response, err error) *TbkSpreadGetResult {
	return &TbkSpreadGetResult{Result: result, Body: body, Http: http, Err: err}
}

// TbkSpreadGet 淘宝客-公用-长链转短链
// https://open.taobao.com/api.htm?docId=27832&docType=2&source=search
func (app *App) TbkSpreadGet(notMustParams ...Params) *TbkSpreadGetResult {
	// 参数
	params := NewParamsWithType("taobao.tbk.spread.get", notMustParams...)
	// 请求
	request, err := app.request(params)
	// 定义
	var response TbkSpreadGetResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewTbkSpreadGetResult(response, request.ResponseBody, request, err)
}
