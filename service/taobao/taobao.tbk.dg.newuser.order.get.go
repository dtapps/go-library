package taobao

import "encoding/json"

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
	Err    error                        // 错误
}

func NewTbkDgNewuserOrderGetResult(result TbkDgNewuserOrderGetResponse, body []byte, err error) *TbkDgNewuserOrderGetResult {
	return &TbkDgNewuserOrderGetResult{Result: result, Body: body, Err: err}
}

// TbkDgNewuserOrderGet 淘宝客-推广者-新用户订单明细查询
// https://open.taobao.com/api.htm?docId=33892&docType=2
func (app *App) TbkDgNewuserOrderGet(notMustParams ...Params) *TbkDgNewuserOrderGetResult {
	// 参数
	params := NewParamsWithType("taobao.tbk.dg.newuser.order.get", notMustParams...)
	// 请求
	body, err := app.request(params)
	// 定义
	var response TbkDgNewuserOrderGetResponse
	err = json.Unmarshal(body, &response)
	return NewTbkDgNewuserOrderGetResult(response, body, err)
}
