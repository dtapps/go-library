package taobao

import (
	"encoding/json"
	"go.dtapp.net/library/gorequest"
)

type TbkTPwdCreateResponse struct {
	TbkTpwdCreateResponse struct {
		Data struct {
			Model          string `json:"model"`
			PasswordSimple string `json:"password_simple"`
		} `json:"data"`
		RequestId string `json:"request_id"`
	} `json:"tbk_tpwd_create_response"`
}

type TbkTPwdCreateResult struct {
	Result TbkTPwdCreateResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
	Err    error                 // 错误
}

func NewTbkTPwdCreateResult(result TbkTPwdCreateResponse, body []byte, http gorequest.Response, err error) *TbkTPwdCreateResult {
	return &TbkTPwdCreateResult{Result: result, Body: body, Http: http, Err: err}
}

// TbkTPwdCreate 淘宝客-公用-淘口令生成
// https://open.taobao.com/api.htm?docId=31127&docType=2&source=search
func (app *App) TbkTPwdCreate(notMustParams ...Params) *TbkTPwdCreateResult {
	// 参数
	params := NewParamsWithType("taobao.tbk.tpwd.create", notMustParams...)
	// 请求
	request, err := app.request(params)
	// 定义
	var response TbkTPwdCreateResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewTbkTPwdCreateResult(response, request.ResponseBody, request, err)
}
