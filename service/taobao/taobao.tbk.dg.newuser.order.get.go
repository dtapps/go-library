package taobao

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
)

type TbkDgNewuserOrderGetResponse struct {
	TbkDgNewuserOrderGetResponse struct {
		Data struct {
			Model          string `json:"model"`
			PasswordSimple string `json:"password_simple"`
		} `json:"data"`
		RequestId string `json:"request_id"`
	} `json:"tbk_tpwd_create_response"`
}

type TbkDgNewuserOrderGetResult struct {
	Result TbkDgNewuserOrderGetResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
	Err    error                        // 错误
}

func NewTbkDgNewuserOrderGetResult(result TbkDgNewuserOrderGetResponse, body []byte, http gorequest.Response, err error) *TbkDgNewuserOrderGetResult {
	return &TbkDgNewuserOrderGetResult{Result: result, Body: body, Http: http, Err: err}
}

// TbkDgNewuserOrderGet 淘宝客-推广者-新用户订单明细查询
// https://open.taobao.com/api.htm?docId=33892&docType=2
func (app *App) TbkDgNewuserOrderGet(notMustParams ...Params) *TbkDgNewuserOrderGetResult {
	// 参数
	params := NewParamsWithType("taobao.tbk.dg.newuser.order.get", notMustParams...)
	// 请求
	request, err := app.request(params)
	// 定义
	var response TbkDgNewuserOrderGetResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewTbkDgNewuserOrderGetResult(response, request.ResponseBody, request, err)
}
