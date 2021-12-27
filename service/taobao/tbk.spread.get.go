package taobao

type TbkSpreadGetResult struct {
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

// TbkSpreadGet 淘宝客-公用-长链转短链 https://open.taobao.com/api.htm?docId=27832&docType=2&source=search
func (app *App) TbkSpreadGet(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := NewParamsWithType("taobao.tbk.spread.get", notMustParams...)
	// 请求
	body, err = app.request(params)
	return
}
