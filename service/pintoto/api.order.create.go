package pintoto

import "encoding/json"

type ApiOrderCreateResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type ApiOrderCreateResult struct {
	Result ApiOrderCreateResponse // 结果
	Body   []byte                 // 内容
	Err    error                  // 错误
}

func NewApiOrderCreateResult(result ApiOrderCreateResponse, body []byte, err error) *ApiOrderCreateResult {
	return &ApiOrderCreateResult{Result: result, Body: body, Err: err}
}

// ApiOrderCreate 下单api https://www.showdoc.com.cn/1154868044931571/5891022916496848
func (app *App) ApiOrderCreate(notMustParams ...Params) *ApiOrderCreateResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request("https://movieapi2.pintoto.cn/api/order/create", params)
	// 定义
	var response ApiOrderCreateResponse
	err = json.Unmarshal(body, &response)
	return NewApiOrderCreateResult(response, body, err)
}
