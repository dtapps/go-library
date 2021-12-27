package taobao

type TbkTPwdCreateResult struct {
	TbkTpwdCreateResponse struct {
		Data struct {
			Model          string `json:"model"`
			PasswordSimple string `json:"password_simple"`
		} `json:"data"`
		RequestId string `json:"request_id"`
	} `json:"tbk_tpwd_create_response"`
}

// TbkTPwdCreate 淘宝客-公用-淘口令生成 https://open.taobao.com/api.htm?docId=31127&docType=2&source=search
func (app *App) TbkTPwdCreate(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := NewParamsWithType("taobao.tbk.tpwd.create", notMustParams...)
	// 请求
	body, err = app.request(params)
	return
}
