package taobao

import "encoding/json"

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
	Err    error                 // 错误
}

func NewTbkTPwdCreateResult(result TbkTPwdCreateResponse, body []byte, err error) *TbkTPwdCreateResult {
	return &TbkTPwdCreateResult{Result: result, Body: body, Err: err}
}

// TbkTPwdCreate 淘宝客-公用-淘口令生成
// https://open.taobao.com/api.htm?docId=31127&docType=2&source=search
func (app *App) TbkTPwdCreate(notMustParams ...Params) *TbkTPwdCreateResult {
	// 参数
	params := NewParamsWithType("taobao.tbk.tpwd.create", notMustParams...)
	// 请求
	body, err := app.request(params)
	// 定义
	var response TbkTPwdCreateResponse
	err = json.Unmarshal(body, &response)
	return NewTbkTPwdCreateResult(response, body, err)
}
