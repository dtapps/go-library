package taobao

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
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
}

func newTbkTPwdCreateResult(result TbkTPwdCreateResponse, body []byte, http gorequest.Response) *TbkTPwdCreateResult {
	return &TbkTPwdCreateResult{Result: result, Body: body, Http: http}
}

// TbkTPwdCreate 淘宝客-公用-淘口令生成
// https://open.taobao.com/api.htm?docId=31127&docType=2&source=search
func (c *Client) TbkTPwdCreate(ctx context.Context, notMustParams ...gorequest.Params) (*TbkTPwdCreateResult, error) {
	// 参数
	params := NewParamsWithType("taobao.tbk.tpwd.create", notMustParams...)
	// 请求
	request, err := c.request(ctx, params)
	if err != nil {
		return newTbkTPwdCreateResult(TbkTPwdCreateResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response TbkTPwdCreateResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newTbkTPwdCreateResult(response, request.ResponseBody, request), err
}
